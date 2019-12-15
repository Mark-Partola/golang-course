package homework2

import (
	"course/common"
	"course/common/io"
	"course/homework-2/logic"
	"fmt"
	"math/big"
)

// RunTask - perform hw2 tasks
func RunTask(num byte) {
	common.RunTask(map[byte]func(){
		1: runTask1,
		2: runTask2,
		3: runTask3,
		4: runTask4,
	}, num)
}

// Написать функцию, которая определяет, четное ли число.
func runTask1() {
	fmt.Print("Enter number: ")
	num, _ := io.ScanNumber()

	isEven, _ := logic.CheckMultiplicity(int(num), 2)

	if isEven {
		fmt.Printf("Number %v is even", num)
	} else {
		fmt.Printf("Number %v is odd", num)
	}
}

// Написать функцию, которая определяет, делится ли число без остатка на 3
func runTask2() {
	fmt.Print("Enter number: ")
	num, _ := io.ScanNumber()

	isMultipleBy3, _ := logic.CheckMultiplicity(int(num), 3)

	if isMultipleBy3 {
		fmt.Printf("Number %v is multiple by 3", num)
	} else {
		fmt.Printf("Number %v is not multiple by 3", num)
	}
}

// Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0
// Следить за переполнением
func runTask3() {
	fmt.Print("Enter number: ")
	num, _ := io.ScanNumber()

	logic.ForEachFibonacci(uint(num), func(value big.Int, idx uint) {
		fmt.Println(fmt.Sprintf("%v: %v", idx+1, value.String()))
	})
}

// Заполнить массив из 100 элементов различными простыми числами по Эратосфену
func runTask4() {
	list := logic.CalculateSimpleNumbersUpTo(100)

	fmt.Println(list)
}
