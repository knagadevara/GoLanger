package circle

type Circle struct {
	radius float32
}

const PI float32 = 3.141592

func (c *Circle) Perimeter() float32 {

	return (2 * c.radius * PI)

}

func (c *Circle) Area() float32 {

	return (c.radius * c.radius) * PI

}
