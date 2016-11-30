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
	"strconv"
	"strings"
)

// Constants
//
// http://golang.org/doc/go_spec.html#Constant_declarations
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/const

// getConst translates a constant.
func (tr *translation) getConst(pos token.Pos, spec []ast.Spec, isGlobal bool) {
	iotaExpr := make([]string, 0) // iota expressions
	isMultipleLine := false
	tr.isConst = true

	if len(spec) > 1 {
		isMultipleLine = true
		tr.addLine(pos)
		tr.WriteString("const ")
	}

	// godoc go/ast ValueSpec
	//  Doc     *CommentGroup // associated documentation; or nil
	//  Names   []*Ident      // value names (len(Names) > 0)
	//  Type    Expr          // value type; or nil
	//  Values  []Expr        // initial values; or nil
	//  Comment *CommentGroup // line comments; or nil
	for i, s := range spec {
		vSpec := s.(*ast.ValueSpec)

		// Type checking
		if tr.getExpression(vSpec.Type).hasError {
			continue
		}

		tr.addLine(vSpec.Pos())
		isFirst := true

		for i, ident := range vSpec.Names {
			if ident.Name == "_" {
				iotaExpr = append(iotaExpr, "")
				continue
			}

			value := strconv.Itoa(ident.Obj.Data.(int)) // possible value of iota

			if vSpec.Values != nil {
				v := vSpec.Values[i]

				expr := tr.getExpression(v)
				if expr.hasError {
					continue
				}

				if expr.useIota {
					exprStr := expr.String()
					value = strings.Replace(exprStr, IOTA, value, -1)
					iotaExpr = append(iotaExpr, exprStr)
				} else {
					value = expr.String()
				}
			} else {
				if tr.hasError {
					continue
				}
				value = strings.Replace(iotaExpr[i], IOTA, value, -1)
			}

			if isGlobal {
				tr.addIfExported(ident)
			}

			// == Write
			name := validIdent(ident.Name)

			if isFirst {
				isFirst = false

				if !isGlobal && isMultipleLine {
					tr.WriteString(strings.Repeat(TAB, tr.tabLevel))
				}
				if isMultipleLine {
					tr.WriteString(name + SP + "=" + SP + value)
				} else {
					tr.WriteString(fmt.Sprintf("const %s=%s", name+SP, SP+value))
				}

			} else {
				tr.WriteString(fmt.Sprintf(",%s=%s", SP+name+SP, SP+value))
			}
		}

		// It is possible that there is only a blank identifier.
		if !isFirst {
			if isMultipleLine {
				if i == len(spec)-1 {
					tr.WriteString(";")
				} else {
					tr.WriteString(",")
				}
			} else {
				tr.WriteString(";")
			}
		}
	}

	tr.isConst = false
}

// Variables
//
// http://golang.org/doc/go_spec.html#Variable_declarations
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/var
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/let

// getVar translates a variable.
func (tr *translation) getVar(spec []ast.Spec, isGlobal bool) {
	isMultipleLine := false

	if len(spec) > 1 {
		isMultipleLine = true
	}

	// godoc go/ast ValueSpec
	for _, s := range spec {
		vSpec := s.(*ast.ValueSpec)

		// It is necessary to add the first variable before of checking
		tr.lastVarName = vSpec.Names[0].Name

		// Type checking
		if tr.getExpression(vSpec.Type).hasError {
			continue
		}

		tr.addLine(vSpec.Pos())
		// Pass token.DEFINE to know that it is a new variable
		tr.writeVar(vSpec.Names, vSpec.Values, vSpec.Type, token.DEFINE,
			isGlobal, isMultipleLine)
	}
}

// Types
//
// http://golang.org/doc/go_spec.html#Type_declarations

// getType translates a custom type.
func (tr *translation) getType(spec []ast.Spec, isGlobal bool) {
	// godoc go/ast TypeSpec
	//  Doc     *CommentGroup // associated documentation; or nil
	//  Name    *Ident        // type name
	//  Type    Expr          // *Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes
	//  Comment *CommentGroup // line comments; or nil
	for _, s := range spec {
		tSpec := s.(*ast.TypeSpec)

		// Type checking
		if tr.getExpression(tSpec.Type).hasError {
			continue
		}
		name := validIdent(tSpec.Name)

		switch typ := tSpec.Type.(type) {
		// godoc go/ast Ident
		//  NamePos token.Pos // identifier position
		//  Name    string    // identifier name
		//  Obj     *Object   // denoted object; or nil
		case *ast.StructType:
			tr.getStruct(typ, name, isGlobal)

		case *ast.ArrayType:
			tr.addLine(tSpec.Pos())

			if typ.Len != nil { // array
				tr.WriteString(fmt.Sprintf("function %s(){}%s.alias(g.ArrayType);",
					name, SP+name))
			} else { // slice
				tr.WriteString(fmt.Sprintf("function %s(){}%s.alias(g.SliceType);",
					name, SP+name))
			}
		case *ast.MapType:
			tr.addLine(tSpec.Pos())
			tr.WriteString(fmt.Sprintf("function %s(){}%s.alias(g.MapType);",
				name, SP+name))
				//"function %s(v,%szero)%s{%sg.mapType.apply(this,%sarguments)%s}%s",
				//name, SP, SP, SP, SP, SP, SP))

		case *ast.Ident:
			tr.addLine(tSpec.Pos())
			tr.WriteString(fmt.Sprintf("function %s(t)%s{%sthis%s=t;%s}",
				name, SP, SP, FIELD_TYPE, SP))
			tr.zeroType[tr.funcId][tr.blockId][name] = "FOO"
//fmt.Println(tSpec.Name.Name)

		default:
			/*tr.addLine(tSpec.Pos())
			tr.WriteString(fmt.Sprintf("function %s(t)%s{%sthis%s=arguments;%s}",
				validIdent(tSpec.Name), SP, SP, FIELD_TYPE, SP))*/
			panic(fmt.Sprintf("unimplemented: %T", typ))
		}

		if tr.hasError {
			continue
		}
		if isGlobal {
			tr.addIfExported(tSpec.Name)
		}
	}
}

// Struct
//

// getStruct translates a custom struct.
func (tr *translation) getStruct(typ *ast.StructType, name string, isGlobal bool) {
	// godoc go/ast StructType
	//  Struct     token.Pos  // position of "struct" keyword
	//  Fields     *FieldList // list of field declarations
	//  Incomplete bool       // true if (source) fields are missing in the Fields list
	//
	// godoc go/ast FieldList
	//  Opening token.Pos // position of opening parenthesis/brace, if any
	//  List    []*Field  // field list; or nil
	//  Closing token.Pos // position of closing parenthesis/brace, if any
	if typ.Incomplete {
		panic("list of fields incomplete ???")
	}

	var fieldNames, fieldLines, fieldsInit string
	//!anonField := make([]bool, 0) // anonymous field

	firstPos := tr.getLine(typ.Fields.Opening)
	posOldField := firstPos
	posNewField := 0
	isFirst := true

	// godoc go/ast Field
	//  Doc     *CommentGroup // associated documentation; or nil
	//  Names   []*Ident      // field/method/parameter names; or nil if anonymous field
	//  Type    Expr          // field/method/parameter type
	//  Tag     *BasicLit     // field tag; or nil
	//  Comment *CommentGroup // line comments; or nil
	for _, field := range typ.Fields.List {
		isPointer := false

		if _, ok := field.Type.(*ast.FuncType); ok {
			tr.addError("%s: function type in struct", tr.fset.Position(field.Pos()))
			continue
		}
		if field.Names == nil {
			tr.addError("%s: anonymous field in struct", tr.fset.Position(field.Pos()))
			continue
		}
		// Type checking
		if expr := tr.getExpression(field.Type); expr.hasError {
			continue
		} else if expr.isPointer {
			isPointer = true
		}

		zero, _ := tr.zeroValue(true, field.Type)

		for _, v := range field.Names {
			fieldName := validIdent(v.Name)
			if fieldName == "_" {
				continue
			}

			if !isFirst {
				fieldNames += "," + SP
				fieldsInit += "," + SP
			}
			fieldNames += fieldName
			fieldsInit += zero
			//!anonField = append(anonField, false)

			// == Printing of fields
			posNewField = tr.getLine(v.Pos())

			if posNewField != posOldField {
				fieldLines += strings.Repeat(NL, posNewField-posOldField)
				fieldLines += strings.Repeat(TAB, tr.tabLevel+1)
			} else {
				fieldLines += SP
			}

			if name != "" {
				fieldLines += fmt.Sprintf("this.%s=", fieldName)
			} else {
				fieldLines += fmt.Sprintf("%s:%s", fieldName, SP)
			}
			if !isPointer {
				fieldLines += fmt.Sprintf("%s", fieldName)
			} else {
				fieldLines += fmt.Sprintf("{p:%s}", fieldName)
			}
			if name != "" {
				fieldLines += ";"
			} else {
				fieldLines += ","
			}

			posOldField = posNewField
			// ==

			if isFirst {
				isFirst = false
			}
		}
	}
	// Remove the last character.
	if fieldLines != "" {
		fieldLines = fieldLines[:len(fieldLines)-1]
	}

	// The right brace
	posNewField = tr.getLine(typ.Fields.Closing)

	if posNewField != posOldField {
		fieldLines += strings.Repeat(NL, posNewField-posOldField)
		fieldLines += strings.Repeat(TAB, tr.tabLevel)
	} else {
		fieldLines += SP
	}

	// Empty structs
	if fieldLines == SP {
		fieldLines = ""
	}

	// == Write
	tr.addLine(typ.Pos())

	if name != "" {
		tr.WriteString(fmt.Sprintf("function %s(%s)%s{%s}",
			name, fieldNames, SP, fieldLines))
		//tr.WriteString(fmt.Sprintf("function %s(%s)%s{%sthis._z=%q;%s}",
		//validIdent(name), fieldNames, SP, SP, fieldsInit, fieldLines))

		// Store the name of new type with its values initialized
		tr.zeroType[tr.funcId][tr.blockId][name] = fieldsInit
	} else {
		tr.WriteString(fmt.Sprintf("_%s=%sfunction(%s)%s{%sreturn%s{%s};};%s",
			SP, SP, fieldNames, SP, SP, SP, fieldLines, SP))

		if _, ok := tr.structSlices[tr.funcId][tr.blockId][tr.lastVarName]; !ok {
			tr.structSlices[tr.funcId][tr.blockId][tr.lastVarName] = void
			tr.insertVar = true
		}
	}

	tr.line += posNewField - firstPos // update the global position
}

// == Utility
//

// writeVar translates variables for both declarations and assignments.
func (tr *translation) writeVar(names interface{}, values []ast.Expr, type_ interface{}, operator token.Token, isGlobal, isMultipleLine bool) {
	var sign string
	var signIsAssign, signIsDefine, isBitClear bool

	tr.isVar = true
	defer func() { tr.isVar = false }()

	if !isGlobal && isMultipleLine {
		tr.WriteString(strings.Repeat(TAB, tr.tabLevel))
	}

	// == Operator
	switch operator {
	case token.DEFINE:
		tr.WriteString("var ")
		sign = "="
		signIsDefine = true
	case token.ASSIGN:
		sign = operator.String()
		signIsAssign = true
	case token.ADD_ASSIGN, token.SUB_ASSIGN, token.MUL_ASSIGN, token.QUO_ASSIGN,
		token.REM_ASSIGN,
		token.AND_ASSIGN, token.OR_ASSIGN, token.XOR_ASSIGN, token.SHL_ASSIGN,
		token.SHR_ASSIGN:

		sign = operator.String()
	case token.AND_NOT_ASSIGN:
		sign = "&="
		isBitClear = true

	default:
		panic(fmt.Sprintf("operator unimplemented: %s", operator.String()))
	}

	// == Names
	// TODO: use this struct
	/*var Name = []struct {
		str      string
		idxValid int
		expr     *expression
	}{}*/

	var _names []string
	var idxValidNames []int // index of variables which are not in blank
	var name_expr []*expression

	switch t := names.(type) {
	case []*ast.Ident:
		_names = make([]string, len(t))
		name_expr = make([]*expression, len(t))

		for i, v := range t {
			expr := tr.getExpression(v)

			_names[i] = validIdent(expr.String())
			name_expr[i] = expr
		}
	case []ast.Expr: // like avobe
		_names = make([]string, len(t))
		name_expr = make([]*expression, len(t))

		for i, v := range t {
			expr := tr.getExpression(v)

			_names[i] = expr.String()
			name_expr[i] = expr
		}
	default:
		panic("unreachable")
	}

	// Check if there is any variable to use; and it is exported
	for i, v := range _names {
		if v != BLANK {
			idxValidNames = append(idxValidNames, i)

			if isGlobal {
				tr.addIfExported(v)
			}
		}
	}
	if len(idxValidNames) == 0 {
		return
	}

	if values != nil {
		// == Function
		if call, ok := values[0].(*ast.CallExpr); ok {

			// Function literal
			if _, ok := call.Fun.(*ast.SelectorExpr); ok {
				goto _noFunc
			}

			// Declaration of slice/array
			fun := call.Fun.(*ast.Ident).Name
			if fun == "make" || fun == "new" {
				goto _noFunc
			}

			// == Assign variable to the output of a function
			fun = tr.getExpression(call).String()

			if len(_names) == 1 {
				if tr.resultUseFunc[0] {
					_names[0] = stripField(_names[0])
				}
				tr.WriteString(_names[0] + SP + sign + SP + fun + ";")
				return
			}
			if len(idxValidNames) == 1 {
				i := idxValidNames[0]
				if tr.resultUseFunc[i] {
					_names[i] = stripField(_names[i])
				}
				tr.WriteString(fmt.Sprintf("%s%s%s[%d];", _names[i], SP+sign+SP, fun, i))
				return
			}

			// multiple variables
			str := fmt.Sprintf("_%s", SP+sign+SP+fun)

			for _, i := range idxValidNames {
				if tr.resultUseFunc[i] {
					_names[i] = stripField(_names[i])
				}
				str += fmt.Sprintf(",%s%s_[%d]", SP+_names[i], SP+sign+SP, i)
			}

			tr.WriteString(str + ";")
			return
		}
	}

_noFunc:
	expr := tr.newExpression(nil)
	typeIs := otherType
	isFuncLit := false
	isZeroValue := false
	isFirst := true
	value := ""
	numericFunc := ""

	if values == nil { // initialization explicit
		value, typeIs = tr.zeroValue(true, type_)
		isZeroValue = true
	}

	for iValidNames, idxName := range idxValidNames {
		name := _names[idxName]
		nameExpr := ""

		tr.lastVarName = name

		// == Name
		if isFirst {
			nameExpr += name
			isFirst = false
		} else {
			nameExpr += "," + SP + name
		}

		if !signIsDefine && len(name_expr[idxName].index) == 0 {
			nameExpr += tagPointer(false, 'P', tr.funcId, tr.blockId, name)
		}

		// == Value
		if isZeroValue {
			if typeIs == sliceType {
				tr.slices[tr.funcId][tr.blockId][name] = void
			}
		} else {
			var valueOfValidName ast.Expr

			// _, ok = m[k]
			if len(values) == 1 && idxName == 1 {
				valueOfValidName = values[0]
			} else {
				valueOfValidName = values[idxName]
			}

			// If the expression is an anonymous function, then, at translating,
			// it is written in the main buffer.
			expr = tr.newExpression(name)
			expr.isValue = true

			if _, ok := valueOfValidName.(*ast.FuncLit); !ok {
				expr.translate(valueOfValidName)
				exprStr := expr.String()

				if isBitClear {
					exprStr = "~(" + exprStr + ")"
				}
				value = exprStr

				_, typeIs = tr.zeroValue(false, type_)

				if expr.isVarAddress {
					tr.addr[tr.funcId][tr.blockId][name] = true
					if !signIsDefine {
						nameExpr += ADDR
					}
				} /*else {
					tr.addr[tr.funcId][tr.blockId][name] = false
				}*/

				// == Map: v, ok := m[k]
				if len(values) == 1 && tr.isType(mapType, expr.mapName) {
					value = value[:len(value)-3] // remove '[0]'

					if len(idxValidNames) == 1 {
						tr.WriteString(fmt.Sprintf("%s%s%s[%d];",
							_names[idxValidNames[0]],
							SP+sign+SP,
							value, idxValidNames[0]))
					} else {
						tr.WriteString(fmt.Sprintf("_%s,%s_[%d],%s_[%d];",
							SP+sign+SP+value,
							SP+_names[0]+SP+sign+SP, 0,
							SP+_names[1]+SP+sign+SP, 1))
					}

					return
				}
				// ==
			} else {
				isFuncLit = true

				tr.WriteString(nameExpr)
				expr.translate(valueOfValidName)
			}

			// Check if new variables assigned to another ones are slices or maps.
			if signIsDefine && expr.isIdent {
				if tr.isType(sliceType, value) {
					tr.slices[tr.funcId][tr.blockId][name] = void
				}
				if tr.isType(mapType, value) {
					tr.maps[tr.funcId][tr.blockId][name] = void
				}
			}
		}

		if signIsDefine {
			typeIsPointer := false
			if typeIs == pointerType {
				typeIsPointer = true
			}

			tr.vars[tr.funcId][tr.blockId][name] = typeIsPointer

			// The value could be a pointer so this new variable has to be it.
			if tr.vars[tr.funcId][tr.blockId][value] {
				tr.vars[tr.funcId][tr.blockId][name] = true
			}

			// Could be addressed ahead
			if value != "" && !expr.isPointer && !expr.isVarAddress && !typeIsPointer {
				value = tagPointer(isZeroValue, 'L', tr.funcId, tr.blockId, name) +
					value +
					tagPointer(isZeroValue, 'R', tr.funcId, tr.blockId, name)
			}
		}

		if !isFuncLit {
			// Insert "var" to variable of anonymous struct.
			if tr.insertVar && tr.isType(structType, name) {
				tr.WriteString("var ")
				tr.insertVar = false
			}
			tr.WriteString(nameExpr)

			/*switch expr.kind {
			case sliceKind:
			}*/

			if name_expr[idxName].addSet {
				tr.WriteString(SP + value + ")")

			} else if expr.kind == sliceKind || expr.isSliceExpr {
				if signIsDefine || signIsAssign {
					tr.slices[tr.funcId][tr.blockId][nameExpr] = void

					if value == "" {
						tr.WriteString(fmt.Sprintf("%sg.MkSlice(0,%s0)", SP+sign+SP, SP))
					} else {
						if expr.isSliceExpr {
							tr.WriteString(fmt.Sprintf("%sg.SliceFrom(%s)", SP+sign+SP, value))
						} else {
							tr.WriteString(fmt.Sprintf("%sg.Slice(%s)", SP+sign+SP, value))
						}
					}
				}
			} else if expr.isMake {
				tr.WriteString(fmt.Sprintf("%sg.MkSlice(%s)", SP+sign+SP, value))
				tr.slices[tr.funcId][tr.blockId][nameExpr] = void

			} else {
				if value != "" {
					// Get the numeric function
					if iValidNames == 0 {
						if ident, ok := type_.(*ast.Ident); ok {
							switch ident.Name {
							case "uint", "uint8", "uint16", "uint32",
								"int", "int8", "int16", "int32",
								"float32", "float64",
								"byte", "rune":
								numericFunc = "g." + strings.Title(ident.Name)
							}
						}
					}
					if numericFunc != "" {
						tr.WriteString(fmt.Sprintf("%s%s(%s)",
							SP+sign+SP, numericFunc, value))
					} else {
						tr.WriteString(SP + sign + SP + value)
					}
				}

				if tr.isArray {
					tr.WriteString(")")
					tr.isArray = false
				}
			}
		}
	}

	if !isFirst {
		tr.WriteString(";")
	}
}

// getTypeFields returns the fields of a custom type.
func (tr *translation) getTypeFields(fields []string) (args, allFields string) {
	for i, f := range fields {
		if i == 0 {
			args = f
		} else {
			args += "," + SP + f
			allFields += SP
		}

		allFields += fmt.Sprintf("this.%s=%s;", f, f)
	}
	return
}

// == Zero value
//

type dataType uint8

const (
	otherType dataType = iota
	mapType
	pointerType
	arrayType
	sliceType
	structType
)

// zeroValue returns the zero value of the value type if "init", and a boolean
// indicating if it is a pointer.
func (tr *translation) zeroValue(init bool, typ interface{}) (value string, dt dataType) {
	var ident *ast.Ident

	switch t := typ.(type) {
	case nil:
		return

	case *ast.ArrayType:
		if t.Len != nil { // array
			return tr.getExpression(t).String(), arrayType
		}

		// slice

		if !Bootstrap {
			return "g.MkSlice()", sliceType
		}
		return "[]", sliceType

	case *ast.InterfaceType: // nil
		return "undefined", otherType

	case *ast.MapType:
		tr.maps[tr.funcId][tr.blockId][tr.lastVarName] = void
		return fmt.Sprintf("g.Map(%s)", tr.zeroOfMap(t)), mapType

	case *ast.StructType:
		return "", structType

	case *ast.Ident:
		ident = t
	case *ast.StarExpr:
		tr.initIsPointer = true
		return tr.zeroValue(init, t.X)
	default:
		panic(fmt.Sprintf("zeroValue(): unexpected type: %T", typ))
	}

	if !init {
		if tr.initIsPointer {
			return "", pointerType
		}
		return
	}

	isType := true
	switch ident.Name {
	case "bool":
		value = "false"
	case "string":
		value = EMPTY
	case "uint", "uint8", "uint16", "uint32", "uint64",
		"int", "int8", "int16", "int32", "int64",
		"float32", "float64",
		"byte", "rune", "uintptr":
		value = "0"
	case "complex64", "complex128":
		value = "(0+0i)"
	default:
		value = ident.Name
		value = fmt.Sprintf("new %s(%s)", value, tr.zeroOfType(value))
		isType = false
	}

	if !Bootstrap && isType {
		//value = fmt.Sprintf("g.%s(%s)", strings.Title(ident.Name), value)
	}
	if tr.initIsPointer {
		value = "{p:undefined}"
		dt = pointerType
		tr.initIsPointer = false
	}
	return
}

// zeroOfMap returns the zero value of a map.
func (tr *translation) zeroOfMap(m *ast.MapType) string {
	if mapT, ok := m.Value.(*ast.MapType); ok { // nested map
		return tr.zeroOfMap(mapT)
	}
	v, _ := tr.zeroValue(true, m.Value)
	return v
}

// zeroOfType returns the zero value of a custom type.
func (tr *translation) zeroOfType(name string) string {
	// In the actual function
	if tr.funcId != 0 {
		for block := tr.blockId; block >= 0; block-- {
			if _, ok := tr.zeroType[tr.funcId][block][name]; ok {
				return tr.zeroType[tr.funcId][block][name]
			}
		}
	}

	// Finally, search in the global variables (funcId = 0).
	for block := tr.blockId; block >= 0; block-- { // block until 0
		if _, ok := tr.zeroType[0][block][name]; ok {
			return tr.zeroType[0][block][name]
		}
	}
	//fmt.Printf("Function %d, block %d, name %s\n", tr.funcId, tr.blockId, name)
	panic("zeroOfType: type not found: " + name)
}

// == Checking
//

// isType checks if a variable name is of a specific data type.
func (tr *translation) isType(t dataType, name string) bool {
	if name == "" {
		return false
	}
	name = strings.SplitN(name, "<<", 2)[0] // could have a tag

	for funcId := tr.funcId; funcId >= 0; funcId-- {
		for blockId := tr.blockId; blockId >= 0; blockId-- {
			// Avoid translation to Go types in functions parameters during bootstrap.
			if Bootstrap && blockId == 0 {
				return false
			}
			if _, ok := tr.vars[funcId][blockId][name]; ok { // variable found
				switch t {
				case arrayType:
					if _, ok = tr.arrays[funcId][blockId][name]; ok {
						return true
					}
				case mapType:
					if _, ok = tr.maps[funcId][blockId][name]; ok {
						return true
					}
				case sliceType:
					if _, ok = tr.slices[funcId][blockId][name]; ok {
						return true
					}
				case structType:
					if _, ok = tr.structSlices[funcId][blockId][name]; ok {
						return true
					}
				}
				return false
			}
		}
	}
	return false
}
