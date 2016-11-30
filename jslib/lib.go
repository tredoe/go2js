// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

// Package g handles the features and Go types in JavaScript.

package g

// The specific type that it represents.
const (
	invalidT int = iota
	arrayT
	mapT
	sliceT
)

func init() {
	// Use the toString() method when Array.isArray isn't implemented:
	// https://developer.mozilla.org/en/JavaScript/Reference/Global_Objects/Array/isArray#Compatibility
	if !Array.isArray {
		Array.isArray = func(arg interface{}) bool {
			return Object.prototype.toString.call(arg) == "[object Array]"
		}
	}

	// Inheritance
	// http://phrogz.net/JS/classes/OOPinJS2.html
	Function.prototype.alias = func(parent interface{}) {
		if parent.constructor == Function { // Normal Inheritance
			this.prototype = parent //new(parent)
			this.prototype.constructor = this
			this.prototype.parent = parent.prototype
		} else { // Pure Virtual Inheritance
			this.prototype = parent
			this.prototype.constructor = this
			this.prototype.parent = parent
		}
		return this
	}
}

// == Boolean
//

type BoolType struct {
	v interface{} // value
	t string      // type
}

// Override the "valueOf" method to convert the object to the primitive value.
// https://developer.mozilla.org/en-US/docs/JavaScript/Reference/Global_Objects/Object/valueOf
func (b BoolType) valueOf() { return b.v }

func Bool(b bool) BoolType { return BoolType{b, "bool"} }

// == String
//

type StringType struct {
	v interface{}
	t string
}

func (s StringType) valueOf() { return s.v }

func String(s string) StringType { return StringType{s, "string"} }

// == Numeric types
//

type NumType struct {
	v interface{}
	t string
}

func (n NumType) valueOf() { return n.v }

//func (n NumType) toString() { return n.v }

// * * *

func Uint(n uint) NumType     { return NumType{n, "uint"} }
func Uint8(n uint8) NumType   { return NumType{n, "uint8"} }
func Uint16(n uint16) NumType { return NumType{n, "uint16"} }
func Uint32(n uint32) NumType { return NumType{n, "uint32"} }

func Int(n int) NumType     { return NumType{n, "int"} }
func Int8(n int8) NumType   { return NumType{n, "int8"} }
func Int16(n int16) NumType { return NumType{n, "int16"} }
func Int32(n int32) NumType { return NumType{n, "int32"} }

func Float32(n float32) NumType { return NumType{n, "float32"} }
func Float64(n float64) NumType { return NumType{n, "float64"} }

func Byte(n byte) NumType { return NumType{n, "byte"} }
func Rune(n rune) NumType { return NumType{n, "rune"} }

// == Array
//

// The array can not be compared with nil.
// The capacity is the same than length.

// ArrayType represents a fixed array type.
type ArrayType struct {
	v []interface{} // array's value

	len_ map[int]int
}

// len returns the length for the given dimension.
func (a ArrayType) len(index int) int {
	if index == nil {
		return a.len_[0]
	}
	return a.len_[len(arguments)]
}

// cap returns the capacity for the given dimension.
func (a ArrayType) cap(index int) int {
	if index == nil {
		return a.len_[0]
	}
	return a.len_[len(arguments)]
}

// str returns the array (of bytes or runes) like a string.
func (a ArrayType) str() string {
	return a.v.join("")
}

// typ returns the type.
func (a ArrayType) typ() int { return arrayT }

// MkArray initializes an array of dimension "index" to value "zero",
// merging the elements of "data" if any.
func MkArray(index []int, zero interface{}, data []interface{}) *ArrayType {
	a := new(ArrayType)

	if data != nil {
		if !equalIndex(index, indexArray(data)) {
			a.v = initArray(index, zero)
			mergeArray(a.v, data)
		} else {
			a.v = data
		}
	} else {
		a.v = initArray(index, zero)
	}

	for i, v := range index {
		a.len_[i] = v
	}

	return a
}

// * * *

// equalIndex reports whether index1 and index2 are equal.
func equalIndex(index1, index2 []int) bool {
	if len(index1) != len(index2) {
		return false
	}
	for i, v := range index1 {
		if v != index2[i] {
			return false
		}
	}
	return true
}

// indexArray returns the dimension of an array.
func indexArray(a []interface{}) (index []int) {
	for {
		index.push(len(a))

		if Array.isArray(a[0]) {
			a = a[0]
		} else {
			break
		}
	}
	return
}

// initArray returns an array of dimension given in "index" initialized to "zero".
func initArray(index []int, zero interface{}) (a []interface{}) {
	if len(index) == 0 {
		return zero
	}
	nextArray := initArray(index.slice(1), zero)

	for i := 0; i < index[0]; i++ {
		a[i] = nextArray
	}
	return
}

// mergeArray merges src in array dst.
func mergeArray(dst, src []interface{}) {
	for i, srcVal := range src {
		if Array.isArray(srcVal) {
			mergeArray(dst[i], srcVal)
		} else {
			isHashMap := false

			// The position is into a hash map, if any
			if typeof(srcVal) == "object" {
				for k, v := range srcVal {
					if srcVal.hasOwnProperty(k) { // identify a hashmap
						isHashMap = true
						i = k
						dst[i] = v
					}
				}
			}
			if !isHashMap {
				dst[i] = srcVal
			}
		}
	}
}

// == Slice
//

// SliceType represents a slice type.
type SliceType struct {
	arr interface{}   // the array where data is got or created from scratch using make
	v   []interface{} // elements appended

	low  int // indexes for the array
	high int
	len  int // total of elements
	cap  int

	nil_ bool // for variables declared like slices
}

func (s SliceType) isNil() bool {
	if s.len != 0 || s.cap != 0 {
		return false
	}
	return s.nil_
}

// typ returns the type.
func (s SliceType) typ() int { return sliceT }

// MkSlice initializes a slice with the zero value.
func MkSlice(zero interface{}, len, cap int) *SliceType {
	s := new(SliceType)

	if zero == nil {
		s.nil_ = true
		return s
	}

	arr := new(ArrayType)
	arr.len_[0] = len
	// The fastest way of fill in an array is when array length is specified first.
	arr.v = Array(len)
	for i := 0; i < len; i++ {
		arr.v[i] = zero
	}

	if cap != nil {
		s.cap = cap
	} else {
		s.cap = len
	}

	s.arr = arr
	s.len = len
	s.high = len

	return s
}

// Slice creates a new slice with the elements in "data".
func Slice(zero interface{}, data []interface{}) *SliceType {
	s := new(SliceType)

	if zero == nil {
		s.nil_ = true
		return s
	}

	arr := new(ArrayType)
	for i, srcVal := range data {
		isHashMap := false

		// The position is into a hash map, if any
		if typeof(srcVal) == "object" {
			for k, v := range srcVal {
				if srcVal.hasOwnProperty(k) { // identify a hashmap
					isHashMap = true

					for i; i < k; i++ {
						arr.v[i] = zero
					}
					arr.v[i] = v
				}
			}
		}
		if !isHashMap {
			arr.v[i] = srcVal
		}
	}
	s.len = len(arr.v)
	arr.len_[0] = s.len
	s.arr = arr

	s.cap = s.len
	s.high = s.len
	return s
}

// SliceFrom creates a new slice from an array or slice using the indexes low and high.
func SliceFrom(src interface{}, low, high int) *SliceType {
	s := new(SliceType)

	if low != nil {
		s.low = low | 0 // to integer
	} else {
		s.low = 0
	}
	if high != nil {
		s.high = high | 0 // to integer
	} else {
		if src.arr != nil { // slice
			s.high = src.len
		} else {
			s.high = len(src.v)
		}
	}

	s.len = s.high - s.low

	if src.arr != nil { // slice
		s.arr = src.arr
		s.cap = src.cap - s.low
		s.low += src.low
		s.high += src.low
	} else { // array
		s.arr = src
		s.cap = src.cap() - s.low
	}
	return s
}

// get gets the slice.
func (s SliceType) get() []interface{} {
	if s.arr != nil {
		arr := s.arr.v.slice(s.low, s.high)

		if len(s.v) != 0 {
			return arr.concat(s.v)
		} else {
			return arr
		}
	}
	return s.v
}

// set sets a value.
func (s SliceType) set(index []int, v interface{}) {
	s.arr.v[index[0]+s.low] = v
}

// str returns the slice (of bytes or runes) like a string.
func (s SliceType) str() string {
	_s := s.get()
	return _s.join("")
}

// * * *

// Append implements the function "append".
func Append(src []interface{}, elt ...interface{}) (dst SliceType) {
	// Copy src to the new slice
	dst.low = src.low
	dst.high = src.high
	dst.len = src.len
	dst.cap = src.cap
	dst.nil_ = src.nil_

	arr := new(ArrayType)
	arr.len_[0] = src.arr.len_[0]
	for _, v := range src.arr.v {
		arr.v.push(v)
	}
	dst.arr = arr

	for _, v := range src.v {
		dst.v.push(v)
	}
	//==

	// TODO: handle len() in interfaces
	// lastIdxElt := len(elt) - 1

	for _, v := range elt {
		if /*i == lastIdxElt &&*/ Array.isArray(v) { // The last field could be an ellipsis
			for _, vArr := range v {
				dst.v.push(vArr)
				if dst.len == dst.cap {
					dst.cap = dst.len * 2
				}
				dst.len++
			}
			break
		}

		dst.v.push(v)
		if dst.len == dst.cap {
			dst.cap = dst.len * 2
		}
		dst.len++
	}
	return dst
}

// Copy implements the function "copy".
func Copy(dst []interface{}, src interface{}) (n int) {
	// []T <= []T
	if src.arr != nil {
		for i := src.low; i < src.high; i++ {
			if n == dst.len {
				return
			}
			dst.arr.v[n] = src.arr.v[i]
			n++
		}
		for _, v := range src.v {
			if n == dst.len {
				return
			}
			dst.v.push(v)
			n++
		}
		return
	}

	// []byte <= string
	for ; n < len(src); n++ {
		if n == dst.len {
			break
		}
		dst.arr.v[n] = src[n]
	}
	return
}

// == Map
//

// The length into a map is rarely used so, in JavaScript, I prefer to calculate
// the length instead of use a field.
//
// A map has not built-in function "cap".

// MapType represents a map type.
// The compiler adds the appropriate zero value for the map (which it is work out
// from the map type).
type MapType struct {
	v    map[interface{}]interface{} // map's value
	zero interface{}                 // zero value for the map's value
}

// len returns the number of elements.
func (m MapType) len() int {
	len := 0
	for key, _ := range m.v {
		if m.v.hasOwnProperty(key) {
			len++
		}
	}
	return len
}

// typ returns the type.
func (m MapType) typ() int { return mapT }

// Map creates a map storing its zero value.
func Map(zero interface{}, v map[interface{}]interface{}) *MapType {
	m := &MapType{v, zero}
	return m
}

// get returns the value for the key "k" if it exists and a boolean indicating it.
// If looking some key up in M's map gets you "nil" ("undefined" in JS),
// then return a copy of the zero value.
func (m MapType) get(k interface{}) (interface{}, bool) {
	v := m.v

	// Allow multi-dimensional index (separated by commas)
	for i := 0; i < len(arguments); i++ {
		v = v[arguments[i]]
	}

	if v == nil {
		return m.zero, false
	}
	return v, true
}

// == Utility
//

// Export adds public names from "exported" to the map "pkg".
func Export(pkg map[interface{}]interface{}, exported []interface{}) {
	for _, v := range exported {
		pkg.v = v
	}
}

/*func Len(v interface{}) {
	
}*/
