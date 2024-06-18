package main

import "fmt"

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		switch {
// 		case i%3 == 0 && i%5 == 0:
// 			println("FizzBuzz")
// 		case i%3 == 0:
// 			println("Fizz")
// 		case i%5 == 0:
// 			println("Buzz")
// 		default:
// 			println(i)
// 		}
// 	}
// }

// программа считает сумму чётных чисел от 0 до 50 включительно и выводит её
// func main() {
// 	var sum int64
// 	for i := 0; i < 50; i++ {
// 		if i%2 == 0 {
// 			sum += int64(i)
// 		}
// 	}

// 	log.Println(sum)
// }

//Напишите функцию, которая находит точку пересечения двух прямых,
// заданных как y = ax + b на плоскости. Коэффициенты a1, b1, a2, b2
// задайте в начале программы или введите из консоли.
// Если прямые не пересекаются, вывести сообщение об этом.

// func main() {
// 	var y, x float64
// 	var a1, a2 float64 = 3, 4
// 	var b1, b2 float64 = 5, 8

// 	x = (b2 - b1) / (a1 - a2)
// 	y = a1*x + b1
// 	fmt.Printf("x: %f, y: %f\n", x, y)
// }

// Напишите программу, выводящую первые 20 простых чисел.
// func main() {
// 	sum := 0
// 	for i := 2; sum < 20; i++ {
// 		var isSimple bool = true
// 		for j := 2; j < i; j++ {
// 			if i%j == 0 {
// 				isSimple = false
// 				break
// 			}
// 		}
// 		if isSimple {
// 			fmt.Println(i)
// 			sum++
// 		}

// 	}
// }

// Напишите функцию, которая ищет факториал натурального числа в итеративном и рекурсивном виде.
// func main() {
// 	FactI(23)
// 	FactR(23)
// }

// func FactI(n int64) int64 {
// 	var factorial int64 = 1
// 	for j := int64(1); j <= n; j++ {
// 		factorial *= int64(j)
// 	}

// 	return factorial
// }

// func FactR(n int64) int64 {
// 	if n < 2 {
// 		return 1
// 	}
// 	return n * FactR(n-1)
// }

// Палиндром
// func main() {
// 	IsPalindrome(13245212)
// }

// func IsPalindrome(n int) bool {
// 	a, b := 0, n
// 	for b > 0 {
// 		a *= 10
// 		a += b % 10
// 		b /= 10

// 	}
// 	return a == n
// }

// Например: “aaaa” => “a4” или “abcccccc” => “a1b1c6”.
func main() {
	str := "abcccccc"
	res := ""
	curSym := ""
	for ind, sym := range str {

	}

	fmt.Println(res)
}
