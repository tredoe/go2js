// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package test

var A string
var B bool
var a int
var b, c, d float64
var e = 0
var f, g float32 = -1, -2
var (
	h       int
	i, j, k = 2.0, 3.0, "bar"
)

var l = true   // l has type bool
var m = 0      // m has type int
var n = 3.0    // n has type float64
var o = "OMDB" // o has type string

// Array
var (
	a1 = new([32]byte)
	a2 = new([2][4]uint)
	//a3 = [2*N] struct { x, y int32 }
	a4 = [10]*float64{}
	a5 = [4]byte{}
	a6 = [3][5]int{}
	a7 = [2][2][2]float64{} // same as [2]([2]([2]float64))

	a8 = [32]byte{1, 2, 3, 4}
	a9 = [4]byte{1, 3: 4} // [1 0 0 4]

	a10 = [...]string{"a", "b", "c"} // [3]string
)

// Slice
var (
	s1 = make([]int, 10)
	s2 = make([]int, 10, 20)

	s3 = []int{2, 4, 6}
	s4 = []int{1, 2: 3} // [1 0 3]
	s5 = []int{}
)

// Map
var (
	m1 = make(map[string]int, 100) // map with initial space for 100 elements
	m2 = make(map[string]int)
	m3 = map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}
	m4 = map[int]interface{}{
		1: "first",
		2: 2,
		3: 3,
	}

	_, found = m4[1] // map lookup; only interested in "found"
)

// Pointer
var (
	p0 *byte
	p1 *int
	p2 *bool
)

func main() {
	Fa, Fb := 0, 10
	var Fc = "c"
	var (
		Fd uint = 20
		Fe float32
	)
}
