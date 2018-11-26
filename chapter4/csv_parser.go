package main

import (
	"errors"
	"fmt"
	"os"
)

// func main() {
// 	var file io.ReadCloser
// 	file, err := OpenCSV("data.csv")
// 	if err != nil {
// 		fmt.Printf("Error: %s\n", err)
// 		return
// 	}
// 	defer file.Close()
// }

func OpenCSV(filename string) (file *os.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()
	file, err = os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file")
		return file, err
	}
	RemoveEmptyLines(file)
	return file, err
}

func RemoveEmptyLines(f *os.File) {
	panic(errors.New("Failed parse"))
}
