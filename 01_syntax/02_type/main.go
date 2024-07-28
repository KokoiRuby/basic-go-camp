package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	a := 456
	b := 123
	println(a + b)
	println(a - b)
	println(a * b)
	if b != 0 { // avoid panic
		println(a / b)
	}
	println(a % b)

	var c float64 = 12.3
	//println(a + c)
	println(a + int(c))

	var d int32 = 12
	// println(a + d)
	println(a + int(d))

	// extreme
	println(math.MaxInt)
	println(math.MinInt)
	println(math.MaxFloat32)
	println(math.SmallestNonzeroFloat32)
	println(math.MaxFloat64)
	println(math.SmallestNonzeroFloat64)

	// string
	println("this is a string")
	println(`this
					is 
					a
					string`)
	// "Hello Go!"
	println("\"Hello Go!\"")

	println(len("abc"))                   // 3
	println(len("你好"))                    // 6
	println(utf8.RuneCountInString("你好")) // 2

	var e byte = 'a'
	println(e) // 97
	fmt.Printf("%c", e)

	var str string = "this is a string"
	var bs []byte = []byte(str)
	bs[0] = 'T'
	println(str, bs)
}
