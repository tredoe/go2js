












var PASS = true;

var x = 10;

(function() {
	x = 13;
}());

function _init() {
	if (x == 13) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: got " + x + ", want 13<br>");
		PASS = false;
	}
}

function singleLine() { document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>"); }

function simpleFunc() {
	var pass = true;


	var max = function(a, b) {
		if (a > b) {
			return a;
		}
		return b;
	};

	var x = 3;
	var y = 4;
	var z = 5;

	var max_xy = max(x, y);
	if (max_xy != 4) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: max(x,y) => got " + max_xy + ", want 4)<br>");
		pass = false, PASS = false;
	}

	var max_xz = max(x, z);
	if (max_xz != 5) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: max(x,z) => got " + max_xz + ", want 5)<br>");
		pass = false, PASS = false;
	}

	if (max(y, z) != 5) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: max(y,z) => got " + max(y, z) + ", want 5)<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function twoOuputValues() {
	var pass = true;


	var SumAndProduct = function(A, B) {
		return [A + B, A * B];
	};

	var x = 3;
	var y = 4;
	var _ = SumAndProduct(x, y), xPLUSy = _[0], xTIMESy = _[1];

	if (xPLUSy != 7) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: sum => got " + xPLUSy + ", want 7)<br>");
		pass = false, PASS = false;
	}
	if (xTIMESy != 12) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: product => got " + xTIMESy + ", want 12)<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function resultVariable() {
	var pass = true;



	var MySqrt = function(f) { var s = 0, ok = false;
		if (f > 0) {
			s = Math.sqrt(f), ok = true;
		}
		return [s, ok];
	};

	var tests = g.Map(0, {
		1: 1,
		2: 1.4142135623730951,
		3: 1.7320508075688772,
		4: 2,
		5: 2.23606797749979,
		6: 2.449489742783178,
		7: 2.6457513110645907,
		8: 2.8284271247461903,
		9: 3,
		10: 3.1622776601683795
	});

	for (var i = -2.0; i <= 10; i++) {
		var _ = MySqrt(i), sqroot = _[0], ok = _[1];
		if (ok) {
			if (JSON.stringify(sqroot) != JSON.stringify(tests.get(i)[0])) {
				document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: square(" + i + ") => got " + sqroot + ", want " + tests.get(i)[0] + "<br>");

				pass = false, PASS = false;
			}
		} else {
			if (i != -2.0 && i != -1.0 && i != 0) {
				document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: square(" + i + ") => should no be run<br>");
				pass = false, PASS = false;
			}
		}
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function _return() {
	var pass = true;

	var MySqrt = function(f) { var squareroot = 0, ok = false;
		if (f > 0) {
			squareroot = Math.sqrt(f), ok = true;
		}
		return [squareroot, ok];
	};

	var ok = MySqrt(5)[1];
	if (!ok) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: MySqrt(5) => got " + ok + ", want " + !ok + "<br>");
		pass = false, PASS = false;
	}

	var ok = MySqrt(0)[1]; if (ok) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: MySqrt(0) => got " + ok + ", want " + !ok + "<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function variadic() {
	var pass = true;

	function person(name, age) {
		this.name=name;
		this.age=age
	}



	var getOlder = function() { var people = arguments;
		if (people.length == 0) {
			return [new person(), false];
		}

		var older = people[0];

		var value; for (var _ in people) { value = people[_];
			if (value.age > older.age) {
				older = value;
			}
		}
		return [older, true];
	};

	
	var ok = false;
	var older = new person("", 0);



	var paul = new person("Paul", 23);
	var jim = new person("Jim", 24);
	var sam = new person("Sam", 84);
	var rob = new person("Rob", 54);
	var karl = new person("Karl", 19);

	var _ = function(msg, out) { return {
		msg: msg,
		out: out
	};}; var tests = [
		_("paul,jim", "Jim"),
		_("paul,jim,sam", "Sam"),
		_("paul,jim,sam,rob", "Sam"),
		_("karl", "Karl")
	];

	older = getOlder(paul, jim)[0];
	if (JSON.stringify(older.name) != JSON.stringify(tests[0].out)) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: (getOlder " + tests[0].msg + ") => got " + older.name + ", want " + tests[0].out + "<br>");

		pass = false, PASS = false;
	}

	older = getOlder(paul, jim, sam)[0];
	if (JSON.stringify(older.name) != JSON.stringify(tests[1].out)) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: (getOlder " + tests[1].msg + ") => got " + older.name + ", want " + tests[1].out + "<br>");

		pass = false, PASS = false;
	}

	older = getOlder(paul, jim, sam, rob)[0];
	if (JSON.stringify(older.name) != JSON.stringify(tests[2].out)) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: (getOlder " + tests[2].msg + ") => got " + older.name + ", want " + tests[2].out + "<br>");

		pass = false, PASS = false;
	}

	older = getOlder(karl)[0];
	if (JSON.stringify(older.name) != JSON.stringify(tests[3].out)) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: (getOlder " + tests[3].msg + ") => got " + older.name + ", want " + tests[3].out + "<br>");

		pass = false, PASS = false;
	}


	_ = getOlder(), older = _[0], ok = _[1];
	if (ok) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: (getOlder) => got " + ok + ", want " + !ok + "<br>");
		pass = false, PASS = false;
	}



	var getUser = function(name, surname, age) { var email = [].slice.call(arguments).slice(3);
		var emails = "";
		var v; for (var _ in email) { v = email[_];
			emails += " " + v;
		}
		return name + " " + surname + ", age " + age + ", emails:" + emails;
	};

	var name = "John";
	var surname = "Smith";
	var age = 17;
	var email1 = "foo@mail.se";
	var email2 = "bar@mail.se";

	var dataUser = getUser(name, surname, age, email1, email2);
	if (JSON.stringify(dataUser) != JSON.stringify(name + " " + surname + ", age " + age + ", emails: " + email1 + " " + email2)) {

		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: multiple parameters => got " + dataUser + "<br>");
		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function Max(slice) {
	if (slice.len == 1) {
		return slice.get()[0];
	}

	var middle = slice.len / 2;
	var m1 = Max(g.SliceFrom(slice, 0, middle));
	var m2 = Max(g.SliceFrom(slice, middle));

	if (m1 > m2) {
		return m1;
	}
	return m2;
}

function Invert(slice) {
	var length = slice.len;
	if (length > 1) {
		slice.set([0], slice.get()[length - 1]), slice.set([length - 1], slice.get()[0]);
		Invert(g.SliceFrom(slice, 1, length - 1));
	}
}

function recursive() {
	var pass = true;

	var s = g.Slice(0, [1, 2, 3, 4, 6, 8]);

	if (Max(s) != 8) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: Max => got " + Max(s) + ", want 8<br>");
		pass = false, PASS = false;
	}

	var slice = g.Slice(0, ['1', '2', '3', '4', '5']);
	Invert(slice);







	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function A() {
	document.write("&nbsp;&nbsp;&nbsp;&nbsp;Running function A<br>");
}

function B(name) {
	document.write("&nbsp;&nbsp;&nbsp;&nbsp;Running function " + name + "<br>");
}















function main() {
	document.write("<br><br>== Functions<br><br>");

	document.write("=== RUN init<br>");
	_init();
	document.write("=== RUN singleLine<br>");
	singleLine();
	document.write("=== RUN simpleFunc<br>");
	simpleFunc();
	document.write("=== RUN twoOuputValues<br>");
	twoOuputValues();
	document.write("=== RUN resultVariable<br>");
	resultVariable();
	document.write("=== RUN return<br>");
	_return();
	document.write("=== RUN variadic<br>");
	variadic();
	document.write("=== RUN recursive<br>");
	recursive();



	if (PASS) {
		document.write("PASS<br>");
	} else {
		document.write("FAIL<br>");
		alert("Fail: Functions");
	}

	throw new Error("unreachable");
	throw new Error("not implemented: " + "foo");
} main();
/* Generated by Go2js (github.com/tredoe/go2js) */
