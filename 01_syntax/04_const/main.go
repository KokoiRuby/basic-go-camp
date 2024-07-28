package main

const External = "External"
const internal = "internal"

const (
	b = 456
	c = 789
)

const (
	StatusA = iota
	StatusB
	StatusC

	StatusD = 100
	StatusE
)

const (
	DayA = iota*12 + 13
	DayB
)

func main() {
	const a = 123
}
