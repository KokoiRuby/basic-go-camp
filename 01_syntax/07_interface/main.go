package main

import (
	"fmt"
)

// type redef
type Integer int
type aliasOfInt int

type MyInterface interface {
	Add(idx int, val any) error
	Append(val any)
}

type Node struct {
	val int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Add(idx int, val any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

type Inner struct {
}

func (i Inner) DoSomething() {

}

type Outer struct {
	Inner
}

func (o Outer) Name() string {
	return "Outer"
}

func (i Inner) SayHello() {
	fmt.Println("Hello, " + i.Name())
}

func (i Inner) Name() string {
	return "Inner"
}

// generic
// interface
type List[T any] interface {
	Add(index int, val T)
	Append(val T)
	Delete(index T)
}

// struct
type Node1[T any] struct {
	val T
}

type LinkedList1[T any] struct {
	head *Node1[T]
}

func (l *LinkedList1[T]) Add(index int, val T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList1[T]) Append(val T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList1[T]) Delete(index T) {
	//TODO implement me
	panic("implement me")
}

// method
func Sum[T Number](vals ...T) T {
	var sum T
	for _, val := range vals {
		sum = sum + val
	}
	return sum
}

// generic constraint
type Number interface {
	int | uint
}

func main() {
	i1 := 10
	i2 := Integer(i1)
	fmt.Printf("%T\n", i1)
	fmt.Printf("%T\n", i2)

	var o Outer
	// same
	o.DoSomething()
	o.Inner.DoSomething()

	o.SayHello()
}
