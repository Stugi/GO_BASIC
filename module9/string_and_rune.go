package main

import "fmt"

func main() {
	strToRune("Привет")
}

func strToRune(str string) {
	// через range
	for ind, el := range str {
		fmt.Println(fmt.Sprintf("%#U", el), ind)
	}

	// через приведение типа string -> []rune
	runes := []rune(str)
	fmt.Println(runes, len(runes))
}
