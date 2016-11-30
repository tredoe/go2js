// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

/*
Command Go2js translates Go into JavaScript so you can continue using a
clean and concise sintaxis.

Really, it is used a subset of Go since JavaScript has not native way to
represent some types neither Go's statements, although some of them could be
emulated.

Advantages:

+ Using one only language for all development. A great advantage for a company.

+ Allows many type errors to be caught early in the development cycle, due to
static typing. (ToDo: compile to checking errors at time of compiling)

+ The mathematical expressions in the constants are calculated at the
translation stage. (ToDo)

+ The lines numbers in the un-minified generated JavaScript match up with the
lines numbers in the original source file.

+ Generates minimized JavaScript.

Go sintaxis not supported:

+ Complex numbers, integers of 64 bits.
+ Function type, interface type excepting the empty interface.
+ Channels, goroutines (could be translated to Web Workers (http://www.html5rocks.com/en/tutorials/workers/basics/).
+ Built-in function recover.
+ Defer statement.
+ Goto, labels. (1) In JavaScript, the labels are restricted to "for" and
"while" loops when they are called from "continue" and "break" directives so
its use is very limited, and (2) it is advised to avoid its use
(https://developer.mozilla.org/en/JavaScript/Reference/Statements/label#Avoid_using_labels).

Note: JavaScript can not actually do meaningful integer arithmetic on anything
bigger than 2^53. Also bitwise logical operations only have defined results (per
the spec) up to 32 bits.  
By this reason, the integers of 64 bits are unsupported.


## Translation

#### Reserved words

The reserved words and keywords used in JavaScript are translated adding "_" at
the end of the name.

See files "testdata/decl_reserved.{go,js}".

#### Initialization

The values are initialized explicitly like Go does, else they would be
"undefined".

	var s1; if (s1 === undefined) { alert("s1 is of value 'undefined'"); }
	var s2 = ""; if (s2 === "") { alert("s2 is an empty string"); }

#### Pointers

There are five primitive values in JavaScript that are stored directly in the
stack: Null, Undefined, Boolean, String and Number. Objects, arrays and
everything else are reference types.  
But you do not pass by reference in JavaScript. Not ever. You always pass a
value from the stack. If the value happens to be a pointer, then it may appear
as if you were passing by reference, but in reality you aren't.  
ECMAScript is simply not able to pass by reference.

The emulation is done using an object with a field named `p`. So:

`*x` and `x` is `x.p` in javascript while `&x` would simply be `x`.

Then, for any variable that is addressed:

`var x *bool` to `var x = {p:false}`

Note: the printing of an address in Go (`&x`) results into an hexadecimal
address. Instead, in JavaScript with this emulation, it prints the value.

Warning: due to JavaScript design, it cann't be guaranteed that the
emulation of pointers can work in other scenes that are not in the test file.

#### Numbers

JavaScript doesn't have an integer division operator like some languages do, so
there is to get the number by rounding down the result:

+ Math.floor() will give the wrong result if the result of the division is a
negative number

	>>> Math.floor(-3/2)
	-2

+ For a 32-bit signed integer, can be used ">>0"

	>>> -3/2 >>0
	-1

+ Using bitwise operations:

	>>> ~~(-3/2)
	-1

	>>> -3/2 |0
	-1

#### Comparison

In JavaScript, when objects are compared then the identity is checked, no
comparison of properties or elements is done.

One way without a custom comparison function is to compare its string
representations, but using the JSON object.  
Here there is an example that shows the why:

	var a = [1, [1, 1], 1], b = [[1, 1], [1, 1]];

	console.log("String() => a: \"" + String(a) + "\" b: \"" + String(b) + "\"")
	if (String(a) === String(b)) { console.log("String(): equals"); }

	console.log("JSON() => a: \"" + JSON.stringify(a) + "\" b: \"" + JSON.stringify(b) + "\"")
	if (JSON.stringify(a) === JSON.stringify(b)) { console.log("JSON(): equals"); }

		String() => a: "1,1,1,1" b: "1,1,1,1"
		String(): equals
		JSON() => a: "[1,[1,1],1]" b: "[[1,1],[1,1]]"

#### Return of multiple values

When a Go function returns more than one value then those values are put into an
array. Then, to access to the different values it is created a variable
`_` assigned to the return of the function, and the variable's names defined in
Go are used to access to each value of that array.

By example, for a Go function like this:

	sum, product := sumAndProduct(x, y)

its translation would be:

	var _ = SumAndProduct(x, y), sum = _[0], product = _[1];

#### Library

JavaScript has several built-in functions and constants which can be translated
from Go. They are defined in the maps "Constant", and "Function".

Since the Go functions "print*" are used to debug, they are translated to
"console.error"; the functions "fmt.Print*" are translated to "console.log"

"panic" is translated to "throw new Error()".

#### Modularity

JavaScript has not some kind of module system built in. To simulate it, all the
code for the package is written inside an anonymous function which is called
directly. Then, it is used a helper function to give to an object (named like
the package) the values that must be exported.

By example, for a package named "foo" with names exported "Add" and "Product":

	var foo = {}; (function() {
	// Code of your package

	g.Export(foo, [Add, Product])
	})();


## Contributing

If you are going to change code related to the compiler then you should run
"go test" after of each change in your forked repository. It will translate the
Go files in the directory "testdata". To see the differences use "git diff",
checking whether the change in the JavaScript files is what you were expecting.  
It is also expected to get some errors and warnings in some of them, which are
validated using the test functions for examples. See file "goscript_test.go".

Then, to checking the generated JavaScript files, use the browser with the
address "file:///PATH_TO/goscript/testdata/test.html".

Ideas:

+ Implement the new JS API for HTML5, translating it from Go functions. See
 both maps *Constant* and *Function* in file "library.go". But you must be sure
 that the API is already implemented in both browsers Firefox and Chrome.
+ The Dart library (http://api.dartlang.org/) could be used like inspiration to
 write web libraries, especially "dom" and "html".
+ JavaScript library to handle integers of 64 bits. Build it in Go since it
 can be translated to JS; see "jslib/lib.go".


## Vision

The great problem of the web developing is that there is one specification and
multiple implementations with the result of that different browsers have
different implementations for both DOM and JavaScript.  
So the solution is obvious, one specification and one implementation; and here
comes Go.

References to develop a DOM library implemented:

	https://developer.mozilla.org/en/Gecko_DOM_Reference  
	http://api.dartlang.org/dom.html  
	http://api.dartlang.org/html.html


## Credits

The tests are basedd in the examples of Big Yuuta's book for novices
(http://go-book.appspot.com/).
*/
package main
