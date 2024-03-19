package main

import (
	"fmt"
	"time"
)

func sendPing(ch chan string) {
	for {
		ch <- "Ping"
		time.Sleep(time.Second * 1)
	}
}

func sendPong(ch chan string) {
	i := 1
	for {
		time.Sleep(time.Second * 1)
		ch <- "Pong"
		i++
		if i >= 10 {
			break
		}
	}
}

func main() {
	pingPongChan := make(chan string)

	go sendPing(pingPongChan)
	go sendPong(pingPongChan)
	for {
		message := <-pingPongChan
		fmt.Println(message)
		time.Sleep(time.Second * 1)
	}

}
