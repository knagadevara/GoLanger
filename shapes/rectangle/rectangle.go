package rectangle

type Rectangle struct {
	length  float32
	breadth float32
}

func (r *Rectangle) Area() float32 {

	return r.length * r.breadth

}

func (r *Rectangle) Perimeter() float32 {

	return 2 + (r.length * r.breadth)

}
