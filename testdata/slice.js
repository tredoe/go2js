









var PASS = true;

function builtIn() {
	var pass = true;

	var s1 = g.MkSlice();
	var s2 = g.MkSlice(0, 0);
	var s3 = g.MkSlice(0, 0);
	var s4 = g.MkSlice(0, 0, 10);
	var s5 = g.Slice(0, [1, 3, 5]);

	var _ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; var tests = [
		_("nil s1", s1.isNil(), true),
		_("nil s2", s2.isNil(), false),
		_("nil s3", s3.isNil(), false),
		_("nil s4", s4.isNil(), false),
		_("nil s5", s5.isNil(), false),
		_("nil s5", !s5.isNil(), true),

		_("len s1", s1.len == 0, true),
		_("len s2", s2.len == 0, true),
		_("len s3", s3.len == 0, true),
		_("len s4", s4.len == 0, true),
		_("len s5", s5.len == 3, true),

		_("cap s1", s1.cap == 0, true),
		_("cap s2", s2.cap == 0, true),
		_("cap s3", s3.cap == 0, true),
		_("cap s4", s4.cap == 10, true),
		_("cap s5", s5.cap == 3, true)
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.in_) != JSON.stringify(t.out)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: " + t.msg + " => got " + t.in_ + ", want " + t.out + "<br>");
			pass = false, PASS = false;
		}
	}
	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function shortHand() {
	var pass = true;

	var array = g.MkArray([10], 0, ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j']);
	var a_slice = g.MkSlice(), b_slice = g.MkSlice();



	a_slice = g.SliceFrom(array, 4, 8);
	if (a_slice.str() == "efgh" && a_slice.len == 4 && a_slice.cap == 6) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [4:8] => got " + a_slice.get() + ", len=" + a_slice.len + ", cap=" + a_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	a_slice = g.SliceFrom(array, 6, 7);
	if (a_slice.str() != "g") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [6:7] => got " + a_slice.get() + "<br>");
		pass = false, PASS = false;
	}

	a_slice = g.SliceFrom(array, 0, 3);
	if (a_slice.str() == "abc" && a_slice.len == 3 && a_slice.cap == 10) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [:3] => got " + a_slice.get() + ", len=" + a_slice.len + ", cap=" + a_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	a_slice = g.SliceFrom(array, 5);
	if (a_slice.str() == "fghij" && a_slice.len == 5 && a_slice.cap == 5) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [5:] => got " + a_slice.get() + ", len=" + a_slice.len + ", cap=" + a_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	a_slice = g.SliceFrom(array, 0);
	if (a_slice.str() == "abcdefghij" && a_slice.len == 10 && a_slice.cap == 10) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [:] => got " + a_slice.get() + ", len=" + a_slice.len + ", cap=" + a_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	a_slice = g.SliceFrom(array, 3, 7);
	if (a_slice.str() == "defg" && a_slice.len == 4 && a_slice.cap == 7) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. [3:7] => got " + a_slice.get() + ", len=" + a_slice.len + ", cap=" + a_slice.cap + "<br>");

		pass = false, PASS = false;
	}



	b_slice = g.SliceFrom(a_slice, 1, 3);
	if (b_slice.str() == "ef" && b_slice.len == 2 && b_slice.cap == 6) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. [1:3] => got " + b_slice.get() + ", len=" + b_slice.len + ", cap=" + b_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	b_slice = g.SliceFrom(a_slice, 0, 3);
	if (b_slice.str() == "def" && b_slice.len == 3 && b_slice.cap == 7) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. [:3] => got " + b_slice.get() + ", len=" + b_slice.len + ", cap=" + b_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	b_slice = g.SliceFrom(a_slice, 0);
	if (b_slice.str() == "defg" && b_slice.len == 4 && b_slice.cap == 7) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. [:] => got " + b_slice.get() + ", len=" + b_slice.len + ", cap=" + b_slice.cap + "<br>");

		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function useFunc() {
	var pass = true;


	var Max = function(slice) {
		var max = slice.get()[0];
		for (var index = 1; index < slice.len; index++) {
			if (slice.get()[index] > max) {
				max = slice.get()[index];
			}
		}
		return max;
	};

	var A1 = g.MkArray([10], 0, [1, 2, 3, 4, 5, 6, 7, 8, 9]);
	var A2 = g.MkArray([4], 0, [1, 2, 3, 4]);
	var A3 = g.MkArray([1], 0, [1]);

	var slice = g.MkSlice();

	slice = g.SliceFrom(A1, 0);
	if (Max(slice) != 9) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: A1 => got " + Max(slice) + ", want 9<br>");
		pass = false, PASS = false;
	}

	slice = g.SliceFrom(A2, 0);
	if (Max(slice) != 4) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: A2 => got " + Max(slice) + ", want 4<br>");
		pass = false, PASS = false;
	}

	slice = g.SliceFrom(A3, 0);
	if (Max(slice) != 1) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: A3 => got " + Max(slice) + ", want 1<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function reference() {
	var pass = true;

	var A = g.MkArray([10], 0, ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j']);
	var slice1 = g.SliceFrom(A, 3, 7);
	var slice2 = g.SliceFrom(A, 5);
	var slice3 = g.SliceFrom(slice1, 0, 2);



	var _ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; var tests = [
		_("A", g.SliceFrom(A, 0).str(), "abcdefghij"),
		_("slice1", slice1.str(), "defg"),
		_("slice2", slice2.str(), "fghij"),
		_("slice3", slice3.str(), "de")
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.in_) != JSON.stringify(t.out)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. " + t.msg + " => got " + t.in_ + ", want " + t.out + "<br>");
			pass = false, PASS = false;
		}
	}


	A.v[4] = 'E';

	_ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; tests = [
		_("A", g.SliceFrom(A, 0).str(), "abcdEfghij"),
		_("slice1", slice1.str(), "dEfg"),
		_("slice2", slice2.str(), "fghij"),
		_("slice3", slice3.str(), "dE")
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.in_) != JSON.stringify(t.out)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. " + t.msg + " => got " + t.in_ + ", want " + t.out + "<br>");
			pass = false, PASS = false;
		}
	}


	slice2.set([1], 'G');

	_ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; tests = [
		_("A", g.SliceFrom(A, 0).str(), "abcdEfGhij"),
		_("slice1", slice1.str(), "dEfG"),
		_("slice2", slice2.str(), "fGhij"),
		_("slice3", slice3.str(), "dE")
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.in_) != JSON.stringify(t.out)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 3. " + t.msg + " => got " + t.in_ + ", want " + t.out + "<br>");
			pass = false, PASS = false;
		}
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function resize() {
	var pass = true;

	var slice = g.MkSlice();


	slice = g.MkSlice(0, 4, 5);

	if (slice.len == 4 && slice.cap == 5 && slice.get()[0] == 0 && slice.get()[1] == 0 && slice.get()[2] == 0 && slice.get()[3] == 0) {


	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. got " + slice.get() + ", want [0 0 0 0]<br>");
		pass = false, PASS = false;
	}


	slice.set([1], 2), slice.set([3], 3);

	if (slice.get()[0] == 0 && slice.get()[1] == 2 && slice.get()[2] == 0 && slice.get()[3] == 3) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. got " + slice.get() + ", want [0 2 0 3]<br>");
		pass = false, PASS = false;
	}


	slice = g.MkSlice(0, 2);

	if (slice.len == 2 && slice.cap == 2 && slice.get()[0] == 0 && slice.get()[1] == 0) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 3. got " + slice.get() + ", want [0 0]<br>");
		pass = false, PASS = false;
	}


	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function grow() {
	var pass = true;


	var GrowIntSlice = function(slice, add) {
		var new_capacity = slice.cap + add;
		var new_slice = g.MkSlice(0, slice.len, new_capacity);
		for (var index = 0; index < slice.len; index++) {
			new_slice.set([index], slice.get()[index]);
		}
		return new_slice;
	};

	var slice = g.Slice(0, [0, 1, 2, 3]);


	if (slice.len == 4 && slice.cap == 4 && slice.get()[0] == 0 && slice.get()[1] == 1 && slice.get()[2] == 2 && slice.get()[3] == 3) {


	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. got " + slice.get() + ", want [0 1 2 3]<br>");
		pass = false, PASS = false;
	}


	slice = GrowIntSlice(slice, 3);

	if (slice.len == 4 && slice.cap == 7 && slice.get()[0] == 0 && slice.get()[1] == 1 && slice.get()[2] == 2 && slice.get()[3] == 3) {


	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. got " + slice.get() + ", want [0 1 2 3]<br>");
		pass = false, PASS = false;
	}





	slice = g.SliceFrom(slice, 0, slice.len + 2);
	slice.set([4], 4), slice.set([5], 5);

	if (slice.len == 6 && slice.cap == 7 && slice.get()[0] == 0 && slice.get()[1] == 1 && slice.get()[2] == 2 && slice.get()[3] == 3 && slice.get()[4] == 4 && slice.get()[5] == 5) {



	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 3. got " + slice.get() + ", want [0 1 2 3 4 5]<br>");
		pass = false, PASS = false;
	}


	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function _copy() {
	var pass = true;

	var a = g.MkArray([8], 0, ['0', '1', '2', '3', '4', '5', '6', '7']);
	var s = g.MkSlice(0, 6);
	var b = g.MkSlice(0, 5);

	var n1 = g.Copy(s, g.SliceFrom(a, 0));
	if (s.str() == "012345" && n1 == 6) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. => got " + s.str() + ", n=" + n1 + "<br>");
		pass = false, PASS = false;
	}

	var n2 = g.Copy(s, g.SliceFrom(s, 2));
	if (s.str() == "234545" && n2 == 4) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. => got " + s.str() + ", n=" + n2 + "<br>");
		pass = false, PASS = false;
	}

	var n3 = g.Copy(b, "Hello, World!");
	if (b.str() == "Hello" && n3 == 5) {

	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 3. => got " + b.str() + ", n=" + n3 + "<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function _append() {
	var pass = true;

	var slice = g.Slice(0, ['1', '2', '3']);
	if (slice.str() == "123" && slice.len == 3) {
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 1. => got " + slice.str() + ", len=" + slice.len + "<br>");
		pass = false, PASS = false;
	}

	slice = g.Append(slice, '4');
	if (slice.str() == "1234" && slice.len == 4) {
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 2. => got " + slice.str() + ", len=" + slice.len + "<br>");
		pass = false, PASS = false;
	}

	slice = g.Append(slice, '5', '6');
	if (slice.str() == "123456" && slice.len == 6) {
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 3. => got " + slice.str() + ", len=" + slice.len + "<br>");
		pass = false, PASS = false;
	}

	slice = g.Append(slice, '7', '8', '9');
	if (slice.str() == "123456789" && slice.len == 9) {
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: 4. => got " + slice.str() + ", len=" + slice.len + "<br>");
		pass = false, PASS = false;
	}



	var a_slice = g.Slice(0, ['1', '2', '3']);
	var b_slice = g.Slice(0, ['7', '8', '9']);

	a_slice = g.Append(a_slice, b_slice.get());
	if (a_slice.str() == "123789" && a_slice.len == 6) {
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: append a slice => got " + a_slice.str() + ", len=" + a_slice.len + "<br>");

		pass = false, PASS = false;
	}













































































	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function main() {
	document.write("<br><br>== Slices<br><br>");

	document.write("=== RUN builtIn<br>");
	builtIn();
	document.write("=== RUN shortHand<br>");
	shortHand();
	document.write("=== RUN useFunc<br>");
	useFunc();
	document.write("=== RUN reference<br>");
	reference();
	document.write("=== RUN resize<br>");
	resize();
	document.write("=== RUN grow<br>");
	grow();
	document.write("=== RUN copy<br>");
	_copy();
	document.write("=== RUN append<br>");
	_append();

	if (PASS) {
		document.write("PASS<br>");
	} else {
		document.write("FAIL<br>");
		alert("Fail: Slices");
	}
} main();
/* Generated by Go2js (github.com/tredoe/go2js) */
