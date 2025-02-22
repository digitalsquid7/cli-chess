package coordinate

type Coordinate [2]int

func Make(x, y int) Coordinate {
	return Coordinate{y, x}
}

func (c Coordinate) Y() int {
	return c[0]
}

func (c Coordinate) X() int {
	return c[1]
}

func (c Coordinate) Left(amount int) Coordinate {
	return Coordinate{c[0], c[1] - amount}
}

func (c Coordinate) Right(amount int) Coordinate {
	return Coordinate{c[0], c[1] + amount}
}

func (c Coordinate) Up(amount int) Coordinate {
	return Coordinate{c[0] - amount, c[1]}
}

func (c Coordinate) Down(amount int) Coordinate {
	return Coordinate{c[0] + amount, c[1]}
}
