package main

import (
	"math"
	"testing"
)

type task2 struct {
	input  float64
	result float64
}

var testsTask2 = []task2{
	{1, 3},
	{0.1, 3.1},
	{0.01, 3.14},
	{0.001, 3.141},
	{0.0001, 3.1415},
}

func TestTask2(t *testing.T) {
	for _, pair := range testsTask2 {
		v := Task2(pair.input)
		if math.Abs(v-pair.result) > pair.input {
			t.Error(
				"for", pair.input,
				"expected", math.Abs(v-pair.result), ">", pair.input,
			)
		}
	}
}
