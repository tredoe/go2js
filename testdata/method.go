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

type Rectangle struct {
	width, height float64
}

func noMethod() {
	pass := true

	area := func(r Rectangle) float64 {
		return r.width * r.height
	}

	r1 := Rectangle{12, 2}

	if area(r1) != 24 {
		fmt.Printf("\tFAIL: area r1 => got %v, want 24)\n", area(r1))
		pass, PASS = false, false
	}
	if area(Rectangle{9, 4}) != 36 {
		fmt.Printf("\tFAIL: area Rectangle{9,4} => got %v, want 36)\n",
			area(Rectangle{9, 4}))
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

// * * *

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func method() {
	pass := true

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	tests := []struct {
		msg string
		in  float64
		out float64
	}{
		{"Rectangle{12,2}", r1.area(), 24},
		{"Rectangle{9,4}", r2.area(), 36},
		{"Circle{10}", c1.area(), 314.1592653589793},
		{"Circle{25}", c2.area(), 1963.4954084936207},
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

// * * *

type sliceOfints []int
type agesByNames map[string]int

func (s sliceOfints) sum() int {
	sum := 0
	for _, value := range s {
		sum += value
	}
	return sum
}

func (people agesByNames) older() string {
	a := 0
	n := ""
	for key, value := range people {
		if value > a {
			a = value
			n = key
		}
	}
	return n
}

func withNamedType() {
	pass := true

	s := sliceOfints{1, 2, 3, 4, 5}
	folks := agesByNames{
		"Bob":   36,
		"Mike":  44,
		"Jane":  30,
		"Popey": 100,
	}

	if s.sum() != 15 {
		fmt.Printf("\tFAIL: s.sum => got %v, want 15)\n", s.sum())
		pass, PASS = false, false
	}
	if folks.older() != "Popey" {
		fmt.Printf("\tFAIL: folks.older => got %s, want Popey)\n",
			folks.older())
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

// * * *

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func complexNamedType() {
	pass := true

	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{20, 20, 20, YELLOW},
		Box{10, 30, 1, WHITE},
	}

	if len(boxes) != 6 {
		fmt.Printf("\tFAIL: len boxes => got %d, want 6\n", len(boxes))
		pass, PASS = false, false
	}
	if boxes[0].Volume() != 64 {
		fmt.Printf("\tFAIL: the volume of the first one => got %d, want 64\n",
			boxes[0].Volume())
		pass, PASS = false, false
	}
	if boxes[len(boxes)-1].color.String() != "WHITE" {
		fmt.Printf("\tFAIL: the color of the last one => got %s, want WHITE\n",
			boxes[len(boxes)-1].color.String())
		pass, PASS = false, false
	}
	if boxes.BiggestsColor().String() != "YELLOW" {
		fmt.Printf("\tFAIL: the biggest one => got %s, want YELLOW\n",
			boxes.BiggestsColor().String())
		pass, PASS = false, false
	}

	// Let's paint them all black
	boxes.PaintItBlack()

	if boxes[1].color.String() != "BLACK" {
		fmt.Printf("\tFAIL: the color of the second one => got %s, want BLACK\n",
			boxes[1].color.String())
		pass, PASS = false, false
	}
	if boxes.BiggestsColor().String() != "BLACK" {
		fmt.Printf("\tFAIL: finally, the biggest one => got %s, want BLACK\n",
			boxes.BiggestsColor().String())
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

// * * *

func main() {
	fmt.Print("\n\n== Methods\n\n")

	fmt.Println("=== RUN noMethod")
	noMethod()
	fmt.Println("=== RUN method")
	method()
	fmt.Println("=== RUN withNamedType")
	withNamedType()
	fmt.Println("=== RUN complexNamedType")
	complexNamedType()

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Methods")
	}
}
