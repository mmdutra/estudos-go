package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá, mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// encapsula a criação de uma goroutine e cria um canal de conexao com a goroutine
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
