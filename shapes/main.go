package main

import (
	"fmt"
	rct "shapes/rectangle/rectangle.go"
	sqr "shapes/rectangle/square.go"
)

type ShapeER interface {
	Area() float32
	Perimeter() float32
}

func AreaInfo(s ShapeER) {
	fmt.Printf("Area: %v\n", s.Area())
	fmt.Println("----")
}

func PerimeterInfo(s ShapeER) {
	fmt.Printf("Perimeter: %v\n", s.Perimeter())
	fmt.Println("----")
}

func main() {
	var rect rct.Rectangle
	rect = 3.5
	rect = 7.5
	var sq sqr.Square
	sq.side = 4.0
	AreaInfo(&sq)
	AreaInfo(&rect)
}
