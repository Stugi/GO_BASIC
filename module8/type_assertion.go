package main

// Иногда, имея переменную, в которой содержится интерфейс,
// мы должны получить конкретный тип, находящийся под этим интерфейсом.
// Для этого можно интерфейс попробовать привести к желаемому типу. Это называется type assertion.

import (
	"fmt"
)

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %s article was written by %s", a.Title, a.Author)
}

func typeCast(s fmt.Stringer) {
	article := s.(Article)
	fmt.Printf("%+v\n", article.Title)
	// fmt.Printf("%+v\n", s.Title)
}

func main() {
	a := Article{
		Title:  "Go for beginners",
		Author: "James Bond",
	}

	typeCast(a)
}

// Если нам нужно узнать, какой тип «лежит» под указанными интерфейсом,
// мы можем использовать следующую конструкцию со switch:
// switch s.(type){
// case nil:
// 	fmt.Println("it is nil!")
// case *Article:
// 	fmt.Println("it is a pointer!")
// case Article:
// 	fmt.Println("it is a structure!")
// default:
// 	fmt.Println("is it a bird?")
// }
