// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import "fmt"

var PASS = true
var rating = map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

func builtIn() {
	pass := true

	var m1 map[string]int
	m2 := map[string]int{}
	m3 := make(map[string]int)
	m4 := make(map[string]int, 10)

	tests := []struct {
		msg string
		in  bool
		out bool
	}{
		{"nil m1", m1 == nil, true},
		{"nil m2", m2 == nil, false},
		{"nil m3", m3 == nil, false},
		{"nil m4", m4 == nil, false},
		{"nil m4", m4 != nil, true},

		{"len m1", len(m1) == 0, true},
		{"len m2", len(m2) == 0, true},
		{"len m3", len(m3) == 0, true},
		{"len m4", len(m4) == 0, true},

		{"nil rating", rating != nil, true},
		{"len rating", len(rating) == 4, true},
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

func declaration() {
	pass := true

	var numbers map[string]float32 // declare a map of strings to ints
	numbers = make(map[string]float32)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["trois"] = 3 // trois is "three" in french

	// A map representing the rating given to some programming languages.
	rating1 := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

	// This is equivalent to writing more verbosely
	rating2 := make(map[string]float32)
	rating2["C"] = 5
	rating2["Go"] = 4.5
	rating2["Python"] = 4.5
	rating2["C++"] = 2

	tests := []struct {
		msg string
		in  float32
		out float32
	}{
		{`numbers["one"]`, numbers["one"], 1},
		{`numbers["ten"]`, numbers["ten"], 10},
		{`numbers["trois"]`, numbers["trois"], 3},

		{`rating["C"]`, rating1["C"], rating2["C"]},
		{`rating["Go"]`, rating1["Go"], rating2["Go"]},
		{`rating["Python"]`, rating1["Python"], rating2["Python"]},
		{`rating["C++"]`, rating1["C++"], rating2["C++"]},
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

func reference() {
	m := make(map[string]string)
	m["Hello"] = "Bonjour"

	m1 := m
	m1["Hello"] = "Salut" // Now: m["Hello"] == "Salut"

	if m["Hello"] == m1["Hello"] {
		fmt.Println("\tpass")
	} else {
		fmt.Printf("\tFAIL: m[\"Hello\"] => got %v, want %v\n", m["Hello"], m1["Hello"])
		PASS = false
	}
}

func keyNoExistent() {
	pass := true

	csharp_rating := rating["C#"]
	csharp_rating2, found := rating["C#"]

	multiDim := map[int]map[int]float32{1: {1: 1.1}, 2: {2: 2.2}}
	k_multiDim := multiDim[1][2]

	tests := []struct {
		msg string
		in  float32
		out float32
	}{
		{"csharp_rating", csharp_rating, 0.00},
		{"csharp_rating2", csharp_rating2, 0},
		{"k_multiDim", k_multiDim, 0},
	}

	for _, t := range tests {
		if t.in != t.out {
			fmt.Printf("\tFAIL: %s => got %v, want %v\n", t.msg, t.in, t.out)
			pass, PASS = false, false
		}
	}
	if found {
		fmt.Printf("\tFAIL: using comma => got %v, want %v\n", found, !found)
		pass, PASS = false, false
	}
	if pass {
		fmt.Println("\tpass")
	}
}

func deleteKey() {
	pass := true

	delete(rating, "C++")
	_, found := rating["C++"]

	if found {
		fmt.Printf("\tFAIL: got %v, want %v\n", found, !found)
		pass, PASS = false, false
	}
	if pass {
		fmt.Println("\tpass")
	}
}

func _range() {
	pass := true

	for key, value := range rating {
		switch key {
		case "C":
			if value != 5 {
				fmt.Printf("\tFAIL: %s => got %v, want 5\n", key, value)
				pass, PASS = false, false
			}
		case "Go":
			if value != 4.5 {
				fmt.Printf("\tFAIL: %s => got %v, want 4.5\n", key, value)
				pass, PASS = false, false
			}
		case "Python":
			if value != 4.5 {
				fmt.Printf("\tFAIL: %s => got %v, want 4.5\n", key, value)
				pass, PASS = false, false
			}
		default:
			fmt.Printf("\tFAIL: %s => no expected\n", key)
			pass, PASS = false, false
		}
	}

	// Omit the value.
	for key := range rating {
		if key != "C" && key != "Go" && key != "Python" {
			fmt.Printf("\tFAIL: key %q no expected\n", key)
			pass, PASS = false, false
		}
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func blankIdInRange() {
	pass := true

	// Return the biggest value in a slice of ints.
	Max := func(slice []int) int {
		max := slice[0] // The first element is the max for now.
		for _, value := range slice {
			if value > max { // We found a bigger value in our slice.
				max = value
			}
		}
		return max
	}

	var slice []int
	// Declare three arrays of different sizes, to test the function Max.
	A1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	A2 := [4]int{1, 2, 3, 4}
	A3 := [1]int{1}

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

func main() {
	fmt.Print("\n\n== Maps\n\n")

	fmt.Println("=== RUN builtIn")
	builtIn()
	fmt.Println("=== RUN declaration")
	declaration()
	fmt.Println("=== RUN reference")
	reference()
	fmt.Println("=== RUN keyNoExistent")
	keyNoExistent()
	fmt.Println("=== RUN deleteKey")
	deleteKey()
	fmt.Println("=== RUN range")
	_range()
	fmt.Println("=== RUN blankIdInRange")
	blankIdInRange()

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Maps")
	}
}
