









var PASS = true;


var i = g.Int({p:undefined});
var hello = {p:undefined};
var p = {p:undefined};

(function() {
	p = i;
	var helloPtr = hello;

	document.write("== init()<br>");
	document.write("&nbsp;&nbsp;&nbsp;&nbsp;\"helloPtr\": " + helloPtr);
}());

function declaration() {
	var i = g.Int({p:undefined});
	var hello = {p:undefined};
	var p = {p:undefined};

	p = i;
	var helloPtr = hello;
	document.write("&nbsp;&nbsp;&nbsp;&nbsp;\"p\": " + p + " " + "<br>&nbsp;&nbsp;&nbsp;&nbsp;\"helloPtr\": " + helloPtr + "<br>");
}

function showAddress() {
	
	var i = g.Int({p:9});
	var hello = {p:"Hello world"};
	var pi = g.Float32({p:3.14});
	var b = {p:true};


	document.write("&nbsp;&nbsp;&nbsp;&nbsp;\"i\": " + i + "<br>");
	document.write("&nbsp;&nbsp;&nbsp;&nbsp;\"hello\": " + hello + "<br>");
	document.write("&nbsp;&nbsp;&nbsp;&nbsp;\"pi\": " + pi + "<br>");
	document.write("&nbsp;&nbsp;&nbsp;&nbsp;\"b\": " + b + "<br>");
}

function nilValue() {
	var pass = true;

	var num = {p:10};
	var p = {p:undefined};

	if (p.p == undefined) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: declaration => got " + p == undefined + "<br>");
		pass = false, PASS = false;
	}

	p = num;
	if (p.p != undefined) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: assignment => got " + p == undefined + "<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function access() {
	var pass = true;

	var hello = {p:"Hello, mina-san!"};
	var helloPtr = {p:undefined};
	helloPtr = hello;

	var i = {p:6};
	var iPtr = i;

	if (helloPtr.p != "Hello, mina-san!") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: *helloPtr => got " + helloPtr.p + ", want " + hello + "<br>");
		pass = false, PASS = false;
	}
	if (iPtr.p != 6) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: *iPtr => got " + iPtr.p + ", want " + i + "<br>");
		pass = false, PASS = false;
	}



	var x = {p:3};
	var y = x;

	y.p++;
	if (x.p != 4) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: x => got " + x + ", want 4<br>");
		pass = false, PASS = false;
	}

	y.p++;
	if (x.p != 5) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: x => got " + x + ", want 5<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function allocation() {
	var sum = 0;
	var doubleSum = {p:undefined};
	for (var i = 0; i < 10; i++) {
		sum += i;
	}

	doubleSum.p = 0;
	doubleSum.p = sum * 2;

	if (sum == 45 && doubleSum.p == 90) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: sum=" + sum + ", *doubleSum=" + doubleSum.p + "<br>");
		PASS = false;
	}
}

function parameterByValue() {

	var add = function(v) {
		v = v + 1;
		return v;
	};

	var x = 3;
	var x1 = add(x);

	if (x == 3 && x1 == 4) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: x=" + x + ", x1=" + x1 + "<br>");
		PASS = false;
	}
}

function byReference_1() {
	var pass = true;

	var add = function(v) {
		v.p = v.p + 1;
		return v.p;
	};

	var x = {p:3};
	var x1 = add(x);

	if (x1 == 4 && x.p == 4) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. x=" + x + ", x1=" + x1 + "<br>");
		pass = false, PASS = false;
	}

	x1 = add(x);
	if (x.p == 5 && x1 == 5) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. x=" + x + ", x1=" + x1 + "<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function byReference_2() {
	var pass = true;

	var add = function(v, i) { v.p += i; };
	var value = {p:6};
	var incr = 1;

	add(value, incr);
	if (value.p == 7) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. value=" + value + "<br>");
		pass = false, PASS = false;
	}

	add(value, incr);
	if (value.p == 8) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. value=" + value + "<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function byReference_3() {
	var x = {p:3};
	var f = function() {
		x.p = 4;
	};
	var y = x;

	f();
	if (y.p == 4) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: *y=" + y.p + "<br>");
		PASS = false;
	}
}

function main() {
	document.write("<br><br>== Pointers<br><br>");






	document.write("=== RUN nilValue<br>");
	nilValue();
	document.write("=== RUN access<br>");
	access();
	document.write("=== RUN allocation<br>");
	allocation();

	document.write("=== RUN parameterByValue<br>");
	parameterByValue();
	document.write("=== RUN byReference_1<br>");
	byReference_1();
	document.write("=== RUN byReference_2<br>");
	byReference_2();
	document.write("=== RUN byReference_3<br>");
	byReference_3();

	if (PASS) {
		document.write("PASS<br>");
	} else {
		document.write("FAIL<br>");
		alert("Fail: Pointers");
	}
} main();
/* Generated by Go2js (github.com/tredoe/go2js) */
