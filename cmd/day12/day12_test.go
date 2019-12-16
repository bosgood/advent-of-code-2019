package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyVelocity(t *testing.T) {
	b := Body{
		Position: Vector{
			X: 1, Y: 2, Z: 3,
		},
		Velocity: Vector{
			X: -2, Y: 0, Z: 3,
		},
	}
	expected := Body{
		Velocity: b.Velocity,
		Position: Vector{
			X: -1, Y: 2, Z: 6,
		},
	}

	actual := b.ApplyVelocity()
	assert.Equal(t, expected, actual)
}

func TestApplyGravity_Single(t *testing.T) {
	b1 := Body{
		Position: Vector{-1, 0, 2},
		Velocity: Vector{0, 0, 0},
	}
	b2 := Body{
		Position: Vector{2, -10, -7},
		Velocity: Vector{0, 0, 0},
	}
	expected := Body{
		Position: b1.Position,
		Velocity: Vector{1, -1, -1},
	}

	mb1 := b1.ApplyGravity(b2)
	assert.Equal(t, expected, mb1)
}

func TestApplyGravity_4(t *testing.T) {
	b1 := Body{
		Position: Vector{-1, 0, 2},
		Velocity: Vector{0, 0, 0},
	}
	b2 := Body{
		Position: Vector{2, -10, -7},
		Velocity: Vector{0, 0, 0},
	}
	b3 := Body{
		Position: Vector{4, -8, 8},
		Velocity: Vector{0, 0, 0},
	}
	b4 := Body{
		Position: Vector{3, 5, -1},
		Velocity: Vector{0, 0, 0},
	}

	for _, b := range []Body{b2, b3, b4} {
		b1 = b1.ApplyGravity(b)
	}
	assert.Equal(t, b1, Body{
		Position: Vector{-1, 0, 2},
		Velocity: Vector{3, -1, -1},
	})
}

func TestStepSystem(t *testing.T) {
	s1 := System{
		Body{
			Position: Vector{-1, 0, 2},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{2, -10, -7},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{4, -8, 8},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{3, 5, -1},
			Velocity: Vector{0, 0, 0},
		},
	}

	s2 := System{
		Body{
			Position: Vector{2, -1, 1},
			Velocity: Vector{3, -1, -1},
		},
		Body{
			Position: Vector{3, -7, -4},
			Velocity: Vector{1, 3, 3},
		},
		Body{
			Position: Vector{1, -7, 5},
			Velocity: Vector{-3, 1, -3},
		},
		Body{
			Position: Vector{2, 2, 0},
			Velocity: Vector{-1, -3, 1},
		},
	}

	assert.Equal(t, s2, s1.Step())
}

func TestStepSystem_10Steps(t *testing.T) {
	s1 := System{
		Body{
			Position: Vector{-1, 0, 2},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{2, -10, -7},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{4, -8, 8},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{3, 5, -1},
			Velocity: Vector{0, 0, 0},
		},
	}

	s2 := System{
		Body{
			Position: Vector{2, 1, -3},
			Velocity: Vector{-3, -2, 1},
		},
		Body{
			Position: Vector{1, -8, 0},
			Velocity: Vector{-1, 1, 3},
		},
		Body{
			Position: Vector{3, -6, 1},
			Velocity: Vector{3, 2, -3},
		},
		Body{
			Position: Vector{2, 0, 4},
			Velocity: Vector{1, -1, -1},
		},
	}

	for i := 0; i < 10; i++ {
		s1 = s1.Step()
	}
	assert.Equal(t, s2, s1)
}

func TestEnergy_10Steps(t *testing.T) {
	s := System{
		Body{
			Position: Vector{-1, 0, 2},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{2, -10, -7},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{4, -8, 8},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{3, 5, -1},
			Velocity: Vector{0, 0, 0},
		},
	}

	for i := 0; i < 10; i++ {
		s = s.Step()
	}
	assert.Equal(t, 6, s[0].PotentialEnergy())
	assert.Equal(t, 6, s[0].KineticEnergy())
	assert.Equal(t, 36, s[0].Energy())

	assert.Equal(t, 9, s[1].PotentialEnergy())
	assert.Equal(t, 5, s[1].KineticEnergy())
	assert.Equal(t, 45, s[1].Energy())

	assert.Equal(t, 10, s[2].PotentialEnergy())
	assert.Equal(t, 8, s[2].KineticEnergy())
	assert.Equal(t, 80, s[2].Energy())

	assert.Equal(t, 6, s[3].PotentialEnergy())
	assert.Equal(t, 3, s[3].KineticEnergy())
	assert.Equal(t, 18, s[3].Energy())

	assert.Equal(t, 179, s.Energy())
}

func TestEnergy_100Steps(t *testing.T) {
	s := System{
		Body{
			Position: Vector{-8, -10, 0},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{5, 5, 10},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{2, -7, 3},
			Velocity: Vector{0, 0, 0},
		},
		Body{
			Position: Vector{9, -8, -3},
			Velocity: Vector{0, 0, 0},
		},
	}

	for i := 0; i < 100; i++ {
		s = s.Step()
	}
	assert.Equal(t, 1940, s.Energy())
}
