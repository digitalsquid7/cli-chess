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

func Up(coor Coordinate, amount int) Coordinate {
	return coor.Up(amount)
}

func Down(coor Coordinate, amount int) Coordinate {
	return coor.Down(amount)
}

func Left(coor Coordinate, amount int) Coordinate {
	return coor.Left(amount)
}

func Right(coor Coordinate, amount int) Coordinate {
	return coor.Right(amount)
}

func TopRight(coor Coordinate, amount int) Coordinate {
	return coor.Up(amount).Right(amount)
}

func TopLeft(coor Coordinate, amount int) Coordinate {
	return coor.Up(amount).Left(amount)
}

func BottomRight(coor Coordinate, amount int) Coordinate {
	return coor.Down(amount).Right(amount)
}

func BottomLeft(coor Coordinate, amount int) Coordinate {
	return coor.Down(amount).Left(amount)
}
