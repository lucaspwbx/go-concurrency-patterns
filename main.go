package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	secondExample()
}

func firstExample() {
	go randomBoring("hello world")
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

func secondExample() {
	c := make(chan string)
	go boringWithChannel("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving")
}

//prints the msg in fixed intervals
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

//prints the msg in random intervals
func randomBoring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

//uses a channel to send boring msgs
func boringWithChannel(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
