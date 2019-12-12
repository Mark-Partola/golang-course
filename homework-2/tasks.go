package homework2

import "course/common"

// RunTask - perform hw2 tasks
func RunTask(num byte) {
	common.RunTask(map[byte]func(){}, num)
}
