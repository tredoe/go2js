// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"math"
)

var PASS = true

var x = 10

func init() {
	x = 13
}

func _init() {
	if x == 13 {
		fmt.Println("\tpass")
	} else {
		fmt.Printf("\tFAIL: got %v, want 13\n", x)
		PASS = false
	}
}

func singleLine() { fmt.Println("\tpass") }

func simpleFunc() {
	pass := true

	// Returns the maximum between two int a, and b.
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) // calling max(x, y)
	if max_xy != 4 {
		fmt.Printf("\tFAIL: max(x,y) => got %v, want 4)\n", max_xy)
		pass, PASS = false, false
	}

	max_xz := max(x, z) // calling max(x, z)
	if max_xz != 5 {
		fmt.Printf("\tFAIL: max(x,z) => got %v, want 5)\n", max_xz)
		pass, PASS = false, false
	}

	if max(y, z) != 5 { // just call it here
		fmt.Printf("\tFAIL: max(y,z) => got %v, want 5)\n", max(y, z))
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func twoOuputValues() {
	pass := true

	// Returns A+B and A*B in a single shot.
	SumAndProduct := func(A, B int) (int, int) {
		return A + B, A * B
	}

	x := 3
	y := 4
	xPLUSy, xTIMESy := SumAndProduct(x, y)

	if xPLUSy != 7 {
		fmt.Printf("\tFAIL: sum => got %v, want 7)\n", xPLUSy)
		pass, PASS = false, false
	}
	if xTIMESy != 12 {
		fmt.Printf("\tFAIL: product => got %v, want 12)\n", xTIMESy)
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func resultVariable() {
	pass := true

	// Returns a bool that is set to true when Sqrt is possible and false when not,
	// and the actual square root of a float64.
	MySqrt := func(f float64) (s float64, ok bool) {
		if f > 0 {
			s, ok = math.Sqrt(f), true
		}
		return s, ok
	}

	tests := map[float64]float64{
		1:  1,
		2:  1.4142135623730951,
		3:  1.7320508075688772,
		4:  2,
		5:  2.23606797749979,
		6:  2.449489742783178,
		7:  2.6457513110645907,
		8:  2.8284271247461903,
		9:  3,
		10: 3.1622776601683795,
	}

	for i := -2.0; i <= 10; i++ {
		sqroot, ok := MySqrt(i)
		if ok {
			if sqroot != tests[i] {
				fmt.Printf("\tFAIL: square(%v) => got %v, want %v\n",
					i, sqroot, tests[i])
				pass, PASS = false, false
			}
		} else {
			if i != -2.0 && i != -1.0 && i != 0 {
				fmt.Printf("\tFAIL: square(%v) => should no be run\n", i)
				pass, PASS = false, false
			}
		}
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func _return() {
	pass := true

	MySqrt := func(f float64) (squareroot float64, ok bool) {
		if f > 0 {
			squareroot, ok = math.Sqrt(f), true
		}
		return // Omitting the output named variables, but keeping the "return".
	}

	_, ok := MySqrt(5)
	if !ok {
		fmt.Printf("\tFAIL: MySqrt(5) => got %v, want %v\n", ok, !ok)
		pass, PASS = false, false
	}

	if _, ok := MySqrt(0); ok {
		fmt.Printf("\tFAIL: MySqrt(0) => got %v, want %v\n", ok, !ok)
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func variadic() {
	pass := true

	type person struct {
		name string
		age  int
	}

	// Returns true and the older person in a group of persons,
	// or false and nil if the group is empty.
	getOlder := func(people ...person) (person, bool) {
		if len(people) == 0 {
			return person{}, false
		}

		older := people[0] // The first one is the older for now.

		for _, value := range people {
			if value.age > older.age {
				older = value
			}
		}
		return older, true
	}

	var (
		ok    bool
		older person
	)

	// Declare some persons.
	paul := person{"Paul", 23}
	jim := person{"Jim", 24}
	sam := person{"Sam", 84}
	rob := person{"Rob", 54}
	karl := person{"Karl", 19}

	tests := []struct {
		msg string
		out string
	}{
		{"paul,jim", "Jim"},
		{"paul,jim,sam", "Sam"},
		{"paul,jim,sam,rob", "Sam"},
		{"karl", "Karl"},
	}

	older, _ = getOlder(paul, jim)
	if older.name != tests[0].out {
		fmt.Printf("\tFAIL: (getOlder %s) => got %v, want %v\n",
			tests[0].msg, older.name, tests[0].out)
		pass, PASS = false, false
	}

	older, _ = getOlder(paul, jim, sam)
	if older.name != tests[1].out {
		fmt.Printf("\tFAIL: (getOlder %s) => got %v, want %v\n",
			tests[1].msg, older.name, tests[1].out)
		pass, PASS = false, false
	}

	older, _ = getOlder(paul, jim, sam, rob)
	if older.name != tests[2].out {
		fmt.Printf("\tFAIL: (getOlder %s) => got %v, want %v\n",
			tests[2].msg, older.name, tests[2].out)
		pass, PASS = false, false
	}

	older, _ = getOlder(karl)
	if older.name != tests[3].out {
		fmt.Printf("\tFAIL: (getOlder %s) => got %v, want %v\n",
			tests[3].msg, older.name, tests[3].out)
		pass, PASS = false, false
	}

	// There is no older person in an empty group.
	older, ok = getOlder()
	if ok {
		fmt.Printf("\tFAIL: (getOlder) => got %v, want %v\n", ok, !ok)
		pass, PASS = false, false
	}

	// == Multiple parameters

	getUser := func(name, surname string, age int, email ...string) string {
		emails := ""
		for _, v := range email {
			emails += " " + v
		}
		return fmt.Sprintf("%s %s, age %d, emails:%s", name, surname, age, emails)
	}

	name := "John"
	surname := "Smith"
	age := 17
	email1 := "foo@mail.se"
	email2 := "bar@mail.se"

	dataUser := getUser(name, surname, age, email1, email2)
	if dataUser != fmt.Sprintf("%s %s, age %d, emails: %s %s",
		name, surname, age, email1, email2) {
		fmt.Printf("\tFAIL: multiple parameters => got %q\n", dataUser)
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func Max(slice []int) int {
	if len(slice) == 1 {
		return slice[0]
	}

	middle := len(slice) / 2
	m1 := Max(slice[:middle])
	m2 := Max(slice[middle:])

	if m1 > m2 {
		return m1
	}
	return m2
}

func Invert(slice []byte) {
	length := len(slice)
	if length > 1 {
		slice[0], slice[length-1] = slice[length-1], slice[0] // Swap first and last ones
		Invert(slice[1 : length-1])
	}
}

func recursive() {
	pass := true

	s := []int{1, 2, 3, 4, 6, 8}

	if Max(s) != 8 {
		fmt.Printf("\tFAIL: Max => got %d, want 8\n", Max(s))
		pass, PASS = false, false
	}

	slice := []byte{'1', '2', '3', '4', '5'}
	Invert(slice)

	// TODO: comment out when multiple assignment is right
	/*if string(slice) != "54321" {
		fmt.Printf("\tFAIL: Invert => got %v, want \"54321\"\n", string(slice))
		pass, PASS = false, false
	}*/

	if pass {
		fmt.Println("\tpass")
	}
}

func A() {
	fmt.Println("\tRunning function A")
}

func B(name string) {
	fmt.Println("\tRunning function " + name)
}

/*func _defer() {
	defer A()

	defer func(){
		fmt.Println("\tRunning in-line function first")
	}()

	defer func(s string){
		fmt.Println("\tRunning in-line function " + s)
	}("second")

	defer B("B")
}*/

func main() {
	fmt.Print("\n\n== Functions\n\n")

	fmt.Println("=== RUN init")
	_init()
	fmt.Println("=== RUN singleLine")
	singleLine()
	fmt.Println("=== RUN simpleFunc")
	simpleFunc()
	fmt.Println("=== RUN twoOuputValues")
	twoOuputValues()
	fmt.Println("=== RUN resultVariable")
	resultVariable()
	fmt.Println("=== RUN return")
	_return()
	fmt.Println("=== RUN variadic")
	variadic()
	fmt.Println("=== RUN recursive")
	recursive()
	//fmt.Println("=== RUN defer")
	//_defer()

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Functions")
	}

	panic("unreachable")
	panic(fmt.Sprintf("not implemented: %s", "foo"))
}
