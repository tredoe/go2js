









var PASS = true;
var rating = g.Map(0, {"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2});

function builtIn() {
	var pass = true;

	var m1 = g.Map(0);
	var m2 = g.Map(0, {});
	var m3 = g.Map(0, {});
	var m4 = g.Map(0, {});

	var _ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; var tests = [
		_("nil m1", m1.v == undefined, true),
		_("nil m2", m2.v == undefined, false),
		_("nil m3", m3.v == undefined, false),
		_("nil m4", m4.v == undefined, false),
		_("nil m4", m4.v != undefined, true),

		_("len m1", m1.len() == 0, true),
		_("len m2", m2.len() == 0, true),
		_("len m3", m3.len() == 0, true),
		_("len m4", m4.len() == 0, true),

		_("nil rating", rating.v != undefined, true),
		_("len rating", rating.len() == 4, true)
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

function declaration() {
	var pass = true;

	var numbers = g.Map(0);
	numbers = g.Map(0, {});
	numbers.v["one"] = 1;
	numbers.v["ten"] = 10;
	numbers.v["trois"] = 3;


	var rating1 = g.Map(0, {"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2});


	var rating2 = g.Map(0, {});
	rating2.v["C"] = 5;
	rating2.v["Go"] = 4.5;
	rating2.v["Python"] = 4.5;
	rating2.v["C++"] = 2;

	var _ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; var tests = [
		_("numbers[\"one\"]", numbers.get("one")[0], 1),
		_("numbers[\"ten\"]", numbers.get("ten")[0], 10),
		_("numbers[\"trois\"]", numbers.get("trois")[0], 3),

		_("rating[\"C\"]", rating1.get("C")[0], rating2.get("C")[0]),
		_("rating[\"Go\"]", rating1.get("Go")[0], rating2.get("Go")[0]),
		_("rating[\"Python\"]", rating1.get("Python")[0], rating2.get("Python")[0]),
		_("rating[\"C++\"]", rating1.get("C++")[0], rating2.get("C++")[0])
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

function reference() {
	var m = g.Map("", {});
	m.v["Hello"] = "Bonjour";

	var m1 = m;
	m1.v["Hello"] = "Salut";

	if (JSON.stringify(m.get("Hello")[0]) == JSON.stringify(m1.get("Hello")[0])) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: m[\"Hello\"] => got " + m.get("Hello")[0] + ", want " + m1.get("Hello")[0] + "<br>");
		PASS = false;
	}
}

function keyNoExistent() {
	var pass = true;

	var csharp_rating = rating.get("C#")[0];
	var _ = rating.get("C#"), csharp_rating2 = _[0], found = _[1];

	var multiDim = g.Map(0, {1: {1: 1.1}, 2: {2: 2.2}});
	var k_multiDim = multiDim.get(1,2)[0];

	var _ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; var tests = [
		_("csharp_rating", csharp_rating, 0.00),
		_("csharp_rating2", csharp_rating2, 0),
		_("k_multiDim", k_multiDim, 0)
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.in_) != JSON.stringify(t.out)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: " + t.msg + " => got " + t.in_ + ", want " + t.out + "<br>");
			pass = false, PASS = false;
		}
	}
	if (found) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: using comma => got " + found + ", want " + !found + "<br>");
		pass = false, PASS = false;
	}
	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function deleteKey() {
	var pass = true;

	delete rating.v["C++"];
	var found = rating.get("C++")[1];

	if (found) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: got " + found + ", want " + !found + "<br>");
		pass = false, PASS = false;
	}
	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function _range() {
	var pass = true;

	var value; for (var key in rating.v) { value = rating.get(key)[0];
		switch (key) {
		case "C":
			if (value != 5) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: " + key + " => got " + value + ", want 5<br>");
			pass = false, PASS = false;
		} break;
		case "Go":
			if (value != 4.5) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: " + key + " => got " + value + ", want 4.5<br>");
			pass = false, PASS = false;
		} break;
		case "Python":
			if (value != 4.5) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: " + key + " => got " + value + ", want 4.5<br>");
			pass = false, PASS = false;
		} break;
		default:
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: " + key + " => no expected<br>");
			pass = false, PASS = false;
		}
	}


	for (var key in rating.v) {
		if (key != "C" && key != "Go" && key != "Python") {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: key " + key + " no expected<br>");
			pass = false, PASS = false;
		}
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function blankIdInRange() {
	var pass = true;


	var Max = function(slice) {
		var max = slice.get()[0];
		var value; for (var _ in slice.get()) { value = slice.get()[_];
			if (value > max) {
				max = value;
			}
		}
		return max;
	};

	var slice = g.MkSlice();

	var A1 = g.MkArray([10], 0, [1, 2, 3, 4, 5, 6, 7, 8, 9]);
	var A2 = g.MkArray([4], 0, [1, 2, 3, 4]);
	var A3 = g.MkArray([1], 0, [1]);

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

function main() {
	document.write("<br><br>== Maps<br><br>");

	document.write("=== RUN builtIn<br>");
	builtIn();
	document.write("=== RUN declaration<br>");
	declaration();
	document.write("=== RUN reference<br>");
	reference();
	document.write("=== RUN keyNoExistent<br>");
	keyNoExistent();
	document.write("=== RUN deleteKey<br>");
	deleteKey();
	document.write("=== RUN range<br>");
	_range();
	document.write("=== RUN blankIdInRange<br>");
	blankIdInRange();

	if (PASS) {
		document.write("PASS<br>");
	} else {
		document.write("FAIL<br>");
		alert("Fail: Maps");
	}
} main();
/* Generated by Go2js (github.com/tredoe/go2js) */
