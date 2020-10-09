package main

import (
	"fmt"
	"sync"
	"time"
)

var ch struct {
	C  chan bool
	C2 chan bool
	C3 chan bool
}
var wg sync.WaitGroup

func initChannels() {
	ch.C = make(chan bool)
	ch.C2 = make(chan bool)
	ch.C3 = make(chan bool)
}

func main() {
	initChannels()

	wg.Add(1)


	go rotinas()
	go rotinas2()
	go rotinas3()

	go observer()

	wg.Wait()

}

func rotinas() {
	fmt.Println("rodou")
	time.Sleep(time.Second * 1)
	ch.C <- false
}

func rotinas2() {
	fmt.Println("rodou2")
	time.Sleep(time.Second * 2)
	ch.C2 <- false
}

func rotinas3() {
	fmt.Println("rodou3")
	time.Sleep(time.Second * 3)
	ch.C3 <- false
}

func observer() {
	defer wg.Done()
	for {
		select {
		case <-ch.C:
			fmt.Println("rodou mesmo 1")
		case <-ch.C2:
			fmt.Println("rodou mesmo 2")
		case <-ch.C3:
			fmt.Println("rodou mesmo 3")
			return
		default:
			fmt.Print("\r")
			fmt.Print("...")
		}
	}
}

