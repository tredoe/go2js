// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

// dataStmt represents data for the statements.
type dataStmt struct {
	funcTotal int // number total of functions
	funcId    int // number of function
	blockId   int // number of block
	tabLevel  int // tabulation level
	idxResult int // for then be used in resultUseFunc

	lenCase int // number of "case" statements
	idxCase int // index in "case" statements

	initIsPointer  bool // the value initialized is a pointer?
	insertVar      bool
	isConst        bool
	isVar          bool
	isArray        bool // to close the parenthesis
	isFunc         bool
	returnBasicLit bool
	skipLbrace     bool // left brace
	skipSemicolon  bool
	wasFallthrough bool // the last statement was "fallthrough"?
	wasReturn      bool // the last statement was "return"?

	lastVarName string // for composite types
	recvVar     string // receiver variable (in methods)
	results     string // variables names that return must use

	resultUseFunc map[int]bool // for JS types: array, slice, map
}

// getStatement translates the Go statement.
func (tr *translation) getStatement(stmt ast.Stmt) {
	switch typ := stmt.(type) {

	// http://golang.org/doc/go_spec.html#Arithmetic_operators
	// https://developer.mozilla.org/en/JavaScript/Reference/Operators/Assignment_Operators
	//
	// godoc go/ast AssignStmt
	//  Lhs    []Expr
	//  TokPos token.Pos   // position of Tok
	//  Tok    token.Token // assignment token, DEFINE
	//  Rhs    []Expr
	case *ast.AssignStmt:
		// There is not variable's type in the assignment.
		tr.writeVar(typ.Lhs, typ.Rhs, nil, typ.Tok, false, false)

	// http://golang.org/doc/go_spec.html#Blocks
	// https://developer.mozilla.org/en/JavaScript/Reference/Statements/block
	//
	// godoc go/ast BlockStmt
	//  Lbrace token.Pos // position of "{"
	//  List   []Stmt
	//  Rbrace token.Pos // position of "}"
	case *ast.BlockStmt:
		tr.blockId++
		tr.vars[tr.funcId][tr.blockId] = make(map[string]bool)
		tr.addr[tr.funcId][tr.blockId] = make(map[string]bool)
		tr.maps[tr.funcId][tr.blockId] = make(map[string]struct{})
		tr.arrays[tr.funcId][tr.blockId] = make(map[string]struct{})
		tr.slices[tr.funcId][tr.blockId] = make(map[string]struct{})
		tr.structSlices[tr.funcId][tr.blockId] = make(map[string]struct{})
		tr.zeroType[tr.funcId][tr.blockId] = make(map[string]string)

		if !tr.skipLbrace {
			tr.WriteString("{")
		} else {
			tr.skipLbrace = false
		}

		for i, v := range typ.List {
			skipTab := false

			// Don't insert tabulation in both "case", "label" clauses
			switch v.(type) {
			case *ast.CaseClause, *ast.LabeledStmt:
				skipTab = true
			default:
				tr.tabLevel++
			}

			// Write tabulation
			if tr.addLine(v.Pos()) {
				tr.WriteString(strings.Repeat(TAB, tr.tabLevel))
			} else if i == 0 {
				tr.WriteString(SP)
			}

			tr.getStatement(v)

			if !skipTab {
				tr.tabLevel--
			}
		}

		if tr.addLine(typ.Rbrace) {
			tr.WriteString(strings.Repeat(TAB, tr.tabLevel))
		} else {
			tr.WriteString(SP)
		}

		tr.WriteString("}")
		tr.blockId--

	// godoc go/ast BranchStmt
	//  TokPos token.Pos   // position of Tok
	//  Tok    token.Token // keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)
	//  Label  *Ident      // label name; or nil
	case *ast.BranchStmt:
		/*label := ";"
		if typ.Label != nil {
			label = SP + typ.Label.Name + ";"
		}*/

		tr.addLine(typ.TokPos)

		switch typ.Tok {
		// http://golang.org/doc/go_spec.html#Break_statements
		// https://developer.mozilla.org/en/JavaScript/Reference/Statements/break
		case token.BREAK:
			tr.WriteString("break;")
		// http://golang.org/doc/go_spec.html#Continue_statements
		// https://developer.mozilla.org/en/JavaScript/Reference/Statements/continue
		case token.CONTINUE:
			tr.WriteString("continue;")
		// http://golang.org/doc/go_spec.html#Goto_statements
		// http://golang.org/doc/go_spec.html#Fallthrough_statements
		case token.FALLTHROUGH:
			tr.wasFallthrough = true
		case token.GOTO: // not used since "label" is not translated
			tr.addError("%s: goto directive", tr.fset.Position(typ.TokPos))
		}

	// godoc go/ast CaseClause
	//  Case  token.Pos // position of "case" or "default" keyword
	//  List  []Expr    // list of expressions or types; nil means default case
	//  Colon token.Pos // position of ":"
	//  Body  []Stmt    // statement list; or nil
	case *ast.CaseClause:
		// To check the last statements
		tr.wasReturn = false
		tr.wasFallthrough = false

		tr.idxCase++
		tr.addLine(typ.Case)

		if typ.List != nil {
			for i, expr := range typ.List {
				if i != 0 {
					tr.WriteString(SP)
				}
				tr.WriteString(fmt.Sprintf("case %s:", tr.getExpression(expr).String()))
			}
		} else {
			tr.WriteString("default:")

			if tr.idxCase != tr.lenCase {
				tr.addWarning("%s: 'default' clause above 'case' clause in switch statement",
					tr.fset.Position(typ.Pos()))
			}
		}

		if typ.Body != nil {
			for _, v := range typ.Body {
				if ok := tr.addLine(v.Pos()); ok {
					tr.WriteString(strings.Repeat(TAB, tr.tabLevel+1))
				} else {
					tr.WriteString(SP)
				}
				tr.getStatement(v)
			}
		}

		if !tr.wasFallthrough && !tr.wasReturn && tr.idxCase != tr.lenCase {
			tr.WriteString(SP + "break;")
		}

	// godoc go/ast DeclStmt
	//  Decl Decl
	case *ast.DeclStmt:
		switch decl := typ.Decl.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.CONST:
				tr.getConst(decl.TokPos, decl.Specs, false)
			case token.VAR:
				tr.getVar(decl.Specs, false)
			case token.TYPE:
				tr.getType(decl.Specs, false)
			default:
				panic("unreachable")
			}
		default:
			panic(fmt.Sprintf("unimplemented: %T", decl))
		}

	// godoc go/ast ExprStmt
	//  X Expr // expression
	case *ast.ExprStmt:
		tr.WriteString(tr.getExpression(typ.X).String() + ";")

	// http://golang.org/doc/go_spec.html#For_statements
	// https://developer.mozilla.org/en/JavaScript/Reference/Statements/for
	//
	// godoc go/ast ForStmt
	//  For  token.Pos // position of "for" keyword
	//  Init Stmt      // initialization statement; or nil
	//  Cond Expr      // condition; or nil
	//  Post Stmt      // post iteration statement; or nil
	//  Body *BlockStmt
	case *ast.ForStmt:
		tr.WriteString("for" + SP + "(")

		if typ.Init != nil {
			tr.getStatement(typ.Init)
		} else {
			tr.WriteString(";")
		}

		if typ.Cond != nil {
			tr.WriteString(SP)
			tr.WriteString(tr.getExpression(typ.Cond).String())
		}
		tr.WriteString(";")

		if typ.Post != nil {
			tr.WriteString(SP)
			tr.skipSemicolon = true
			tr.getStatement(typ.Post)
		}

		tr.WriteString(")" + SP)
		tr.getStatement(typ.Body)

	// http://golang.org/doc/go_spec.html#Go_statements
	//
	// godoc go/ast GoStmt
	//  Go   token.Pos // position of "go" keyword
	//  Call *CallExpr
	case *ast.GoStmt:
		tr.addError("%s: goroutine", tr.fset.Position(typ.Go))

	// http://golang.org/doc/go_spec.html#If_statements
	// https://developer.mozilla.org/en/JavaScript/Reference/Statements/if...else
	//
	// godoc go/ast IfStmt
	//  If   token.Pos // position of "if" keyword
	//  Init Stmt      // initialization statement; or nil
	//  Cond Expr      // condition
	//  Body *BlockStmt
	//  Else Stmt // else branch; or nil
	case *ast.IfStmt:
		if typ.Init != nil {
			tr.getStatement(typ.Init)
			tr.WriteString(SP)
		}

		tr.WriteString(fmt.Sprintf("if%s(%s)%s", SP, tr.getExpression(typ.Cond).String(), SP))
		tr.getStatement(typ.Body)

		if typ.Else != nil {
			tr.WriteString(SP + "else ")
			tr.getStatement(typ.Else)
		}

	// godoc go/ast IncDecStmt
	//  X      Expr
	//  TokPos token.Pos   // position of Tok
	//  Tok    token.Token // INC or DEC
	case *ast.IncDecStmt:
		tr.WriteString(tr.getExpression(typ.X).String() + typ.Tok.String())

		if tr.skipSemicolon {
			tr.skipSemicolon = false
		} else {
			tr.WriteString(";")
		}

	// http://golang.org/doc/go_spec.html#For_statements
	// https://developer.mozilla.org/en/JavaScript/Reference/Statements/for...in
	//
	// godoc go/ast RangeStmt
	//  For        token.Pos   // position of "for" keyword
	//  Key, Value Expr        // Value may be nil
	//  TokPos     token.Pos   // position of Tok
	//  Tok        token.Token // ASSIGN, DEFINE
	//  X          Expr        // value to range over
	//  Body       *BlockStmt
	case *ast.RangeStmt:
		expr := tr.getExpression(typ.X).String()
		key := tr.getExpression(typ.Key).String()
		value := ""
		isMap := false

		if tr.isType(structType, stripField(expr)) {
			expr = stripField(expr)
		} else if tr.isType(mapType, expr) {
			isMap = true
		}

		if typ.Value != nil {
			value = tr.getExpression(typ.Value).String()
			if typ.Tok == token.DEFINE {
				tr.WriteString(fmt.Sprintf("var %s;%s", value, SP))
			}
		}

		tr.WriteString(fmt.Sprintf("for%s(var %s in %s", SP, key, expr))
		if isMap {
			tr.WriteString(".v")
		}
		tr.WriteString(")" + SP)

		if typ.Value != nil {
			tr.WriteString(fmt.Sprintf("{%s=%s", SP+value+SP, SP+expr))
			if isMap {
				tr.WriteString(".get(" + key + ")[0];")
			} else {
				tr.WriteString("[" + key + "];")
			}

			tr.skipLbrace = true
		}

		tr.getStatement(typ.Body)

	// http://golang.org/doc/go_spec.html#Return_statements
	// https://developer.mozilla.org/en/JavaScript/Reference/Statements/return
	//
	// godoc go/ast ReturnStmt
	//  Return  token.Pos // position of "return" keyword
	//  Results []Expr    // result expressions; or nil
	case *ast.ReturnStmt:
		tr.wasReturn = true

		if typ.Results == nil {
			if tr.results != "" {
				tr.WriteString(tr.results)
			} else {
				tr.WriteString("return;")
			}
			tr.wasReturn = false
			break
		}

		// Multiple values
		if len(typ.Results) != 1 {
			results := ""
			for i, v := range typ.Results {
				if i != 0 {
					results += "," + SP
				}
				tr.idxResult = i
				results += tr.getExpression(v).String()
			}

			tr.WriteString("return [" + results + "];")
		} else {
			tr.idxResult = 0
			tr.WriteString("return " + tr.getExpression(typ.Results[0]).String() + ";")
		}
		tr.wasReturn = false

	// http://golang.org/doc/go_spec.html#Switch_statements
	// https://developer.mozilla.org/en/JavaScript/Reference/Statements/switch
	//
	// godoc go/ast SwitchStmt
	//  Switch token.Pos  // position of "switch" keyword
	//  Init   Stmt       // initialization statement; or nil
	//  Tag    Expr       // tag expression; or nil
	//  Body   *BlockStmt // CaseClauses only
	case *ast.SwitchStmt:
		tag := "true"
		tr.lenCase = len(typ.Body.List)
		tr.idxCase = 0

		if typ.Tag != nil {
			tag = tr.getExpression(typ.Tag).String()
		} else if typ.Init != nil {
			tr.getStatement(typ.Init)
			tr.WriteString(SP)
		}

		tr.WriteString(fmt.Sprintf("switch%s(%s)%s", SP, tag, SP))
		tr.getStatement(typ.Body)

	// == Not supported

	// http://golang.org/doc/go_spec.html#Defer_statements
	//
	// godoc go/ast DeferStmt
	//  Defer token.Pos // position of "defer" keyword
	//  Call  *CallExpr
	case *ast.DeferStmt:
		tr.addError("%s: defer directive", tr.fset.Position(typ.Defer))

	// http://golang.org/doc/go_spec.html#Labeled_statements
	// https://developer.mozilla.org/en/JavaScript/Reference/Statements/label
	//
	// godoc go/ast LabeledStmt
	//  Label *Ident
	//  Colon token.Pos // position of ":"
	//  Stmt  Stmt
	case *ast.LabeledStmt:
		tr.addError("%s: use of label", tr.fset.Position(typ.Pos()))

	default:
		panic(fmt.Sprintf("unimplemented: %T", stmt))
	}
}
