// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

// Directives not supported

package test

func unvalidDirectives() {
	ch := make(chan int)

	go print("hello!")
	defer println("bye!")

	panic("problem")
	recover()
}

func _goto() {
	isFirst := true

_skipPoint:
	println("Using label")

	if isFirst {
		isFirst = false
		goto _skipPoint
		print("This part is skipped")
	}
}
