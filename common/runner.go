package common

func RunTask(tasks map[byte]func(), num byte) {
	task, exists := tasks[num]
	if !exists {
		panic("Incorrect task number")
	}

	task()
}
