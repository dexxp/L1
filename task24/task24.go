package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// Конструктор Point
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Метод для вычисления расстояния между двумя точками
func (p *Point) DistanceTo(q *Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	distance := point1.DistanceTo(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}