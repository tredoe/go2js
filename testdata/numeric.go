// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import "fmt"

var PASS = true

var (
	u   uint   = 1
	u_         = uint(1)
	u8  uint8  = 8
	u16 uint16 = 16
	u32 uint32 = 32

	i   int   = -1
	i_        = int(-1)
	i8  int8  = -8
	i16 int16 = -16
	i32 int32 = -32

	f32  float32 = 3.2
	f32_         = float32(3.2)
	f64  float64 = 6.4
	f64_         = float64(6.4)

	b  byte = 8
	b_      = byte(8)

	r  rune = 32
	r_      = rune(32)
)

//b3      = '1'
//b4      = 'a'
//r3      = '9'
//r4      = '€'

func value() {
	pass := true

	if u != 1 || u_ != 1 || u8 != 8 || u16 != 16 || u32 != 32 {
		fmt.Print("\tFAIL: uint\n")
		pass, PASS = false, false
	}
	if i != -1 || i_ != -1 || i8 != -8 || i16 != -16 || i32 != -32 {
		fmt.Print("\tFAIL: int\n")
		pass, PASS = false, false
	}
	if f32 != 3.2 || f32_ != 3.2 || f64 != 6.4 || f64_ != 6.4 {
		fmt.Print("\tFAIL: float\n")
		pass, PASS = false, false
	}
	if b != 8 || b_ != 8 {
		fmt.Print("\tFAIL: byte\n")
		pass, PASS = false, false
	}
	if r != 32 || r_ != 32 {
		fmt.Print("\tFAIL: rune\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func calculation() {
	pass := true

	if u+1 != 2 || u_+1 != 2 || u8+1 != 9 || u16+1 != 17 || u32+1 != 33 {
		fmt.Print("\tFAIL: add uint\n")
		pass, PASS = false, false
	}
	if i+1 != 0 || i_+1 != 0 || i8+1 != -7 || i16+1 != -15 || i32+1 != -31 {
		fmt.Print("\tFAIL: add int\n")
		pass, PASS = false, false
	}
	if f32+1 != 4.2 || f32_+1 != 4.2 || f64+1 != 7.4 || f64_+1 != 7.4 {
		fmt.Print("\tFAIL: add float\n")
		pass, PASS = false, false
	}
	if b+1 != 9 || b_+1 != 9 {
		fmt.Print("\tFAIL: add byte\n")
		pass, PASS = false, false
	}
	if r+1 != 33 || r_+1 != 33 {
		fmt.Print("\tFAIL: add rune\n")
		pass, PASS = false, false
	}

	if u8-1 != 7 || i8-1 != -9 || f32-1 != 2.2 || b-1 != 7 || r-1 != 31 {
		fmt.Print("\tFAIL: subtract\n")
		pass, PASS = false, false
	}

	if u16*2 != 32 || i16*2 != -32 || f64*2 != 12.8 || b*2 != 16 || r*2 != 64 {
		fmt.Print("\tFAIL: multiplication\n")
		pass, PASS = false, false
	}

	if u/1 != 1 || i/1 != -1 || f32/2 != 1.6 || b/2 != 4 || r/2 != 16 {
		fmt.Print("\tFAIL: division (quotient)\n")
		pass, PASS = false, false
	}
	if u8%3 != 2 || u16%3 != 1 || i8%3 != -2 || i16%3 != -1 {
		fmt.Print("\tFAIL: division (remainder)\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func bitwise() {
	pass := true

	if u16>>1 != 8 || u16<<1 != 32 || i16>>1 != -8 || i16<<1 != -32 {
		fmt.Print("\tFAIL: Shift\n")
		pass, PASS = false, false
	}

	if 7&9 != 1 || -7&9 != 9 {
		fmt.Print("\tFAIL: AND\n")
		pass, PASS = false, false
	}

	if 7|9 != 15 || -7|9 != -7 {
		fmt.Print("\tFAIL: OR\n")
		pass, PASS = false, false
	}

	if 7^9 != 14 || -7^9 != -16 {
		fmt.Print("\tFAIL: XOR\n")
		pass, PASS = false, false
	}

	if 7&^9 != 6 || -7&^9 != -16 {
		fmt.Print("\tFAIL: AND NOT\n")
		pass, PASS = false, false
	}

	n := 7
	n &^= 9
	if n != 6 {
		fmt.Print("\tFAIL: AND NOT (assignment)\n")
		pass, PASS = false, false
	}

	if ^-7 != 6 || ^7 != -8 {
		fmt.Print("\tFAIL: NOT\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func main() {
	fmt.Print("\n\n== Numeric\n\n")

	fmt.Println("=== RUN value")
	value()
	fmt.Println("=== RUN calculation")
	calculation()
	fmt.Println("=== RUN bitwise")
	bitwise()

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Numeric")
	}
}
