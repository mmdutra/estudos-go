package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan int)

	go escrever("Olá, mundo", canal)

	fmt.Println("Escrevendo antes da funcao `escrever` acabar de ser executada")

	for {
		numero, aberto := <-canal

		if !aberto {
			break
		}

		fmt.Println(numero)
	}

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim da execução")
}

func escrever(texto string, canal chan int) {
	for i := 0; i < 5; i++ {
		canal <- i
		time.Sleep(time.Second)
	}

	close(canal)
}
