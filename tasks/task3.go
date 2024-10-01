package tasks

import (
	"fmt"
	"math"
	"math/rand"
)

// seq - отсортированный массив
// x - элемент который нужно найти
// Возвращает (индекс элемента, количество итераций)
func BinarySearch(seq []int, x int) (res int, iters int) {
	if len(seq) == 0 || x < seq[0] || x > seq[len(seq)-1] {
		return -1, 0
	}

	left := 0
	right := len(seq) - 1

	for left < right-1 {
		mid := (right + left) / 2
		iters++
		if seq[mid] > x {
			right = mid
		} else if seq[mid] < x {
			left = mid
		} else {
			return mid, iters
		}
	}

	if seq[left] == x {
		return left, iters
	} else if seq[right] == x {
		return right, iters
	}

	return -1, iters
}

// seq - отсортированный массив
// x - элемент который нужно найти
// Возвращает (индекс элемента, количество итераций)
func InterpolationSearch(seq []int, x int) (res int, iters int) {
	if len(seq) == 0 || x < seq[0] || x > seq[len(seq)-1] {
		return -1, 0
	}

	left := 0
	right := len(seq) - 1

	for seq[left] < x && seq[right] > x {
		mid := left + (x-seq[left])*(right-left)/(seq[right]-seq[left])
		iters++
		if seq[mid] > x {
			right = mid - 1
		} else if seq[mid] < x {
			left = mid + 1
		} else {
			return mid, iters
		}
	}

	if seq[left] == x {
		return left, iters
	} else if seq[right] == x {
		return right, iters
	}

	return -1, iters
}

// генерирует возврастающую экспоненциальную последовательность целых чисел
// n - количество элементов
// base - основание
func genSortedExpSeq(n, base int) []int {
	result := make([]int, n)
	for i := 1; i < n; i++ {
		result[i] = int(math.Pow(float64(base), float64(i)))
		if result[i] < 0 {
			panic("Overflow")
		}
	}
	return result
}

// генерирует возврастающую линейную последовательность целых чисел
// n - количество элементов
// a, b - диапазон шага (рандомный)
func genSortedSeq(n, a, b int) []int {
	result := make([]int, n)

	for i := 1; i < n; i++ {
		result[i] = result[i-1] + rand.Intn(b-a) + a
	}
	return result
}

// генерирует возврастающую кубическую последовательность целых чисел
// n - количество элементов
// a, b - диапазон шага (рандомный)
func genSortedCubeSeq(n, a, b int) []int {
	result := make([]int, n)

	for i := 1; i < n; i++ {
		result[i] = result[i-1] + int(math.Pow(float64(rand.Intn(b-a)+a), 3))
	}
	return result
}

/* Задание 3
* Дана возрастающая последовательность целых чисел.
* Требуется написать функции:
* - бинарный поиск
* - интерполяционный поиск
*	которые возвращают позицию числа
* если не нашли, то -1
* Сравнить эти алгоритмы по количеству шагов поиска.
 */
func Task3() {

	// Количество элементов в последовательностях
	seqIters := []int{30, 3_000, 3_000_000}

	// Названия последовательностей
	seqNames := []string{"экспоненциальной", "кубической", "линейной"}

	// Генерация последовательностей
	seqes := [][]int{
		// Экспоненциальная
		genSortedExpSeq(seqIters[0], 2),
		// Кубическая
		genSortedCubeSeq(seqIters[1], 1, 10_000),
		// Линейная
		genSortedSeq(seqIters[2], 1, 10_000),
	}

	// Случайные элементы которые нужно найти
	xs := []int{seqes[0][rand.Intn(seqIters[0])], seqes[1][rand.Intn(seqIters[1])], seqes[2][rand.Intn(seqIters[2])]}

	for i := range len(seqIters) {
		fmt.Printf("Для %v последовательности из %v элементов\n", seqNames[i], seqIters[i])

		_, iters := InterpolationSearch(seqes[i], xs[i])
		fmt.Printf("- InterpolationSearch выполнился за %v итераций\n", iters)

		_, iters = BinarySearch(seqes[i], xs[i])
		fmt.Printf("- BinarySearch выполнился за %v итераций\n\n", iters)
	}
}
