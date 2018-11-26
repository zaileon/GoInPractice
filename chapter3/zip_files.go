package main

import (
	"compress/gzip"
	"io"
	"os"
)

// func main() {
// 	var wg sync.WaitGroup
// 	for _, file := range os.Args[1:] {
// 		wg.Add(1)
// 		go func(filename string) {
// 			compress(filename)
// 			wg.Done()
// 		}(file)
// 	}
// 	wg.Wait()
// 	fmt.Printf("Compressed completed\n")
// }

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
