package main

import "fmt"

// map Карта

func main() {
	// map
	// Карта — это ссылочный тип.
	var x map[string]int
	//инициализировать карту можно с помощью make
	x = make(map[string]int)

	x["key"] = 10

	// существует ли значение в карте
	value, ok := x["key"]

	// удалить значение
	delete(x, "key")

	fmt.Println(value, ok)

	// Ключ может быть:
	// -	строкой,
	// -	числом,
	// -	указателем на структуру,
	// -	структурой.
	// Значением мапы может быть всё что угодно, даже другая мапа или массив.

	// Ключом мапы НЕ может быть:
	// -	другая мапа,
	// -	функция
	// -	слайс.

	// Значение мапы может быть другой мапа
	elements := map[string]map[string]string{
		// полная запись
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		// сокращенная запись
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
	}

	fmt.Println(elements)

	// range
	mapLoop()
}

func mapLoop() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}
