package treeutil

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	endConnector        = "└"
	middleConnector     = "├"
	strictLine          = "─"
	parentConnector     = "│"
	lastParentConnector = " "
)

// PrintTree -
func PrintTree(path string) {
	renderLevel(path, []bool{})
}

// TODO: defer close file
// TODO: don't panic
// TODO: simplify
func renderLevel(path string, history []bool) {
	path, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Cannot read file or directory %v", err))
	}

	stat, err := file.Stat()
	if err != nil {
		panic(fmt.Sprintf("Cannot read stat of file or directory %v", err))
	}

	if stat.IsDir() {
		connector := calcConnector(history)
		renderLine(connector, stat)

		files, err := file.Readdirnames(-1)
		if err != nil {
			panic(fmt.Sprintf("Cannot retrieve directory files %v", err))
		}

		sort.Strings(files)

		for idx, file := range files {
			isCurrentLast := idx == len(files)-1
			next := filepath.Join(path, file)

			updatedHistory := make([]bool, len(history))
			copy(updatedHistory, history)
			updatedHistory = append(updatedHistory, isCurrentLast)
			renderLevel(next, updatedHistory)
		}
	} else {
		connector := calcConnector(history)
		renderLine(connector, stat)
	}
}

func calcConnector(history []bool) string {
	historyConnector := ""
	for idx, isLast := range history {
		if idx == len(history)-1 {
			if isLast {
				historyConnector += endConnector
			} else {
				historyConnector += middleConnector
			}

			break
		}

		if isLast {
			historyConnector += lastParentConnector
		} else {
			historyConnector += parentConnector
		}
	}

	return historyConnector + strings.Repeat(strictLine, 2)
}

func renderLine(connector string, stat os.FileInfo) {
	if stat.IsDir() {
		fmt.Printf("%v%v\n", connector, stat.Name())
	} else {
		fmt.Printf("%v%v (%vb)\n", connector, stat.Name(), stat.Size())
	}
}
