package helpers

import (
	"bufio"
	"log"
	"os"
)

// LoadStringList loads the input as a slice of strings
func LoadStringList(inputFileName string) []string {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}

// Convertable assumes that the zero value is valid and usable
// e.g. not a nil pointer
type Convertable[T any] interface {
	New() T
	Convert(string) T
}

// LoadStringList loads the input as a slice of strings
func LoadParsedList[T Convertable[T]](inputFileName string) []T {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list []T
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var temp T
		list = append(list, temp.New().Convert(scanner.Text()))
	}
	return list
}
