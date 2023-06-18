package main

import (
	"fmt"
)

type ShapeER interface {
	Area() float32
	Perimeter() float32
}

type MyRectangle struct {
	Length  float32
	Breadth float32
}

type MySquare struct {
	Side float32
}

type MyCircle struct {
	Radius float32
}

func AreaInfo(s ShapeER) {
	fmt.Printf("Area:\t%.3f\n", s.Area())
	fmt.Println("----")
}

func PerimeterInfo(s ShapeER) {
	fmt.Printf("Perimeter:\t%.3f\n", s.Perimeter())
	fmt.Println("----")
}

const PI float32 = 3.141592

func (c *MyCircle) Perimeter() float32 {
	fmt.Printf("Circle ")
	return (2 * c.Radius * PI)
}

func (c *MyCircle) Area() float32 {
	fmt.Printf("Circle ")
	return (c.Radius * c.Radius) * PI
}

func (r *MyRectangle) Area() float32 {
	fmt.Printf("Rectangle ")
	return r.Length * r.Breadth
}

func (r *MyRectangle) Perimeter() float32 {
	fmt.Printf("Rectangle ")
	return 2 + (r.Length * r.Breadth)
}

func (s *MySquare) Area() float32 {
	fmt.Printf("Square ")
	return s.Side * s.Side
}

func (s *MySquare) Perimeter() float32 {
	fmt.Printf("Square ")
	return (s.Side * 4.0)
}

func main() {

	var sqr MySquare
	var crl MyCircle
	var rtc MyRectangle

	sqr.Side = 4
	crl.Radius = 3.5
	rtc = MyRectangle{Breadth: 5.5, Length: 7.5}

	PerimeterInfo(&rtc)
	PerimeterInfo(&sqr)
	PerimeterInfo(&crl)
}
