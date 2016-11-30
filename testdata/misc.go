// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import "fmt"

var PASS = true

// Functions using arguments of custom types in JS.

func argArray(arr [3]int) [3]int {
	pass := true

	if len(arr) == 3 && cap(arr) == 3 && arr[0] == 1 && arr[1] == 2 && arr[2] == 3 {
	} else {
		fmt.Print("\tFAIL: argArray\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
	return arr
}

func argEllipsis(arr [2]int) [2]int {
	pass := true

	if len(arr) == 2 && cap(arr) == 2 && arr[0] == 5 && arr[1] == 6 {
	} else {
		fmt.Print("\tFAIL: argEllipsis\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
	return arr
}

func argSlice(s []byte) []byte {
	pass := true

	if len(s) == 2 && cap(s) == 2 && string(s) == "89" && s[0] == '8' && s[1] == '9' {
	} else {
		fmt.Print("\tFAIL: argSlice\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
	return s
}

func argMap(m map[int]string) map[int]string {
	pass := true

	if len(m) == 2 && m[1] == "foo" && m[2] == "bar" {
	} else {
		fmt.Print("\tFAIL: argSlice\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
	return m
}

func main() {
	fmt.Print("\n\n== Miscellaneous\n\n")

	fmt.Println("=== RUN argArray")
	a := [3]int{1, 2, 3}
	a = argArray(a)
	argArray([3]int{1, 2, 3})

	fmt.Println("=== RUN argEllipsis")
	ell := [...]int{5, 6}
	ell = argEllipsis(ell)
	argEllipsis([...]int{5, 6})

	fmt.Println("=== RUN argSlice")
	s := []byte{'8', '9'}
	s = argSlice(s)
	argSlice([]byte{'8', '9'})

	fmt.Println("=== RUN argMap")
	m := map[int]string{1: "foo", 2: "bar"}
	m = argMap(m)
	argMap(map[int]string{1: "foo", 2: "bar"})

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Miscellaneous")
	}
}
