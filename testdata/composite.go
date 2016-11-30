// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import "fmt"

var PASS = true

type person struct {
	name string
	age  int
}

// Return the older person of p1 and p2, and the difference in their ages.
func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

// Return the older person in a group of 10 persons.
func older10(people [10]person) person {
	older := people[0] // The first one is the older for now.

	// Loop through the array and check if we could find an older person.
	for index := 1; index < 10; index++ { // We skipped the first element here.
		if people[index].age > older.age {
			older = people[index]
		}
	}
	return older
}

// == Array
//

func builtInArray() {
	pass := true

	// TODO
	//var pa1 *[4]int
	//pa2 := new([4]int) // *[4]int

	var a1 [5]int
	a2 := [5]int{}
	a3 := [5]int{2}
	a4 := [5]int{2, 4}

	a5 := [3][4]int{}
	a6 := [3][4][2]int{}

	tests := []struct {
		msg string
		in  bool
		out bool
	}{
		//{"nil a1", a1 == nil, true},
		//{"nil a2", a2 == nil, false},
		{"len a1", len(a1) == 5, true},
		{"len a2", len(a2) == 5, true},
		{"len a3", len(a3) == 5, true},
		{"len a4", len(a4) == 5, true},
		{"len a4", len(a4) != 5, false},

		{"cap a1", cap(a1) == 5, true},
		{"cap a2", cap(a2) == 5, true},
		{"cap a3", cap(a3) == 5, true},
		{"cap a4", cap(a4) == 5, true},

		{"len a5", len(a5) == 3, true},
		{"cap a5", cap(a5) == 3, true},
		{"len a5[0]", len(a5[0]) == 4, true},
		{"cap a5[0]", cap(a5[0]) == 4, true},
		{"len a5[1000]", len(a5[1000]) == 4, true},
		{"cap a5[1000]", cap(a5[1000]) == 4, true},

		{"len a6", len(a6) == 3, true},
		{"cap a6", cap(a6) == 3, true},
		{"len a6[0]", len(a6[0]) == 4, true},
		{"cap a6[0]", cap(a6[0]) == 4, true},
		{"len a6[0][0]", len(a6[0][0]) == 2, true},
		{"cap a6[0][0]", cap(a6[0][0]) == 2, true},
		{"len a6[0][1000]", len(a6[0][1000]) == 2, true},
		{"cap a6[0][1000]", cap(a6[0][1000]) == 2, true},
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

func initArray() {
	pass := true

	// Declare and initialize an array A of 10 person.
	array1 := [10]person{
		person{"", 0},
		person{"Paul", 23},
		person{"Jim", 24},
		person{"Sam", 84},
		person{"Rob", 54},
		person{"", 0},
		person{"", 0},
		person{"", 0},
		person{"Karl", 10},
		person{"", 0},
	}

	// Declare and initialize an array of 10 persons, but let the compiler guess the size.
	array2 := [...]person{ // Substitute '...' instead of an integer size.
		person{"", 0},
		person{"Paul", 23},
		person{"Jim", 24},
		person{"Sam", 84},
		person{"Rob", 54},
		person{"", 0},
		person{"", 0},
		person{"", 0},
		person{"Karl", 10},
		person{"", 0}}

	tests := []struct {
		msg string
		in  bool
		out bool
	}{
		{"len", len(array1) == len(array2), true},
		{"cap", cap(array1) == cap(array2), true},
		{"equality", array1 == array2, true},
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

func _array() {
	// Declare an example array variable of 10 person called 'array'.
	var array [10]person

	// Initialize some of the elements of the array, the others are by default
	// set to person{"", 0}
	array[1] = person{"Paul", 23}
	array[2] = person{"Jim", 24}
	array[3] = person{"Sam", 84}
	array[4] = person{"Rob", 54}
	array[8] = person{"Karl", 19}

	older := older10(array) // Call the function by passing it our array.

	if older.name == "Sam" {
		fmt.Println("\tpass")
	} else {
		fmt.Printf("\tFAIL: got %v, want Sam\n", older.name)
		PASS = false
	}
}

func multiArray() {
	// Declare and initialize an array of 2 arrays of 4 ints
	doubleArray_1 := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// Simplify the previous declaration, with the '...' syntax
	doubleArray_2 := [2][4]int{
		[...]int{1, 2, 3, 4}, [...]int{5, 6, 7, 8}}

	// Super simpification!
	doubleArray_3 := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	if doubleArray_1 == doubleArray_2 && doubleArray_2 == doubleArray_3 {
		fmt.Println("\tpass")
	} else {
		fmt.Print("\tFAIL: got different arraies\n")
		PASS = false
	}
}

// == Struct
//

func _struct() {
	pass := true

	var tom person
	tom.name, tom.age = "Tom", 18

	bob := person{age: 25, name: "Bob"} // specify the fields and their values
	paul := person{"Paul", 43}          // specify values of fields in their order

	TB_older, TB_diff := older(tom, bob)
	TP_older, TP_diff := older(tom, paul)
	BP_older, BP_diff := older(bob, paul)

	tests := []struct {
		msg       string
		inPerson  person
		outPerson person
		inDiff    int
		outDiff   int
	}{
		{"Tom,Bob", TB_older, bob, TB_diff, 7},
		{"Tom,Paul", TP_older, paul, TP_diff, 25},
		{"Bob,Paul", BP_older, paul, BP_diff, 18},
	}

	for _, t := range tests {
		if t.inPerson != t.outPerson {
			fmt.Printf("\tFAIL: %s => person got %v, want %v\n",
				t.msg, t.inPerson, t.outPerson)
			pass, PASS = false, false
		}
		if t.inDiff != t.outDiff {
			fmt.Printf("\tFAIL: %s => difference got %v, want %v\n",
				t.msg, t.inDiff, t.outDiff)
			pass, PASS = false, false
		}
	}
	if pass {
		fmt.Println("\tpass")
	}
}

func main() {
	fmt.Print("\n\n== Composite types\n\n")

	fmt.Println("=== RUN builtInArray")
	builtInArray()
	fmt.Println("=== RUN initArray")
	initArray()
	fmt.Println("=== RUN array")
	_array()
	fmt.Println("=== RUN multiArray")
	multiArray()

	fmt.Println("=== RUN struct")
	_struct()

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Composite types")
	}
}
