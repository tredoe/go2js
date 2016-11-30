









var PASS = true;

function person(name, age) {
	this.name=name;
	this.age=age
}


function older(p1, p2) {
	if (p1.age > p2.age) {
		return [p1, p1.age - p2.age];
	}
	return [p2, p2.age - p1.age];
}


function older10(people) {
	var older = people.v[0];


	for (var index = 1; index < 10; index++) {
		if (people.v[index].age > older.age) {
			older = people.v[index];
		}
	}
	return older;
}




function builtInArray() {
	var pass = true;





	var a1 = g.MkArray([5], 0);
	var a2 = g.MkArray([5], 0);
	var a3 = g.MkArray([5], 0, [2]);
	var a4 = g.MkArray([5], 0, [2, 4]);

	var a5 = g.MkArray([3,4], 0);
	var a6 = g.MkArray([3,4,2], 0);

	var _ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; var tests = [


		_("len a1", a1.len() == 5, true),
		_("len a2", a2.len() == 5, true),
		_("len a3", a3.len() == 5, true),
		_("len a4", a4.len() == 5, true),
		_("len a4", a4.len() != 5, false),

		_("cap a1", a1.cap() == 5, true),
		_("cap a2", a2.cap() == 5, true),
		_("cap a3", a3.cap() == 5, true),
		_("cap a4", a4.cap() == 5, true),

		_("len a5", a5.len() == 3, true),
		_("cap a5", a5.cap() == 3, true),
		_("len a5[0]", a5.len(0) == 4, true),
		_("cap a5[0]", a5.cap(0) == 4, true),
		_("len a5[1000]", a5.len(1000) == 4, true),
		_("cap a5[1000]", a5.cap(1000) == 4, true),

		_("len a6", a6.len() == 3, true),
		_("cap a6", a6.cap() == 3, true),
		_("len a6[0]", a6.len(0) == 4, true),
		_("cap a6[0]", a6.cap(0) == 4, true),
		_("len a6[0][0]", a6.len(0,0) == 2, true),
		_("cap a6[0][0]", a6.cap(0,0) == 2, true),
		_("len a6[0][1000]", a6.len(0,1000) == 2, true),
		_("cap a6[0][1000]", a6.cap(0,1000) == 2, true)
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

function initArray() {
	var pass = true;


	var array1 = g.MkArray([10], new person("", 0), [
		new person("", 0),
		new person("Paul", 23),
		new person("Jim", 24),
		new person("Sam", 84),
		new person("Rob", 54),
		new person("", 0),
		new person("", 0),
		new person("", 0),
		new person("Karl", 10),
		new person("", 0)
	]);


	var array2 = g.MkArray([10], new person("", 0), [
		new person("", 0),
		new person("Paul", 23),
		new person("Jim", 24),
		new person("Sam", 84),
		new person("Rob", 54),
		new person("", 0),
		new person("", 0),
		new person("", 0),
		new person("Karl", 10),
		new person("", 0)]);

	var _ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; var tests = [
		_("len", array1.len() == array2.len(), true),
		_("cap", array1.cap() == array2.cap(), true),
		_("equality", JSON.stringify(array1.v) == JSON.stringify(array2.v), true)
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

function _array() {

	var array = g.MkArray([10], new person("", 0));



	array.v[1] = new person("Paul", 23);
	array.v[2] = new person("Jim", 24);
	array.v[3] = new person("Sam", 84);
	array.v[4] = new person("Rob", 54);
	array.v[8] = new person("Karl", 19);

	var older = older10(array);

	if (older.name == "Sam") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: got " + older.name + ", want Sam<br>");
		PASS = false;
	}
}

function multiArray() {

	var doubleArray_1 = g.MkArray([2,4], 0, [[1, 2, 3, 4], [5, 6, 7, 8]]);


	var doubleArray_2 = g.MkArray([2,4], 0, [
		[1, 2, 3, 4], [5, 6, 7, 8]]);


	var doubleArray_3 = g.MkArray([2,4], 0, [
		[1, 2, 3, 4],
		[5, 6, 7, 8]
	]);

	if (JSON.stringify(doubleArray_1.v) == JSON.stringify(doubleArray_2.v) && JSON.stringify(doubleArray_2.v) == JSON.stringify(doubleArray_3.v)) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	} else {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: got different arraies<br>");
		PASS = false;
	}
}




function _struct() {
	var pass = true;

	var tom = new person("", 0);
	tom.name = "Tom", tom.age = 18;

	var bob = new person(); bob.age = 25, bob.name = "Bob";
	var paul = new person("Paul", 43);

	var _ = older(tom, bob), TB_older = _[0], TB_diff = _[1];
	var _ = older(tom, paul), TP_older = _[0], TP_diff = _[1];
	var _ = older(bob, paul), BP_older = _[0], BP_diff = _[1];

	var _ = function(msg, inPerson, outPerson, inDiff, outDiff) { return {
		msg: msg,
		inPerson: inPerson,
		outPerson: outPerson,
		inDiff: inDiff,
		outDiff: outDiff
	};}; var tests = [
		_("Tom,Bob", TB_older, bob, TB_diff, 7),
		_("Tom,Paul", TP_older, paul, TP_diff, 25),
		_("Bob,Paul", BP_older, paul, BP_diff, 18)
	];

	var t; for (var _ in tests) { t = tests[_];
		if (JSON.stringify(t.inPerson) != JSON.stringify(t.outPerson)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: " + t.msg + " => person got " + t.inPerson + ", want " + t.outPerson + "<br>");

			pass = false, PASS = false;
		}
		if (JSON.stringify(t.inDiff) != JSON.stringify(t.outDiff)) {
			document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: " + t.msg + " => difference got " + t.inDiff + ", want " + t.outDiff + "<br>");

			pass = false, PASS = false;
		}
	}
	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}

function main() {
	document.write("<br><br>== Composite types<br><br>");

	document.write("=== RUN builtInArray<br>");
	builtInArray();
	document.write("=== RUN initArray<br>");
	initArray();
	document.write("=== RUN array<br>");
	_array();
	document.write("=== RUN multiArray<br>");
	multiArray();

	document.write("=== RUN struct<br>");
	_struct();

	if (PASS) {
		document.write("PASS<br>");
	} else {
		document.write("FAIL<br>");
		alert("Fail: Composite types");
	}
} main();
/* Generated by Go2js (github.com/tredoe/go2js) */
