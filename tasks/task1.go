package tasks

import (
	"math"
	"math/rand"
)

func IsInCircle(x, y float64) bool {
	const circleRadius = 1

	return circleRadius >= math.Sqrt(x*x+y*y)
}

/*
Задание 1
Написать программу, которая будет за счёт иммитационного моделирования вычислять значение числа Пи.
Количество бросков задаёт пользователь, например, 10_000.
*/
func Task1(n int) float64 {
	var k int

	for range n {
		randX := rand.Float64()
		randY := rand.Float64()

		if IsInCircle(randX, randY) {
			k++
		}
	}

	pi := 4.0 * float64(k) / float64(n)

	return pi
}
