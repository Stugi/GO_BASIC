package main

import "fmt"

// slice
// слайс — это ссылочный тип

func main() {
	// Слайс, или срез (от англ. slice)
	// - это структура данных, содержащая указатель на массив,
	// значения длины (len — от англ. length)
	// и ёмкости (cap — от англ. capacity).
	// Как и массивы, слайсы индексируются.

	// var x []float64

	// Для создания слайса на определённое какое-то количество элементов используем функцию make:
	x := make([]float64, 5)
	fmt.Println(x, len(x), cap(x))

	// У нас получится слайс с len = 3, cap = 5.
	sl := make([]float64, 3, 5)
	sl = append(sl, 1, 2, 3) // добавляем 3 элемента
	fmt.Println(sl, len(sl), cap(sl))

	// Low : High
	// Low - начало среза
	// High - конец среза
	sl2 := sl[1:3]
	fmt.Println(sl2, len(sl2), cap(sl2))

	// Ссылочный тип
	sliceFromArray()

	//COPY
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)

	//Копируем второй слайс в первый.
	// Обратите внимание, что во второй слайс скопируются только числа 1 2.
	fmt.Println(slice1, slice2)
}

func sliceFromArray() {
	arr := [5]int{1, 1, 1, 1, 1}
	sl := arr[:]

	fmt.Println(arr, sl)

	sl[0] = 10

	fmt.Println(arr, sl)

	arr[1] = 20

	fmt.Println(arr, sl)

}
