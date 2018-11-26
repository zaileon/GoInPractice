package main

import (
	"time"
)

// func main() {
// 	msg := make(chan string)
// 	done := make(chan bool)
// 	until := time.After(5 * time.Second)

// 	go send(msg, done)

// 	for {
// 		select {
// 		case m := <-msg:
// 			fmt.Println(m)
// 		case <-until:
// 			done <- true
// 			time.Sleep(500 * time.Millisecond)
// 			return
// 		}
// 	}
// }

func send(msg chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			println("Done")
			close(msg)
			return
		default:
			msg <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
