// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"go/ast"
	"strings"
)

// Functions
//
// http://golang.org/doc/go_spec.html#Function_declarations
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/function

// getFunc translates a function.
func (tr *translation) getFunc(decl *ast.FuncDecl) {
	// godoc go/ast FuncDecl
	//  Doc  *CommentGroup // associated documentation; or nil
	//  Recv *FieldList    // receiver (methods); or nil (functions)
	//  Name *Ident        // function/method name
	//  Type *FuncType     // position of Func keyword, parameters and results
	//  Body *BlockStmt    // function body; or nil (forward declaration)

	// Check empty functions
	if len(decl.Body.List) == 0 {
		return
	}

	isFuncInit := false // function init()

	// == Initialization to save variables created on this function
	//if decl.Name != nil { // discard literal functions //TODO: remove
		tr.funcTotal++
		tr.funcId = tr.funcTotal
		tr.blockId = 0

		tr.vars[tr.funcId] = make(map[int]map[string]bool)
		tr.addr[tr.funcId] = make(map[int]map[string]bool)
		tr.maps[tr.funcId] = make(map[int]map[string]struct{})
		tr.arrays[tr.funcId] = make(map[int]map[string]struct{})
		tr.slices[tr.funcId] = make(map[int]map[string]struct{})
		tr.structSlices[tr.funcId] = make(map[int]map[string]struct{})
		tr.zeroType[tr.funcId] = make(map[int]map[string]string)

		// The blockId 0 holds variables of functions arguments
		tr.vars[tr.funcId][tr.blockId] = make(map[string]bool)
		tr.maps[tr.funcId][tr.blockId] = make(map[string]struct{})
		tr.arrays[tr.funcId][tr.blockId] = make(map[string]struct{})
		tr.slices[tr.funcId][tr.blockId] = make(map[string]struct{})
	//}
	// ==

	tr.addLine(decl.Pos())
	tr.addIfExported(decl.Name)

	if decl.Name.Name != "init" {
		tr.writeFunc(decl.Recv, decl.Name, decl.Type)
	} else {
		isFuncInit = true
		tr.WriteString("(function()" + SP)
	}

	tr.getStatement(decl.Body)

	if isFuncInit {
		tr.WriteString("());")
	}

	if decl.Name != nil {
		// At exiting of the function, it returns at the global scope.
		tr.funcId = 0
		tr.blockId = 0

		if decl.Name.Name == "main" {
			tr.WriteString(SP + "main();") // call to function main
		}
	}
	if decl.Recv != nil {
		tr.recvVar = ""
	}
}

// godoc go/ast FuncType
//  Func    token.Pos  // position of "func" keyword
//  Params  *FieldList // (incoming) parameters; or nil
//  Results *FieldList // (outgoing) results; or nil

// godoc go/ast FieldList
//  Opening token.Pos // position of opening parenthesis/brace, if any
//  List    []*Field  // field list; or nil
//  Closing token.Pos // position of closing parenthesis/brace, if any

// godoc go/ast Field
//  Doc     *CommentGroup // associated documentation; or nil
//  Names   []*Ident      // field/method/parameter names; or nil if anonymous field
//  Type    Expr          // field/method/parameter type
//  Tag     *BasicLit     // field tag; or nil
//  Comment *CommentGroup // line comments; or nil

// writeFunc writes the function declaration.
func (tr *translation) writeFunc(recv *ast.FieldList, name *ast.Ident, typ *ast.FuncType) {
	if recv != nil { // method
		field := recv.List[0]
		tr.recvVar = field.Names[0].Name

		fType := tr.getExpression(field.Type).String()
		if strings.HasSuffix(fType, FIELD_POINTER) { // is it a pointer?
			fType = fType[:len(fType)-2]
		}

		tr.WriteString(fmt.Sprintf("%s.prototype.%s=%sfunction",
			fType, validIdent(name)+SP, SP))
	} else if name != nil {
		tr.WriteString(fmt.Sprintf("function %s", validIdent(name)))
		tr.recvVar = "_" // avoid that been added "this" in selectors
	} else { // Literal function
		tr.WriteString(fmt.Sprintf("%s=%sfunction", SP, SP))
		tr.recvVar = "_" // avoid "this" in selectors
	}

	// Get the parameters
	paramFix, paramVar := tr.joinParams(typ)

	tr.WriteString(fmt.Sprintf("(%s)%s", paramFix, SP))

	if paramVar != "" {
		tr.WriteString("{" + SP + paramVar)
		tr.skipLbrace = true
	}

	// To return multiple values
	declResults, declReturn := tr.joinResults(typ)

	if declResults != "" {
		if !tr.skipLbrace {
			tr.WriteString("{")
			tr.skipLbrace = true
		}
		tr.WriteString(SP + declResults)
		tr.results = declReturn
	} else {
		tr.results = ""
	}
}

// joinParams gets the parameters.
func (tr *translation) joinParams(f *ast.FuncType) (paramFix, paramVar string) {
	if f.Params == nil {
		return
	}
	isFirst := true
	i := 0

L:
	for _, list := range f.Params.List {
		typ := otherType

		switch t := list.Type.(type) {
		case *ast.Ellipsis:
			paramVar = fmt.Sprintf("var %s=%s",
				validIdent(list.Names[0].Name)+SP, SP)

			if i != 0 {
				paramVar += fmt.Sprintf("[].slice.call(arguments).slice(%d);", i)
			} else {
				paramVar += "arguments;"
			}
			break L // an ellipsis is the last parameter

		case *ast.ArrayType:
			if t.Len != nil {
				typ = arrayType
			} else {
				typ = sliceType
			}
		case *ast.MapType:
			typ = mapType
		}

		for _, v := range list.Names {
			if !isFirst {
				paramFix += "," + SP
			}
			i++
			_name := validIdent(v.Name)
			paramFix += _name

			if typ != otherType {
				tr.vars[tr.funcId][tr.blockId][_name] = false

				switch typ {
				case sliceType:
					tr.slices[tr.funcId][tr.blockId][_name] = void
				case arrayType:
					tr.arrays[tr.funcId][tr.blockId][_name] = void
				case mapType:
					tr.maps[tr.funcId][tr.blockId][_name] = void
				}
			}
			if isFirst {
				isFirst = false
			}
		}
	}
	return
}

// joinResults gets the results to use both in the declaration and in its return.
func (tr *translation) joinResults(f *ast.FuncType) (decl, ret string) {
	if f.Results == nil {
		return
	}
	isFirst := true
	isMultiple := false
	typeUseFunc := false
	i := 0

	for _, list := range f.Results.List {
		switch list.Type.(type) {
		case *ast.ArrayType, *ast.MapType, *ast.Ellipsis:
			typeUseFunc = true
		default:
			typeUseFunc = false
		}

		if list.Names == nil {
			tr.resultUseFunc[i] = typeUseFunc
			i++
			continue
		}

		value, _ := tr.zeroValue(true, list.Type)

		for _, v := range list.Names {
			if !isFirst {
				decl += "," + SP
				ret += "," + SP
				isMultiple = true
			} else {
				isFirst = false
			}
			decl += fmt.Sprintf("%s=%s", validIdent(v.Name)+SP, SP+value)
			ret += v.Name

			tr.resultUseFunc[i] = typeUseFunc
			i++
		}
	}

	if decl != "" {
		decl = "var " + decl + ";"
	}

	if isMultiple {
		ret = "[" + ret + "]"
	}
	ret = "return " + ret + ";"

	return
}
