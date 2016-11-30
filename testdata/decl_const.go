// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package test

const Pi float64 = 3.14159265358979323846
const pi2 = Pi
const zero = 0.0 // untyped floating-point constant
const (
	size int32 = 1024
	eof        = -1 // untyped integer constant
)
const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", untyped integer and string constants
const u, v float32 = 0, 3   // u = 0.0, v = 3.0

// == iota

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays // this constant is not exported
)
const ( // iota is reset to 0
	a0 = iota // a0 == 0
	a1 = iota // a1 == 1
	a2 = iota // a2 == 2
)
const (
	b0 = 1 << iota // b0 == 1 (iota has been reset)
	b1 = 1 << iota // b1 == 2
	b2 = 1 << iota // b2 == 4
)
const (
	c0         = iota * 42 // c0 == 0     (untyped integer constant)
	c1 float64 = iota * 42 // c1 == 42.0  (float64 constant)
	c2         = iota * 42 // c2 == 84    (untyped integer constant)
)

const x = iota // x == 0 (iota has been reset)
const y = iota // y == 0 (iota has been reset)

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
	bit1, mask1                          // bit1 == 2, mask1 == 1
	_, _                                 // skips iota == 2
	bit3, mask3                          // bit3 == 8, mask3 == 7
)

func main() {
	const F = 1

	const (
		Fa = 2
		fb = 3
	)
}
