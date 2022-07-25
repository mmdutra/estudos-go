package main

import (
	"fmt"
	"time"
)

func main() {
	go contarComIntervalo(2, "primeiro")
	contarComIntervalo(1, "segundo")
}

func contarComIntervalo(intervalo int, contador string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second * time.Duration(intervalo))
		fmt.Println("Contador : ", contador, " -> ", i)
	}
}
