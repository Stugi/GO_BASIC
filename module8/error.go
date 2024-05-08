package main

import (
	"fmt"
)

//Для упаковки ошибок используется функция Errorf и спецификатор %w.
// Переменная — аргумент для %w и есть ошибка.

// Создав таким образом ошибку с добавлением глагола и передав туда ошибку,
// она будет содержать метод Unwrap,
// который возвращает нашу вложенную ошибку во внешнюю функцию.

func main() {
	i, err := someFunc(0)
	fmt.Println(i, err)

	i, err = someFunc(10)
	fmt.Println(i, err)
}

func someFunc(i int) (int, error) {
	j, err := funcReturningError(i)
	if err != nil {
		return 0, fmt.Errorf("wrap error: %w", err)
	}

	return j, nil
}

func funcReturningError(i int) (int, error) {
	if i == 0 {
		return 0, fmt.Errorf("got %d", i)
	}

	return i, nil
}

type CustomError struct {
	msg string
}

func NewCustomError(msg string) *CustomError {
	return &CustomError{msg: msg}
}

func (err *CustomError) Error() string {
	return err.msg
}

// Конструкция вида var _ error = (*CustomerErr)(nil) позволяет нам проверить,
// что структура (CustomerErr) соответствует интерфейсу (error).
var _ error = (*CustomError)(nil)
