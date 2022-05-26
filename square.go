package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {

	return Point{s.start.x + int(s.a), s.start.y + int(s.a)}

	//s.start.x + s.a

	//return s
	// implement me
}

func (s Square) Area() uint {
	return s.a * s.a
	// implement me
}

func (s Square) Perimeter() uint {
	return s.a * 4
	// implement me
}
