package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		contarComIntervalo(2, "PRIMEIRA GOROUTINE")
		waitGroup.Done()
	}()

	go func() {
		contarComIntervalo(1, "SEGUNDA GOROUTINE")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func contarComIntervalo(intervalo int, contador string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * time.Duration(intervalo))
		fmt.Println("Contador : ", contador, " -> ", i)
	}
}
