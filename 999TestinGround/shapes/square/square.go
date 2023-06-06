package square

type Square struct {
	side float32
}

func (s *Square) Area() float32 {

	return s.side * s.side

}

func (s *Square) Perimeter() float32 {

	return (s.side * 4.0)

}
