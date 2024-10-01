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

type search struct {
	in_seq  []int
	in_elem int
	out     int
}

var searchTests = []search{
	{[]int{}, 1, -1},

	{[]int{1, 2, 3}, 1, 0},
	{[]int{1, 2, 3}, 3, 2},
	{[]int{1, 2, 3}, 2, 1},
	{[]int{1, 2, 3}, 0, -1},
	{[]int{1, 2, 3}, 4, -1},

	{[]int{1, 2, 3, 4}, 1, 0},
	{[]int{1, 2, 2, 4}, 3, -1},
	{[]int{1, 2, 2, 4}, 4, 3},
	{[]int{1, 2, 3, 4}, 2, 1},
	{[]int{1, 2, 3, 4}, 0, -1},
	{[]int{1, 2, 3, 4}, 5, -1},
}

func TestBinarySearch(t *testing.T) {
	for _, data := range searchTests {
		v, _ := tasks.BinarySearch(data.in_seq, data.in_elem)
		if v != data.out {
			t.Error(
				"for", data.in_elem,
				"in", data.in_seq,
				"expected", data.out,
				"got", v,
			)
		}
	}
}

func TestInterpolationSearch(t *testing.T) {
	for _, data := range searchTests {
		v, _ := tasks.InterpolationSearch(data.in_seq, data.in_elem)
		if v != data.out {
			t.Error(
				"for", data.in_elem,
				"in", data.in_seq,
				"expected", data.out,
				"got", v,
			)
		}
	}
}
