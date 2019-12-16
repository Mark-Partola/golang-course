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
	space               = "  "
)

// PrintTree -
func PrintTree(path string) error {
	return renderLevel(path, []bool{})
}

func renderLevel(path string, history []bool) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Cannot read file or directory %v", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Cannot read stat of file or directory %v", err)
	}

	if stat.IsDir() {
		if len(history) != 0 {
			connector := calcConnector(history)
			renderLine(connector, stat)
		}

		files, err := file.Readdirnames(-1)
		if err != nil {
			return fmt.Errorf("Cannot retrieve directory files %v", err)
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

	return nil
}

func calcConnector(history []bool) string {
	historyConnector := ""
	for idx, isLast := range history {
		if idx == len(history)-1 {
			if isLast {
				historyConnector += space + endConnector
			} else {
				historyConnector += space + middleConnector
			}

			break
		}

		if isLast {
			historyConnector += space + lastParentConnector
		} else {
			historyConnector += space + parentConnector
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
