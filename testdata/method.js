












var PASS = true;

function Rectangle(width, height) {
	this.width=width; this.height=height
}

function noMethod() {
	var pass = true;

	var area = function(r) {
		return r.width * r.height;
	};

	var r1 = new Rectangle(12, 2);

	if (area(r1) != 24) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: area r1 => got " + area(r1) + ", want 24)<br>");
		pass = false, PASS = false;
	}
	if (area(new Rectangle(9, 4)) != 36) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: area Rectangle{9,4} => got " + area(new Rectangle(9, 4)) + ", want 36)<br>");

		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}



Rectangle.prototype.area = function() {
	return this.width * this.height;
}

function Circle(radius) {
	this.radius=radius
}

Circle.prototype.area = function() {
	return this.radius * this.radius * Math.PI;
}

function method() {
	var pass = true;

	var r1 = new Rectangle(12, 2);
	var r2 = new Rectangle(9, 4);
	var c1 = new Circle(10);
	var c2 = new Circle(25);

	var _ = function(msg, in_, out) { return {
		msg: msg,
		in_: in_,
		out: out
	};}; var tests = [
		_("Rectangle{12,2}", r1.area(), 24),
		_("Rectangle{9,4}", r2.area(), 36),
		_("Circle{10}", c1.area(), 314.1592653589793),
		_("Circle{25}", c2.area(), 1963.4954084936207)
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



function sliceOfints(){} sliceOfints.alias(g.SliceType);
function agesByNames(){} agesByNames.alias(g.MapType);

sliceOfints.prototype.sum = function() {
	var sum = 0;
	var value; for (var _ in this.t) { value = this.t[_];
		sum += value;
	}
	return sum;
}

agesByNames.prototype.older = function() {
	var a = 0;
	var n = "";
	var value; for (var key in this.t) { value = this.t[key];
		if (value > a) {
			a = value;
			n = key;
		}
	}
	return n;
}

function withNamedType() {
	var pass = true;

	var s = new sliceOfints(1, 2, 3, 4, 5);
	var folks = new agesByNames();
	folks["Bob"] = 36,
	folks["Mike"] = 44,
	folks["Jane"] = 30,
	folks["Popey"] = 100;


	if (s.sum() != 15) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: s.sum => got " + s.sum() + ", want 15)<br>");
		pass = false, PASS = false;
	}
	if (folks.older() != "Popey") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: folks.older => got " + folks.older() + ", want Popey)<br>");

		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}



const 
WHITE = 0,
BLACK = 1,
BLUE = 2,
RED = 3,
YELLOW = 4;


function Color(t) { this.t=t; }

function Box(width, height, depth, color) {
	this.width=width; this.height=height; this.depth=depth;
	this.color=color
}

function BoxList(){} BoxList.alias(g.SliceType);

Box.prototype.Volume = function() {
	return this.width * this.height * this.depth;
}

Box.prototype.SetColor = function(c) {
	this.color = c;
}

BoxList.prototype.BiggestsColor = function() {
	var v = 0.00;
	var k = Color(WHITE);
	var b; for (var _ in this.t) { b = this.t[_];
		if (b.Volume() > v) {
			v = b.Volume();
			k = b.color;
		}
	}
	return k;
}

BoxList.prototype.PaintItBlack = function() {
	var _; for (var i in this.t) { _ = this.t[i];
		this.t[i].SetColor(BLACK);
	}
}

Color.prototype.String = function() {
	var strings = g.Slice("", ["WHITE", "BLACK", "BLUE", "RED", "YELLOW"]);
	return strings.get()[this.t];
}

function complexNamedType() {
	var pass = true;

	var boxes = new BoxList(
		new Box(4, 4, 4, RED),
		new Box(10, 10, 1, YELLOW),
		new Box(1, 1, 20, BLACK),
		new Box(10, 10, 1, BLUE),
		new Box(20, 20, 20, YELLOW),
		new Box(10, 30, 1, WHITE)
	);

	if (boxes.length != 6) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: len boxes => got " + boxes.length + ", want 6<br>");
		pass = false, PASS = false;
	}
	if (boxes[0].Volume() != 64) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: the volume of the first one => got " + boxes[0].Volume() + ", want 64<br>");

		pass = false, PASS = false;
	}
	if (boxes[boxes.length - 1].color.String() != "WHITE") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: the color of the last one => got " + boxes[boxes.length - 1].color.String() + ", want WHITE<br>");

		pass = false, PASS = false;
	}
	if (boxes.BiggestsColor().String() != "YELLOW") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: the biggest one => got " + boxes.BiggestsColor().String() + ", want YELLOW<br>");

		pass = false, PASS = false;
	}


	boxes.PaintItBlack();

	if (boxes[1].color.String() != "BLACK") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: the color of the second one => got " + boxes[1].color.String() + ", want BLACK<br>");

		pass = false, PASS = false;
	}
	if (boxes.BiggestsColor().String() != "BLACK") {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;FAIL: finally, the biggest one => got " + boxes.BiggestsColor().String() + ", want BLACK<br>");

		pass = false, PASS = false;
	}

	if (pass) {
		document.write("&nbsp;&nbsp;&nbsp;&nbsp;pass<br>");
	}
}



function main() {
	document.write("<br><br>== Methods<br><br>");

	document.write("=== RUN noMethod<br>");
	noMethod();
	document.write("=== RUN method<br>");
	method();
	document.write("=== RUN withNamedType<br>");
	withNamedType();
	document.write("=== RUN complexNamedType<br>");
	complexNamedType();

	if (PASS) {
		document.write("PASS<br>");
	} else {
		document.write("FAIL<br>");
		alert("Fail: Methods");
	}
} main();
/* Generated by Go2js (github.com/tredoe/go2js) */
