// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
## Pointers

To identify variables that could be addressed ahead, it is used the map:

	{number of function: {number of block: {variable name: is pointer?} }}

In the generated code, it is added a tag before and after of each new variable
but pointer. The tag uses the schema `<<side:funcId:blockId:varName>>`

	*side:* *L* or *R* if the tag is on the left or on the right of the variable.
		*i* indicates that its value is zero for the value type.
	*funcId:* identifier of function. '0' is for global declarations
	*blockId:* number of block inner of that function. Start with '1'
	*varName:* variable's name

It is also added the tag `<<P:funcId:blockId:varName>>` after of each variable
name, and `<<&>>` after of it when the assignment is an address.
*/

// To remove tags related to pointers
var reTagPointer = regexp.MustCompile(`<<z?[LRP]:\d+:\d+:[^>]+>>`)

// tagPointer returns a tag to identify pointers.
// The argument field indicates if the variable is zero.
func tagPointer(zero bool, typ rune, funcId, blockId int, name string) string {
	/*if typ != 'L' && typ != 'R' && typ != 'P' {
		panic("invalid identifier for pointer: " + string(typ))
	}*/

	zeroStr := ""
	if zero {
		zeroStr = "z"
	}

	return fmt.Sprintf("<<%s:%d:%d:%s>>", zeroStr+string(typ), funcId, blockId, name)
}

// addPointer searches the point where the variable was declared for tag it as pointer.
func (tr *translation) addPointer(name string) {
	// In the actual function
	if tr.funcId != 0 {
		for block := tr.blockId; block >= 1; block-- {
			if _, ok := tr.vars[tr.funcId][block][name]; ok {
				tr.vars[tr.funcId][block][name] = true
				return
			}
		}
	}

	// Finally, search in the global variables (funcId = 0).
	for block := tr.blockId; block >= 0; block-- { // block until 0
		if _, ok := tr.vars[0][block][name]; ok {
			tr.vars[0][block][name] = true
			return
		}
	}
	//fmt.Printf("Function %d, block %d, name %s\n", tr.funcId, tr.blockId, name)
	panic("addPointer: unreachable")
}

// replacePointers replaces tags related to variables addressed.
func (tr *translation) replacePointers(str *string) {
	// Replaces tags in variables that access to pointers.
	replaceLocal := func(funcId, startBlock, endBlock int, varName string) {
		for block := startBlock; block <= endBlock; block++ {
			// Check if there is a variable named like the pointer in another block.
			if block != startBlock {
				if _, ok := tr.vars[funcId][block][varName]; ok {
					break
				}
			}

			pointer := tagPointer(false, 'P', funcId, block, varName)

			// Comparing pointers with value nil
			reNil := regexp.MustCompile(pointer + NIL)
			*str = reNil.ReplaceAllString(*str, FIELD_POINTER)

			if tr.addr[funcId][block][varName] {
				*str = strings.Replace(*str, pointer+ADDR, "", -1)
			} else {
				*str = strings.Replace(*str, pointer, FIELD_POINTER, -1)
			}
		}
	}

	// In each function
	for funcId, blocks := range tr.vars {
		for blockId := 0; blockId <= len(blocks); blockId++ {
			for name, isPointer := range tr.vars[funcId][blockId] {
				if isPointer {
					replaceLocal(funcId, blockId, len(blocks), name)

					// Replace brackets around variables addressed.
					lBrack := tagPointer(false, 'L', funcId, blockId, name)
					rBrack := tagPointer(false, 'R', funcId, blockId, name)

					*str = strings.Replace(*str, lBrack, "{p:", 1)
					*str = strings.Replace(*str, rBrack, "}", 1)

					// The declaration of pointers without initial value
					// have type "nil" in Go
					iLBrack := tagPointer(true, 'L', funcId, blockId, name)
					iRBrack := tagPointer(true, 'R', funcId, blockId, name)
					re := regexp.MustCompile(iLBrack + `[^<]+` + iRBrack)

					*str = re.ReplaceAllString(*str, "{p:undefined}")
				}
			}
		}
	}

	// == Global pointers
	globalScope := 0

	replaceGlobal := func(globVarName string) {
		for funcId, blocks := range tr.vars {
			if funcId == globalScope {
				continue
			}

			for blockId := 1; blockId <= len(blocks); blockId++ {
				if _, ok := tr.vars[funcId][blockId][globVarName]; ok {
					continue
				}

				pointer := tagPointer(false, 'P', funcId, blockId, globVarName)

				reNil := regexp.MustCompile(pointer + NIL)
				*str = reNil.ReplaceAllString(*str, FIELD_POINTER)

				if tr.addr[funcId][blockId][globVarName] {
					*str = strings.Replace(*str, pointer+ADDR, "", -1)
				} else {
					*str = strings.Replace(*str, pointer, FIELD_POINTER, -1)
				}
			}
		}
	}

	for globBlockId := 0; globBlockId <= len(tr.vars[globalScope]); globBlockId++ {
		for globName, isPointer := range tr.vars[globalScope][globBlockId] {
			if isPointer {
				replaceGlobal(globName)
			}
		}
	}

	// * * *

	// Remove the tags.
	*str = reTagPointer.ReplaceAllString(*str, "")
	*str = strings.Replace(*str, NIL, "", -1)
}
