// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

// Declarations not supported

package test

import fmt "fmt" // Package implementing formatted I/O.
import (
	"os"

	"github.com/tredoe/goscript"
)

var (
	a       complex128
	b, c, d complex128
	e       complex128 = 1
	f, g    complex128 = -1, -2

	h = complex(1, 2)
)

// Array
var (
	a1 = new([2][4]complex64)
	a2 = [3][5]complex64{}
	a3 = [32]complex64{1, 2, 3, 4}
)

// Slice
var (
	s1 = make([]complex128, 10)
	s2 = []complex128{2, 4, 6}
	s3 = [...]complex128{1, 2, 3}
)

// Map
var (
	m1 = make(map[complex64]int, 100)
	m2 = make(map[string]complex64)
	m3 = map[complex64]string{
		1: "first",
		2: "second",
		3: "third",
	}
)

// Chan
var (
	c1 = make(chan int, 10)
	c2 = make(chan bool)
	c3 = <-0
)

// == Struct
type i int

type s1 struct {
	a, b int
	c    float64
	_    float32 // padding
	F    func()
}

type s2 struct {
	a int64
	i
	f complex128
}

func main() {}
