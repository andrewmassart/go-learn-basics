package main

import (
	"fmt"
	"math"
)

type Measurable interface {
	magnitude() float64
}

type Vector2d struct {
	X float64
	Y float64
}

func (vector Vector2d) magnitude() float64 {
	return math.Sqrt(vector.X*vector.X + vector.Y*vector.Y)
}

func (vector *Vector2d) scale(factor float64) {
	vector.X *= factor
	vector.Y *= factor
}

type Vector3d struct {
	X float64
	Y float64
	Z float64
}

func (vector Vector3d) magnitude() float64 {
	return math.Sqrt(vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z)
}

type Particle struct {
	Vector2d
	Mass float64
}

func main() {
	velocity := Vector2d{X: 3, Y: 4}
	fmt.Println("magnitude:", velocity.magnitude())

	velocity.scale(2)
	fmt.Println("after scale * 2:", velocity, velocity.magnitude())

	var measurable Measurable = velocity
	fmt.Println("vector2 as Measurable:", measurable.magnitude())

	measurable = Vector3d{X: 1, Y: 2, Z: 2}
	fmt.Println("vector3 as Measurable:", measurable.magnitude())

	particle := Particle{X: 6, Y: 8, Mass: 2}
	fmt.Println("particle speed:", particle.magnitude())
}
