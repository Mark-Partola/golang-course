package homework1

import (
	"course/homework-1/io"
	"course/homework-1/logic"
	"fmt"
)

// RunTask - perform hw1 task
func RunTask(num byte) {
	switch num {
	case 1:
		runTask1()
	case 2:
		runTask2()
	case 3:
		runTask3()
	default:
		panic("Incorrect task number")
	}
}

// Написать программу для конвертации рублей в доллары.
// Программа запрашивает сумму в рублях и выводит сумму в долларах.
// Курс доллара задайте константой.
func runTask1() {
	fmt.Print("Please, input number: ")
	count := io.ScanNumber()
	result := logic.ConvertFromRubToUsd(count)
	fmt.Println(fmt.Sprintf("$%v = %v₽", count, result))
}

// Даны катеты прямоугольного треугольника.
// Найти его площадь, периметр и гипотенузу.
// Используйте тип данных float64 и функции из пакета math.
func runTask2() {
	fmt.Print("Please, input length of first cathet: ")
	cathet1 := io.ScanNumber()
	fmt.Print("Please, input length of second cathet: ")
	cathet2 := io.ScanNumber()

	hypotenuse := logic.GetHypotenuse(cathet1, cathet2)
	perimeter := logic.GetPerimeter(cathet1, cathet2)
	square := logic.GetSquare(cathet1, cathet2)

	fmt.Printf("Hypotenuse = %v\nPerimeter = %v\nSquare = %v", hypotenuse, perimeter, square)
}

// Пользователь вводит сумму вклада в банк и годовой процент.
// Найти сумму вклада через 5 лет.
func runTask3() {
	fmt.Print("Amount: ")
	initialAmount := io.ScanNumber()
	fmt.Print("Interest rate: ")
	interestRate := io.ScanNumber()
	fmt.Print("Years: ")
	years := io.ScanNumber()

	perspectiveAmount := logic.GetPerspectiveAmount(initialAmount, interestRate, years)

	fmt.Printf("Amount %v in %v years with %v interest rate will be %v", initialAmount, years, interestRate, perspectiveAmount)
}
