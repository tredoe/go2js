// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import "fmt"

var PASS = true

func builtIn() {
	pass := true

	var s1 []byte
	s2 := []byte{}
	s3 := make([]byte, 0)
	s4 := make([]byte, 0, 10)
	s5 := []int{1, 3, 5}

	tests := []struct {
		msg string
		in  bool
		out bool
	}{
		{"nil s1", s1 == nil, true},
		{"nil s2", s2 == nil, false},
		{"nil s3", s3 == nil, false},
		{"nil s4", s4 == nil, false},
		{"nil s5", s5 == nil, false},
		{"nil s5", s5 != nil, true},

		{"len s1", len(s1) == 0, true},
		{"len s2", len(s2) == 0, true},
		{"len s3", len(s3) == 0, true},
		{"len s4", len(s4) == 0, true},
		{"len s5", len(s5) == 3, true},

		{"cap s1", cap(s1) == 0, true},
		{"cap s2", cap(s2) == 0, true},
		{"cap s3", cap(s3) == 0, true},
		{"cap s4", cap(s4) == 10, true},
		{"cap s5", cap(s5) == 3, true},
	}

	for _, t := range tests {
		if t.in != t.out {
			fmt.Printf("\tFAIL: %s => got %v, want %v\n", t.msg, t.in, t.out)
			pass, PASS = false, false
		}
	}
	if pass {
		fmt.Println("\tpass")
	}
}

func shortHand() {
	pass := true

	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a_slice, b_slice []byte

	// == 1. Slice of an array

	a_slice = array[4:8]
	if string(a_slice) == "efgh" && len(a_slice) == 4 && cap(a_slice) == 6 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 1. [4:8] => got %v, len=%v, cap=%v\n",
			a_slice, len(a_slice), cap(a_slice))
		pass, PASS = false, false
	}

	a_slice = array[6:7]
	if string(a_slice) != "g" {
		fmt.Printf("\tFAIL: 1. [6:7] => got %v\n", a_slice)
		pass, PASS = false, false
	}

	a_slice = array[:3]
	if string(a_slice) == "abc" && len(a_slice) == 3 && cap(a_slice) == 10 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 1. [:3] => got %v, len=%v, cap=%v\n",
			a_slice, len(a_slice), cap(a_slice))
		pass, PASS = false, false
	}

	a_slice = array[5:]
	if string(a_slice) == "fghij" && len(a_slice) == 5 && cap(a_slice) == 5 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 1. [5:] => got %v, len=%v, cap=%v\n",
			a_slice, len(a_slice), cap(a_slice))
		pass, PASS = false, false
	}

	a_slice = array[:]
	if string(a_slice) == "abcdefghij" && len(a_slice) == 10 && cap(a_slice) == 10 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 1. [:] => got %v, len=%v, cap=%v\n",
			a_slice, len(a_slice), cap(a_slice))
		pass, PASS = false, false
	}

	a_slice = array[3:7]
	if string(a_slice) == "defg" && len(a_slice) == 4 && cap(a_slice) == 7 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 1. [3:7] => got %v, len=%v, cap=%v\n",
			a_slice, len(a_slice), cap(a_slice))
		pass, PASS = false, false
	}

	// == 2. Slice of a slice

	b_slice = a_slice[1:3]
	if string(b_slice) == "ef" && len(b_slice) == 2 && cap(b_slice) == 6 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 2. [1:3] => got %v, len=%v, cap=%v\n",
			b_slice, len(b_slice), cap(b_slice))
		pass, PASS = false, false
	}

	b_slice = a_slice[:3]
	if string(b_slice) == "def" && len(b_slice) == 3 && cap(b_slice) == 7 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 2. [:3] => got %v, len=%v, cap=%v\n",
			b_slice, len(b_slice), cap(b_slice))
		pass, PASS = false, false
	}

	b_slice = a_slice[:]
	if string(b_slice) == "defg" && len(b_slice) == 4 && cap(b_slice) == 7 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 2. [:] => got %v, len=%v, cap=%v\n",
			b_slice, len(b_slice), cap(b_slice))
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func useFunc() {
	pass := true

	// Returns the biggest value in a slice of ints.
	Max := func(slice []int) int {
		max := slice[0] // The first element is the max for now.
		for index := 1; index < len(slice); index++ {
			if slice[index] > max {
				max = slice[index]
			}
		}
		return max
	}

	A1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	A2 := [4]int{1, 2, 3, 4}
	A3 := [1]int{1}

	var slice []int

	slice = A1[:] // Take all A1 elements.
	if Max(slice) != 9 {
		fmt.Printf("\tFAIL: A1 => got %v, want 9\n", Max(slice))
		pass, PASS = false, false
	}

	slice = A2[:]
	if Max(slice) != 4 {
		fmt.Printf("\tFAIL: A2 => got %v, want 4\n", Max(slice))
		pass, PASS = false, false
	}

	slice = A3[:]
	if Max(slice) != 1 {
		fmt.Printf("\tFAIL: A3 => got %v, want 1\n", Max(slice))
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func reference() {
	pass := true

	A := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	slice1 := A[3:7]
	slice2 := A[5:]
	slice3 := slice1[:2]

	// == 1. Current content of A and the slices.

	tests := []struct {
		msg string
		in  string
		out string
	}{
		{"A", string(A[:]), "abcdefghij"},
		{"slice1", string(slice1), "defg"},
		{"slice2", string(slice2), "fghij"},
		{"slice3", string(slice3), "de"},
	}

	for _, t := range tests {
		if t.in != t.out {
			fmt.Printf("\tFAIL: 1. %s => got %v, want %v\n", t.msg, t.in, t.out)
			pass, PASS = false, false
		}
	}

	// == 2. Let's change the 'e' in A to 'E'.
	A[4] = 'E'

	tests = []struct {
		msg string
		in  string
		out string
	}{
		{"A", string(A[:]), "abcdEfghij"},
		{"slice1", string(slice1), "dEfg"},
		{"slice2", string(slice2), "fghij"},
		{"slice3", string(slice3), "dE"},
	}

	for _, t := range tests {
		if t.in != t.out {
			fmt.Printf("\tFAIL: 2. %s => got %v, want %v\n", t.msg, t.in, t.out)
			pass, PASS = false, false
		}
	}

	// == 3. Let's change the 'g' in slice2 to 'G'.
	slice2[1] = 'G'

	tests = []struct {
		msg string
		in  string
		out string
	}{
		{"A", string(A[:]), "abcdEfGhij"},
		{"slice1", string(slice1), "dEfG"},
		{"slice2", string(slice2), "fGhij"},
		{"slice3", string(slice3), "dE"},
	}

	for _, t := range tests {
		if t.in != t.out {
			fmt.Printf("\tFAIL: 3. %s => got %v, want %v\n", t.msg, t.in, t.out)
			pass, PASS = false, false
		}
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func resize() {
	pass := true

	var slice []byte

	// == 1.
	slice = make([]byte, 4, 5) // [0 0 0 0]

	if len(slice) == 4 && cap(slice) == 5 &&
		slice[0] == 0 && slice[1] == 0 && slice[2] == 0 && slice[3] == 0 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 1. got %v, want [0 0 0 0]\n", slice)
		pass, PASS = false, false
	}

	// == 2.
	slice[1], slice[3] = 2, 3 // [0 2 0 3]

	if slice[0] == 0 && slice[1] == 2 && slice[2] == 0 && slice[3] == 3 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 2. got %v, want [0 2 0 3]\n", slice)
		pass, PASS = false, false
	}

	// == 3.
	slice = make([]byte, 2) // Resize: [0 0]

	if len(slice) == 2 && cap(slice) == 2 && slice[0] == 0 && slice[1] == 0 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 3. got %v, want [0 0]\n", slice)
		pass, PASS = false, false
	}
	//==

	if pass {
		fmt.Println("\tpass")
	}
}

func grow() {
	pass := true

	// Add elements to the slice.
	GrowIntSlice := func(slice []int, add int) []int {
		new_capacity := cap(slice) + add
		new_slice := make([]int, len(slice), new_capacity)
		for index := 0; index < len(slice); index++ {
			new_slice[index] = slice[index]
		}
		return new_slice
	}

	slice := []int{0, 1, 2, 3}

	// == 1.
	if len(slice) == 4 && cap(slice) == 4 &&
		slice[0] == 0 && slice[1] == 1 && slice[2] == 2 && slice[3] == 3 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 1. got %v, want [0 1 2 3]\n", slice)
		pass, PASS = false, false
	}

	// == 2.
	slice = GrowIntSlice(slice, 3)

	if len(slice) == 4 && cap(slice) == 7 &&
		slice[0] == 0 && slice[1] == 1 && slice[2] == 2 && slice[3] == 3 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 2. got %v, want [0 1 2 3]\n", slice)
		pass, PASS = false, false
	}

	// == 3.

	// Let's two elements to the slice
	// So we reslice the slice to add 2 to its original length
	slice = slice[:len(slice)+2] // We can do this because cap(slice) == 7
	slice[4], slice[5] = 4, 5

	if len(slice) == 6 && cap(slice) == 7 &&
		slice[0] == 0 && slice[1] == 1 && slice[2] == 2 && slice[3] == 3 &&
		slice[4] == 4 && slice[5] == 5 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 3. got %v, want [0 1 2 3 4 5]\n", slice)
		pass, PASS = false, false
	}
	//==

	if pass {
		fmt.Println("\tpass")
	}
}

func _copy() {
	pass := true

	var a = [...]byte{'0', '1', '2', '3', '4', '5', '6', '7'}
	var s = make([]byte, 6)
	var b = make([]byte, 5)

	n1 := copy(s, a[0:])
	if string(s) == "012345" && n1 == 6 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 1. => got %q, n=%v\n", string(s), n1)
		pass, PASS = false, false
	}

	n2 := copy(s, s[2:])
	if string(s) == "234545" && n2 == 4 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 2. => got %q, n=%v\n", string(s), n2)
		pass, PASS = false, false
	}

	n3 := copy(b, "Hello, World!")
	if string(b) == "Hello" && n3 == 5 {
		// ok
	} else {
		fmt.Printf("\tFAIL: 3. => got %q, n=%v\n", string(b), n3)
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func _append() {
	pass := true

	slice := []byte{'1', '2', '3'}
	if string(slice) == "123" && len(slice) == 3 {
	} else {
		fmt.Printf("\tFAIL: 1. => got %q, len=%d\n", string(slice), len(slice))
		pass, PASS = false, false
	}

	slice = append(slice, '4')
	if string(slice) == "1234" && len(slice) == 4 {
	} else {
		fmt.Printf("\tFAIL: 2. => got %q, len=%d\n", string(slice), len(slice))
		pass, PASS = false, false
	}

	slice = append(slice, '5', '6')
	if string(slice) == "123456" && len(slice) == 6 {
	} else {
		fmt.Printf("\tFAIL: 3. => got %q, len=%d\n", string(slice), len(slice))
		pass, PASS = false, false
	}

	slice = append(slice, '7', '8', '9')
	if string(slice) == "123456789" && len(slice) == 9 {
	} else {
		fmt.Printf("\tFAIL: 4. => got %q, len=%d\n", string(slice), len(slice))
		pass, PASS = false, false
	}

	// == A slice

	a_slice := []byte{'1', '2', '3'}
	b_slice := []byte{'7', '8', '9'}

	a_slice = append(a_slice, b_slice...)
	if string(a_slice) == "123789" && len(a_slice) == 6 {
	} else {
		fmt.Printf("\tFAIL: append a slice => got %q, len=%d\n",
			string(a_slice), len(a_slice))
		pass, PASS = false, false
	}

	// == Delete

	/*del := func(i int, slice []byte) []byte {
		switch i {
		case 0:
			slice = slice[1:]
		case len(slice) - 1:
			slice = slice[:len(slice)-1]
		default:
			slice = append(slice[:i], slice[i+1:]...)
		}
		return slice
	}

	slice = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	slice = del(5, slice)
	if string(slice) == "012346789" && len(slice) == 9 {
	} else {
		fmt.Printf("\tFAIL: delete 5th element => got %q, len=%d\n",
			string(slice), len(slice))
		pass, PASS = false, false
	}

	slice = del(0, slice)
	if string(slice) == "12346789" && len(slice) == 8 {
	} else {
		fmt.Printf("\tFAIL: delete first element => got %q, len=%d\n",
			string(slice), len(slice))
		pass, PASS = false, false
	}

	slice = del(len(slice)-1, slice)
	if string(slice) == "1234678" && len(slice) == 7 {
	} else {
		fmt.Printf("\tFAIL: delete last element => got %q, len=%d\n",
			string(slice), len(slice))
		pass, PASS = false, false
	}

	// == Simple delete

	simpleDel := func(i int, slice []byte) []byte {
		slice = append(slice[:i], slice[i+1:]...)
		return slice
	}

	slice = simpleDel(3, slice)
	if string(slice) == "123678" && len(slice) == 6 {
	} else {
		fmt.Printf("\tFAIL: (simple) delete 3rd element => got %q, len=%d\n",
			string(slice), len(slice))
		pass, PASS = false, false
	}

	slice = simpleDel(0, slice)
	if string(slice) == "23678" && len(slice) == 5 {
	} else {
		fmt.Printf("\tFAIL: (simple) delete first element => got %q, len=%d\n",
			string(slice), len(slice))
		pass, PASS = false, false
	}

	slice = simpleDel(len(slice)-1, slice)
	if string(slice) == "2367" && len(slice) == 4 {
	} else {
		fmt.Printf("\tFAIL: (simple) delete last element => got %q, len=%d\n",
			string(slice), len(slice))
		pass, PASS = false, false
	}*/

	// == []interface

	//var t []interface{}
	//t = append(t, 42, 3.1415, "foo") //  t == []interface{}{42, 3.1415, "foo"}

	if pass {
		fmt.Println("\tpass")
	}
}

func main() {
	fmt.Print("\n\n== Slices\n\n")

	fmt.Println("=== RUN builtIn")
	builtIn()
	fmt.Println("=== RUN shortHand")
	shortHand()
	fmt.Println("=== RUN useFunc")
	useFunc()
	fmt.Println("=== RUN reference")
	reference()
	fmt.Println("=== RUN resize")
	resize()
	fmt.Println("=== RUN grow")
	grow()
	fmt.Println("=== RUN copy")
	_copy()
	fmt.Println("=== RUN append")
	_append()

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Slices")
	}
}
