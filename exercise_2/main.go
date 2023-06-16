package main

import "fmt"

type triangle struct {
	height float64
	base   float64
	sideA  float64
	sideB  float64
}

type square struct {
	length float64
}

func (t triangle) area() float64 {
	return t.base * t.height * 0.5
}

func (t triangle) perimeter() float64 {
	return t.sideA + t.sideB + t.base
}

func (s square) perimeter() float64 {
	return 4 * s.length
}
func (s square) area() float64 {
	return s.length * 2
}

type measure interface {
	area() float64
	perimeter() float64
}

func printPeri(m measure) {
	fmt.Println(m.perimeter())
}
func printMeasure(m measure) {
	fmt.Println(m.area())
}

func main() {
	t := triangle{height: 5.5, base: 2.5, sideA: 5, sideB: 2}
	s := square{length: 2.5}
	printMeasure(t)
	printPeri(t)
	printMeasure(s)
}
