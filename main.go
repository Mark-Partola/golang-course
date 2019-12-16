package main

import (
	treeutil "course/tree-util"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		panic("You must provide directory in arguments")
	}

	treeutil.PrintTree(args[0])
}
