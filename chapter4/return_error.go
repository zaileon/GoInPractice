package main

import (
	"errors"
	"strings"
)

func Concat(words ...string) (string, error) {
	if len(words) == 0 {
		return "", errors.New("no word provided")
	}
	return strings.Join(words, " "), nil
}

// func main() {
// 	words := os.Args[1:]
// 	if res, err := Concat(words...); err != nil {
// 		fmt.Printf("Error: %s\n", err)
// 	} else {
// 		fmt.Println(res)
// 	}
// }
