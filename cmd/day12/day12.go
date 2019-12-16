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
	Part1()
	Part2()
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

func Part1() {
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

	fmt.Printf("Energy after 1000 iterations: %d\n", s.Energy())
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Part2() {
	start := System{
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

	iterations := 0
	current := start
	var periodX int
	var periodY int
	var periodZ int
	for {
		iterations++
		current = current.Step()
		if periodX == 0 && repeatedX(start, current) {
			periodX = iterations
		}
		if periodY == 0 && repeatedY(start, current) {
			periodY = iterations
		}
		if periodZ == 0 && repeatedZ(start, current) {
			periodZ = iterations
		}

		if periodX != 0 && periodY != 0 && periodZ != 0 {
			break
		}
	}

	fmt.Printf("%d iterations for X\n", periodX)
	fmt.Printf("%d iterations for Y\n", periodY)
	fmt.Printf("%d iterations for Z\n", periodZ)

	lcm := LCM(periodX, periodY, periodZ)
	fmt.Printf("LCM for periods: %d\n", lcm)
}

func repeated(s1, s2 System, pred func(Body, Body) bool) bool {
	for i := range s1 {
		if !pred(s1[i], s2[i]) {
			return false
		}
	}
	return true
}

func repeatedX(s1, s2 System) bool {
	return repeated(s1, s2, func(b1, b2 Body) bool {
		return b1.Position.X == b2.Position.X &&
			b1.Velocity.X == b2.Velocity.X
	})
}
func repeatedY(s1, s2 System) bool {
	return repeated(s1, s2, func(b1, b2 Body) bool {
		return b1.Position.Y == b2.Position.Y &&
			b1.Velocity.Y == b2.Velocity.Y
	})
}
func repeatedZ(s1, s2 System) bool {
	return repeated(s1, s2, func(b1, b2 Body) bool {
		return b1.Position.Z == b2.Position.Z &&
			b1.Velocity.Z == b2.Velocity.Z
	})
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
