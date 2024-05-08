package main

import (
	"fmt"
)

// Для обработки паники существует встроенная функция recover().
// Пример использования:
func main() {
	Cascade()
	fmt.Println("Everything is fine")
}

func Cascade() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic: %v\n", r)
		}
	}()
	panicFunc()
	fmt.Println("Cascade end")
}

func panicFunc() {
	fmt.Println("just before panic")
	if true {
		panic("this is an artificial panic")
	}
	fmt.Println("panicFunc without panic")
}
