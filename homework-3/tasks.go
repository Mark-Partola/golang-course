package homework3

import (
	"course/common"
	"course/homework-3/dictionary"
	"course/homework-3/models"
	"course/homework-3/queue"
	"fmt"
)

// RunTask - perform hw2 tasks
func RunTask(num byte) {
	common.RunTask(map[byte]func(){
		1: runTask1and2,
		2: runTask1and2,
		3: runTask3,
		4: runTask4,
	}, num)
}

// Описать несколько структур — любой легковой автомобиль и грузовик.
// Структуры должны содержать марку авто, год выпуска, объем багажника/кузова,
// запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.
// Инициализировать несколько экземпляров структур.
// Применить к ним различные действия. Вывести значения свойств экземпляров в консоль.
func runTask1and2() {
	audi := models.Car{
		IssuedYear: 2019,
		Name:       "audi",
		Body: models.Body{
			Color:         "#000",
			Capacity:      400,
			TrunkCapacity: 100,
		},
		Engine: models.Engine{
			Acceleration: 10,
			FuelCapacity: 30,
			FuelRemains:  21,
		},
	}

	audi.OpenWindow(models.LeftBackWindow)
	err := audi.Fill(300)
	if err != nil {
		fmt.Println("Over the capacity. Fuel spilled out!")
	}

	audi.TurnOn()

	fmt.Println(audi)
}

// * Реализовать очередь.
func runTask3() {
	// by factory
	push, shift := queue.CreateQueue()

	push(1)
	push(2)
	push(3)

	val1, _ := shift()
	val2, _ := shift()
	val3, _ := shift()
	val4, err := shift()

	fmt.Println(val1, val2, val3, val4, err)

	// by structure
	q := queue.Queue{}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	val1, _ = q.Shift()
	val2, _ = q.Shift()
	val3, _ = q.Shift()
	val4, err = q.Shift()

	fmt.Println(val1, val2, val3, val4, err)
}

// * Внести в телефонный справочник дополнительные данные.
// Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.
func runTask4() {
	fileName := "dump.json"
	dict := dictionary.Dictionary{}

	err := dict.Restore(fileName)
	if err != nil {
		panic(err)
	}

	dict.Add(dictionary.Record{
		Name:   "John",
		Phones: []int{797712312332},
	})
	dict.Add(dictionary.Record{
		Name:   "Mike",
		Phones: []int{79179438259},
	})
	dict.Add(dictionary.Record{
		Name:   "Eva",
		Phones: []int{232249592},
	})
	dict.Append(dictionary.Record{
		Name:   "Mike",
		Phones: []int{79483687328},
	})

	err = dict.Persist(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(dict.String())
}
