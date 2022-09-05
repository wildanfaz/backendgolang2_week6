package taskbasic3

import "math"

type Calculating2d interface {
	Area() float64
	Perimeter() float64
}

type Calculating3d interface {
	Volume() float64
}

type AllCalculate interface {
	Calculating2d
	Calculating3d
}

type Cube struct {
	Side float64
}

func (num *Cube) Area() float64 {
	var area = math.Pow(num.Side, 2)
	return area
}

func (num *Cube) Perimeter() float64 {
	var perimeter = num.Side * 4
	return perimeter
}

func (num *Cube) Volume() float64 {
	var volume = math.Pow(num.Side, 3)
	return volume
}
