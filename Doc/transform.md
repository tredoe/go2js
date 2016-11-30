Transforming to JavaScript
==========================

## TODO

1. Do the type inference. There is some (incomplete) typechecker next to Go
parser. Type inference seems not to be complex, but it is quite tricky,
especially with << (see shift2.go in test directory).

2. Give types to the polymorphic functions (infix +, etc). Note that some
of arithmetic functions (+ for strings, <<, >>) may works not like in C or
your target lang, thus it would be better to convert them to functions
(__shl_32_32, __plus_str, ...) at the AST level.

3. Resolve consts. That implies infinite-precision arithmetic with complex
numbers. There are also tricky issues with type casting inside the const
declarations.

4. Resove the syntax which cannot be emited as-is to C or javascript, for
example { x := 1; { x := x + 1; /* two different x'es in a expression */ } }

5. Remove syntax sugar from AST to make further processing simpler (It may
be worth to define "Kernel Go" (same way as Kernel Mozart/Oz), which is
still compilable Go but has minimal number of construct).

 - get rid of constructions with duplicated semantic (&([100]int{}) and
new([100]int), make([]int, 50, 100) and new([100]int)[:50])

 - get rid of :=, "x := 1", "var x int = 1", "var x = 1" to be the same [OK]

 - desugar multielement assigment ( *a(),*b()=c(),d() ), see Go 1 spec for
the evaluation order.

 - desugar swap(swap(a,b))

 - implement 'switch' and 'for' via 'goto' (and get rid of
break/continue/fallthrough)

 - convert struct comparision to elementwise comparision

 - desugar named returns to local vars with uniq names

 - make zero-initialization of vars explicit [OK]

+ Substitute "var" by "let" for local variables, when browsers use ECMAScript 6

http://kishorelive.com/2011/11/22/ecmascript-6-looks-promising/

+ Should be translated empty structs? Or simply ignored?


## Initialization

The values must be initialized explicitly, else they will be "undefined".

	// Faills
	var s1;
	if (s1 === "") { alert("undefined"); }

	// It's ok
	var s2 = "";
	if (s2 === "") { alert("empty string"); }


## Numbers

All JavaScript numbers are double floats (64 bits) but 11 bits are used to store
the position of the fractional dot within the number. And one extra bit is used
with the signed numbers.  
So you can't actually do meaningful integer arithmetic on anything bigger than
2^53.

	// Maximum size for integers in JavaScript.
	MAX_UINT_JS = 1<<53 - 1
	MAX_INT_JS  = 1<<52 - 1

http://www.jwz.org/blog/2010/10/every-day-i-learn-something-new-and-stupid/
http://rx4ajax-jscore.com/ecmacore/datatype/number.html


One soluction would be to checking if every integer variable fills in that
capacity, but the effort is not worth the work since there would be to checking
every binary operation related to it.  
The best option is to use a JavaScript library which handles integers of 64 bits.

https://github.com/jtobey/javascript-bignum


## Struct

> A struct in an object-oriented environment is a "public class". Classes
> represent objects (they become objects when instantiated) composed of
> properties and methods. Properties of classes probably have a direct
> correspondence to members of a struct type.

> While there are no classes in JS, you effectively create them when you
> create the constructor for an object in JS, in which you can assign names
> of properties and methods within the constructor, and then populate them
> with their values through arguments or constants/other assignments within
> the function block of the constructor. You would instantiate 'structs' or
> 'classes' in this manner.

[Reference](http://bytes.com/topic/javascript/answers/441203-structs-javascript)


## Map

It is defined using an object.

	`var m1 = new Object();` or `var m1 = {};`

	m1['one']='first';
	m1['two']='second';
	m1['three']='third';

To create a map with values initialized:

	var m2 = {
		1: 'Joe',
		3: 'Sam',
		8: 'Eve',
	};

+ Loop: `for(var i in m1)`


## Testing

The JavaScript output of the files in directory "test" have been checked using
[JavaScript Lint](http://javascriptlint.com/download.htm):

	wget http://javascriptlint.com/download/jsl-0.3.0-src.tar.gz
	tar xzf jsl-*.tar.gz && cd jsl-*/src && make -f Makefile.ref BUILD_OPT=1 &&
	sudo cp Linux_All_*.OBJ/jsl /usr/local/bin/ && cd - && rm -rf jsl-*

	jsl -process test/var.js


## HTML5

+ [Wrapper to Filesystem API]
(http://ericbidelman.tumblr.com/post/14866798359/introducing-filer-js)

