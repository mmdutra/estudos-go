package main

import (
	"fmt"
)

func main() {
	// so irá bloquear caso atinja o tamanho do buffer (2)
	canal := make(chan string, 2)

	canal <- "Olá, mundo!"

	mensagem := <-canal
	fmt.Println(mensagem)
}
