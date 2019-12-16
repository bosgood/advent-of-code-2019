package main

import (
	"fmt"
	"math"
	"strings"
)

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// f, err := os.Open("inputs/day12_part1.txt")
	// try(err)
	// b, err := ioutil.ReadAll(f)
	// try(err)
	// trimmed := bytes.TrimSpace(b)
	fmt.Println(Part1(""))
}

type System []Body

func (s System) Step() System {
	s2 := make(System, len(s))
	for i, b1 := range s {
		res := b1
		for j, b2 := range s {
			if i == j {
				continue
			}
			res = res.ApplyGravity(b2)
		}
		s2[i] = res.ApplyVelocity()
	}
	return s2
}

func (s System) Energy() int {
	total := 0
	for _, b := range s {
		total += b.Energy()
	}
	return total
}

type Vector struct {
	X int
	Y int
	Z int
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v Vector) AbsSum() int {
	return int(math.Abs(float64(v.X))) +
		int(math.Abs(float64(v.Y))) +
		int(math.Abs(float64(v.Z)))
}

type Body struct {
	Position Vector
	Velocity Vector
}

func (m Body) ApplyGravity(m2 Body) Body {
	dx1 := compareAxis(m.Position.X, m2.Position.X)
	dy1 := compareAxis(m.Position.Y, m2.Position.Y)
	dz1 := compareAxis(m.Position.Z, m2.Position.Z)

	return Body{
		Position: m.Position,
		Velocity: m.Velocity.Add(Vector{dx1, dy1, dz1}),
	}
}

func (m Body) ApplyVelocity() Body {
	return Body{
		Velocity: m.Velocity,
		Position: m.Position.Add(m.Velocity),
	}
}

func (m Body) Energy() int {
	return m.PotentialEnergy() * m.KineticEnergy()
}

func (m Body) PotentialEnergy() int {
	return m.Position.AbsSum()
}

func (m Body) KineticEnergy() int {
	return m.Velocity.AbsSum()
}

func compareAxis(v1, v2 int) int {
	if v1 > v2 {
		return -1
	} else if v1 < v2 {
		return +1
	} else {
		return 0
	}
}

func Part1(input string) string {
	s := System{
		Body{
			Position: Vector{-4, 3, 15},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{-11, -10, 13},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{2, 2, 18},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{7, -1, 0},
			Velocity: Vector{0, 0, 0},
		},
	}

	for i := 0; i < 1000; i++ {
		s = s.Step()
	}

	return fmt.Sprintf("%d", s.Energy())
}

var (
	puzzle = strings.TrimSpace(
		`
<x=-4, y=3, z=15>
<x=-11, y=-10, z=13>
<x=2, y=2, z=18>
<x=7, y=-1, z=0>
`)
)
