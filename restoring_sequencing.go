package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	c := make(chan Message)
	go boring("joe", c)
	go boring("ann", c)
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
}

func boring(msg string, c chan Message) {
	var waitForIt chan bool
	for i := 0; i < 5; i++ {
		waitForIt = make(chan bool)
		c <- Message{
			fmt.Sprintf("%s: %d", msg, i),
			waitForIt,
		}
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		<-waitForIt
	}
}
