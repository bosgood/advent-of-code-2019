package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeIntcode_SampleInputs(t *testing.T) {
	for _, c := range []struct {
		in  string
		out string
	}{
		{
			in:  "1,0,0,0,99",
			out: "2,0,0,0,99",
		},
		{
			in:  "2,3,0,3,99",
			out: "2,3,0,6,99",
		},
		{
			in:  "2,4,4,5,99,0",
			out: "2,4,4,5,99,9801",
		},
		{
			in:  "1,1,1,4,99,5,6,0,99",
			out: "30,1,1,4,2,5,6,0,99",
		},
	} {
		assert.Equal(t, c.out, ComputeIntcode(c.in))
	}
}
