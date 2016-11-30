// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import "testing"

const (
	DIR_PKG  = "./jslib/"
	DIR_TEST = "./testdata/"
)

func init() {
	MaxMessage = 100 // to show all errors

	// Want see the tests in the HTML page.

	for _, v := range []string{"", "ln", "f"} {
		Function["fmt.Print"+v] = "document.write"
	}
	Function["print"] = "alert"
	Function["println"] = "alert"

	Char['\n'] = "<br>"
	Char['\t'] = "&nbsp;&nbsp;&nbsp;&nbsp;"
}

func TestConst(t *testing.T)    { translate('t', "decl_const.go", t) }
func TestVar(t *testing.T)      { translate('t', "decl_var.go", t) }
func TestStruct(t *testing.T)   { translate('t', "decl_struct.go", t) }
func TestReserved(t *testing.T) { translate('t', "decl_reserved.go", t) }
func TestPointer(t *testing.T)  { translate('t', "pointer.go", t) }

func TestFunc(t *testing.T)  { translate('t', "func.go", t) }
func TestCompo(t *testing.T) { translate('t', "composite.go", t) }
func TestSlice(t *testing.T) { translate('t', "slice.go", t) }
func TestMap(t *testing.T)   { translate('t', "map.go", t) }

func TestMethod(t *testing.T) { translate('t', "method.go", t) }

func TestNumeric(t *testing.T) { translate('t', "numeric.go", t) }
func TestMisc(t *testing.T)    { translate('t', "misc.go", t) }

func ExampleControl() {
	Translate(DIR_TEST+"control.go", true)

	// Output:
	// == Warnings
	//
	// ./testdata/control.go:58:2: 'default' clause above 'case' clause in switch statement
}

func ExampleDecl() {
	Translate(DIR_TEST+"error_decl.go", true)

	// Output:
	// == Errors
	//
	// os: import from core library
	// ./testdata/error_decl.go:19:10: complex128 type
	// ./testdata/error_decl.go:20:10: complex128 type
	// ./testdata/error_decl.go:21:10: complex128 type
	// ./testdata/error_decl.go:22:10: complex128 type
	// ./testdata/error_decl.go:24:6: built-in function complex()
	// ./testdata/error_decl.go:29:17: complex64 type
	// ./testdata/error_decl.go:30:13: complex64 type
	// ./testdata/error_decl.go:31:11: complex64 type
	// ./testdata/error_decl.go:36:14: complex128 type
	// ./testdata/error_decl.go:37:9: complex128 type
	// ./testdata/error_decl.go:38:12: complex128 type
	// ./testdata/error_decl.go:43:16: complex64 type
	// ./testdata/error_decl.go:44:23: complex64 type
	// ./testdata/error_decl.go:45:11: complex64 type
	// ./testdata/error_decl.go:54:12: channel type
	// ./testdata/error_decl.go:55:12: channel type
	// ./testdata/error_decl.go:56:7: channel operator
	// ./testdata/error_decl.go:66:2: function type in struct
	// ./testdata/error_decl.go:70:4: int64 type
	// ./testdata/error_decl.go:71:2: anonymous field in struct
	// ./testdata/error_decl.go:72:4: complex128 type
}

func ExampleStmt() {
	Translate(DIR_TEST+"error_stmt.go", true)

	// Output:
	// == Errors
	//
	// ./testdata/error_stmt.go:12:13: channel type
	// ./testdata/error_stmt.go:14:2: goroutine
	// ./testdata/error_stmt.go:15:2: defer directive
	// ./testdata/error_stmt.go:18:2: built-in function recover()
	// ./testdata/error_stmt.go:24:1: use of label
	// ./testdata/error_stmt.go:29:3: goto directive
}

// == JavaScript library

func TestLib(t *testing.T) { translate('p', "lib.go", t) }

// == Utility
//

func translate(kind rune, filename string, t *testing.T) {
	dir := ""

	if kind == 't' {
		dir = DIR_TEST
		Bootstrap = false
	} else if kind == 'p' {
		dir = DIR_PKG
		Bootstrap = true
	} else {
		panic("Wrong kind")
	}

	if err := Translate(dir+filename, true); err != nil {
		t.Fatalf("expected parse file: %s", err)
	}
}
