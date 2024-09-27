package main

import (
	"math"
	"tasks/tasks"
	"testing"
)

type task2 struct {
	in  float64
	out float64
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
		v := tasks.Task2(pair.in)
		if math.Abs(v-pair.out) > pair.in {
			t.Error(
				"for", pair.in,
				"expected", pair.out,
				"got", v,
			)
		}
	}
}
