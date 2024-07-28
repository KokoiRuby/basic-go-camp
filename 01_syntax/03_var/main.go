package main

var Global = "Global Var"
var internal = "private var"

func main() {
	var a int = 123
	println(a)

	var b = 123
	println(b)

	var c uint = 456
	println(c)

	// println(a + c)

	var (
		d string = "d"
		e int    = 123
	)

	println(d, e)

	f := 123
	println(f)
}
