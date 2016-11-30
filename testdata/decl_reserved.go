// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package new

const static = 0

const (
	public = iota
	foo
	private
)

var class bool

var (
	enum int
	bar  int
	let  int
)

type function string

type try struct {
	private bool
	public  bool
}

func (t *try) with(in string) (void, super string) {
	new := func(this string) (typeof string) {
		v := this
		return v
	}
}

func do(in, void string) (super string) {
	return ""
}
