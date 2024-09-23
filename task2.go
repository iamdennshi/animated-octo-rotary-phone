package main

import (
	"fmt"
	"math"
	"math/rand"
)

func countNumbersAfterDot(number float64) int {
	var count int

	for number < 1 {
		count++
		number *= 10
	}
	return count
}

func shift(number float64, digit int) int {
	for range digit {
		number *= 10
	}
	return int(number)
}

/*
Задание 2
Написать программу, которая будет за счёт иммитационного моделирования вычислять значение числа Пи.
Количество бросков заранее не задано.
Пользователь задаёт точность вычисления, например, 0.01 или 0.1.
*/
func Task2(precision float64) float64 {

	if precision > 1 {
		panic("precision > 1")
	}

	numbersAfterDot := countNumbersAfterDot(precision)

	if precision*math.Pow(10, float64(numbersAfterDot)) != 1 {
		expected := 1 / (float64(numbersAfterDot) * 10)
		msg := fmt.Sprintf("precision is bad expected %v, got %v", expected, precision)
		panic(msg)
	}

	var n, k uint64
	var curDigit int
	var cur float64
	prevDigit := -1
	matched := 0
	step := uint64(1000 * math.Pow(10, float64(numbersAfterDot)))

LOOP:
	for {
		n += step
		for range step {
			randX := rand.Float64()
			randY := rand.Float64()

			if IsInCircle(randX, randY) {
				k++
			}
		}

		cur = 4.0 * float64(k) / float64(n)

		// numbersAfterDot+1 чтобы была стабильна цифра которая идет после нужной
		curDigit = shift(cur, numbersAfterDot+1)

		if prevDigit == curDigit {
			matched++

		} else {
			matched = 0
			prevDigit = curDigit
		}

		if matched == numbersAfterDot {
			break LOOP
		}

	}

	return cur
}
