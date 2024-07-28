package main

import "fmt"

func Func1() (name string, age int) {
	return "name", 18
}

func Func2() (name string, age int) {
	name = "name"
	age = 18
	return
}

func Func3() (name string, age int) {
	// zero-value
	// return "", 0
	return
}

func A() { // circular dep
	C()
}

func B() {
	B()
}

func C() {
	A()
}

func Func4() {
	varName := Func3
	//varName := Func3() // nok
	_, _ = varName()
}

func Func5() {
	fn := func() string {
		return "hello"
	}

	fn()
}

func Func6(name string) func() string { // closure + return func
	return func() string {
		return name
	}
}

func Func7() {
	fn := func() string {
		return "hello"
	}()

	println(fn)
}

func Closure1() func() int {
	age := 0
	return func() int {
		age++
		return age
	}
}

func Func8(name string, alias ...string) {

}

func DeferClosure1() {
	i := 0
	defer func() {
		println(i) // 1
	}()
	i = 1
}

func DeferClosure2() {
	i := 0
	defer func(val int) {
		println(val) // 0
	}(i)
	i = 1
}

func DeferReturn1() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturn2() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferClosureLoop1() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("addr is: %p, value is %v\n", &i, i)
		}()
	}
}

func DeferClosureLoop2() {
	for i := 0; i < 5; i++ {
		defer func(val int) {
			fmt.Printf("addr is: %p, value is %v\n", &val, val)
		}(i)
	}
}

func DeferClosureLoop3() {
	for i := 0; i < 5; i++ {
		j := i
		defer func() {
			fmt.Printf("addr is: %p, value is %v\n", &j, j)
		}()
	}
}

func DeferClosureLoop4() {
	var j int
	for i := 0; i < 5; i++ {
		j = i
		defer func() {
			fmt.Printf("addr is: %p, value is %v\n", &j, j)
		}()
	}
}

type User struct {
	name string
}

func Loop() {
	users := []User{
		{
			name: "A",
		},
		{
			name: "B",
		},
		{
			name: "C",
		},
	}
	m := make(map[string]*User, 3)
	for _, u := range users {
		fmt.Printf("%p\n", &u)
		m[u.name] = &u
	}

	for k, v := range m {
		fmt.Printf("name: %s, user: %v\n", k, v)
	}
}

func main() {
	str1 := Func6("Hello")
	str2 := Func6("world")
	println(str1(), str2())

	getAge := Closure1()
	println(getAge()) // 1
	println(getAge()) // 2
	println(getAge()) // 3
	getAgeAgain := Closure1()
	println(getAgeAgain()) // 1
	println(getAgeAgain()) // 2
	println(getAgeAgain()) // 3

	Func8("A")
	Func8("A", "B")
	Func8("A", "B", "C")
	aliases := []string{"B", "C"}
	Func8("A", aliases...) // don't forget ...

	DeferClosure1()
	DeferClosure2()

	println(DeferReturn1())
	println(DeferReturn2())

	DeferClosureLoop1()
	DeferClosureLoop2()
	DeferClosureLoop3()
	DeferClosureLoop4()

	Loop()

}
