package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	a int
	b int
}

type square struct {
	a int
	b int
}

func main() {
	t := triangle{a: 2, b: 3}
	s := square{a: 2, b: 3}
	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return float64((t.a * t.b) / 2)
}

func (s square) getArea() float64 {
	return float64(s.a * s.b)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
