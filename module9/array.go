package main

import "fmt"

func main() {
	// одномерный массив
	// в Go массив является не ссылочным типом

	var x [6]int
	x[4] = 100

	fmt.Println(x)

	x[0] = 10
	x[1] = 20
	x[2] = 34
	x[3] = 45
	x[4] = 56
	x[5] = 67

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += float64(x[i])
	}

	fmt.Println(total / 5)

	// ФУНКЦИЯ LEN
	total = 0
	for i := 0; i < len(x); i++ {
		total += float64(x[i])
	}
	fmt.Println(total / float64(len(x)))

	//ЦИКЛ FOR-RANGE
	total = 0
	for i, item := range x {
		total += float64(item)
		fmt.Println(i, item)
	}

	fmt.Println(total / float64(len(x)))

	// короткая	 запись
	x = [6]int{10, 45, 34, 56, 67, 78}

	//Теперь наша функция принимает указатель на массив.
	// Сначала мы создаём массив со всеми единицами, берём указатель на этот массив,
	// передаём его в функцию, в функции меняем значение последнего элемента, и после этого,
	// даже когда мы вышли за пределы функции, наш массив будет изменен.
	arr := [5]float64{1, 1, 1, 1, 1}
	changeArray(&arr)
	fmt.Println(arr)

	//→ Что будет, если обратиться к индексу, которого в массиве не существует?
	// Будет паника!  → Ошибка: runtime error: index out of range [5] with length 5

	// ДВУМЕРНЫЙ МАССИВ
	twoDeimensionalArray()

	//Создание массива с элементами типа int,
	// размерность которого будет автоматически определена в зависимости от количества значений в фигурных скобках.
	var intArr = [...]int{100, 54, 43, 22}
}

func changeArray(arr *[5]float64) {
	arr[4] = 100
}

func twoDeimensionalArray() {
	arr := [2][4]int{
		[4]int{1, 2, 3, 4},
		[4]int{5, 6, 7},
	}
	return arr
}
