package types

import "math"

type Vec2 struct {
	X int
	Y int
}

func (c Vec2) Up() Vec2 {
	return Vec2{c.X, c.Y - 1}
}

func (c Vec2) Down() Vec2 {
	return Vec2{c.X, c.Y + 1}
}

func (c Vec2) Left() Vec2 {
	return Vec2{c.X - 1, c.Y}
}

func (c Vec2) Right() Vec2 {
	return Vec2{c.X + 1, c.Y}
}

func (c Vec2) Around() []Vec2 {
	return []Vec2{
		c.Up(),
		c.Left(),
		c.Down(),
		c.Right(),
	}
}

func (c Vec2) Add(a *Vec2) Vec2 {
	return Vec2{c.X + a.X, c.Y + a.Y}
}

func (c Vec2) Subtract(a *Vec2) Vec2 {
	return Vec2{c.X - a.X, c.Y - a.Y}
}

func (c Vec2) RotateLeft() Vec2 {
	return Vec2{
		X: c.X*int(math.Cos(3*math.Pi/2)) - c.Y*int(math.Sin(3*math.Pi/2)),
		Y: c.X*int(math.Sin(3*math.Pi/2)) + c.Y*int(math.Cos(3*math.Pi/2)),
	}
}

func (c Vec2) RotateRight() Vec2 {
	return Vec2{
		X: c.X*int(math.Cos(math.Pi/2)) - c.Y*int(math.Sin(math.Pi/2)),
		Y: c.X*int(math.Sin(math.Pi/2)) + c.Y*int(math.Cos(math.Pi/2)),
	}
}

func (c Vec2) Multiply(i int) Vec2 {
	return Vec2{X: i * c.X, Y: i * c.Y}
}
