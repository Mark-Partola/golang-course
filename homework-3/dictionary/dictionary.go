package dictionary

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Dictionary -
type Dictionary map[string][]int

// Record -
type Record struct {
	Name   string
	Phones []int
}

// Add -
func (dict *Dictionary) Add(record Record) {
	(*dict)[record.Name] = record.Phones
}

// Append -
func (dict *Dictionary) Append(record Record) {
	(*dict)[record.Name] = append((*dict)[record.Name], record.Phones...)
}

func (dict Dictionary) String() string {
	result := ""
	for name, numbers := range dict {

		result += fmt.Sprintln("Абонент:", name)
		for i, number := range numbers {
			result += fmt.Sprintf("\t %v) %v \n", i+1, number)
		}
	}

	return result
}

// Persist -
func (dict *Dictionary) Persist(fileName string) error {
	value, err := json.Marshal(dict)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.Write(value)

	writer.Flush()
	file.Sync()

	return nil
}

// Restore -
func (dict *Dictionary) Restore(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &dict)

	return nil
}
