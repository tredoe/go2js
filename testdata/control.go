// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import "fmt"

var PASS = true

func _if() {
	pass := true

	// == Simple
	x := 5

	if x > 10 {
		fmt.Print("\tFAIL: simple\n")
		pass, PASS = false, false
	}

	// == Leading initial short
	if x := 12; x > 10 {
		// ok
	} else {
		fmt.Print("\tFAIL: with statement\n")
		pass, PASS = false, false
	}

	// == Multiple if/else
	i := 7

	if i == 3 {
		fmt.Print("\tFAIL: multiple (i == 3)\n")
		pass, PASS = false, false
	} else if i < 3 {
		fmt.Print("\tFAIL: multiple (i < 3)\n")
		pass, PASS = false, false
	} else {
		// ok
	}
	// ==

	if pass {
		fmt.Println("\tpass")
	}
}

func _switch() {
	pass := true

	// == Simple
	i := 10

	switch i {
	default:
		fmt.Print("\tFAIL: simple (default)\n")
		pass, PASS = false, false
	case 1:
		fmt.Print("\tFAIL: simple (1)\n")
		pass, PASS = false, false
	case 2, 3, 4:
		fmt.Print("\tFAIL: simple (2,3,4)\n")
		pass, PASS = false, false
	case 10:
		// ok
	}

	// == Without expression
	switch i = 5; {
	case i < 10:
		// ok
	case i > 10, i < 0:
		fmt.Print("\tFAIL: without expression (i>10, i<0)\n")
		pass, PASS = false, false
	case i == 10:
		fmt.Print("\tFAIL: without expression (i==10)\n")
		pass, PASS = false, false
	default:
		fmt.Print("\tFAIL: without expression (default)\n")
		pass, PASS = false, false
	}

	// == Without expression 2
	switch {
	case i == 5:
		// ok
	default:
		fmt.Print("\tFAIL: without expression 2 (default)\n")
		pass, PASS = false, false
	}

	// == With fallthrough
	switch i {
	case 4:
		pass = false
		fallthrough
	case 5:
		pass = false
		fallthrough
	case 6:
		pass = false
		fallthrough
	case 7:
		pass = true
	case 8:
		fmt.Print("\tFAIL: with fallthrough (8)\n")
		pass, PASS = false, false
	default:
		fmt.Print("\tFAIL: with fallthrough (default)\n")
		pass, PASS = false, false
	}

	if pass == false && PASS == true {
		fmt.Print("\tFAIL: with fallthrough (4,5,6)\n")
		PASS = false
	}
	// ==

	if pass {
		fmt.Println("\tpass")
	}
}

func _for() {
	pass := true

	// == Simple
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	if sum == 45 {
		// ok
	} else {
		fmt.Print("\tFAIL: simple\n")
		pass, PASS = false, false
	}

	// == Expression1 and expression3 are omitted here
	sum = 1
	for sum < 1000 {
		sum += sum
	}

	if sum == 1024 {
		// ok
	} else {
		fmt.Print("\tFAIL: 2 expressions omitted\n")
		pass, PASS = false, false
	}

	// == Expression1 and expression3 are omitted here, and semicolons gone
	sum = 1
	for sum < 1000 {
		sum += sum
	}

	if sum == 1024 {
		// ok
	} else {
		fmt.Print("\tFAIL: 2 expressions omitted, no semicolons\n")
		pass, PASS = false, false
	}

	// == Infinite loop (limited to show the output), no semicolons at all
	i := 0
	s := ""

	for {
		i++
		if i == 3 {
			s = fmt.Sprintf("%d", i)
			break
		}
	}

	if s == "3" {
		// ok
	} else {
		fmt.Print("\tFAIL: infinite loop\n")
		pass, PASS = false, false
	}

	// == break
	s = ""
	for i := 10; i > 0; i-- {
		if i < 5 {
			break
		}
		s += fmt.Sprintf("%d ", i)
	}

	if s == "10 9 8 7 6 5 " {
		// ok
	} else {
		fmt.Print("\tFAIL: break\n")
		pass, PASS = false, false
	}

	// == continue
	s = ""
	for i := 10; i > 0; i-- {
		if i == 5 {
			continue
		}
		s += fmt.Sprintf("%d ", i)
	}

	if s == "10 9 8 7 6 4 3 2 1 " {
		// ok
	} else {
		fmt.Print("\tFAIL: continue\n")
		pass, PASS = false, false
	}
	//==

	if pass {
		fmt.Println("\tpass")
	}
}

func _range() {
	pass := true

	s := []int{2, 3, 5}

	tests := map[int]int{
		0: 2,
		1: 3,
		2: 5,
	}

	for i, v := range s {
		if tests[i] != v {
			fmt.Printf("\tFAIL: %d. got %v, want %v\n", i, v, tests[i])
			pass, PASS = false, false
		}
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func main() {
	fmt.Print("\n\n== Control statements\n\n")

	fmt.Println("=== RUN if")
	_if()
	fmt.Println("=== RUN switch")
	_switch()
	fmt.Println("=== RUN for")
	_for()
	fmt.Println("=== RUN range")
	_range()

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Control statements")
	}
}
