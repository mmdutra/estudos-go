package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numero := 20

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// limita a variável ao escopo do if/else, fora disso nao eh mais acessivel
	if outroNumero := criaNumero(); outroNumero > 50 {
		fmt.Println(outroNumero, " é maior que 50")
	} else {
		fmt.Println(outroNumero, " é menor ou igual a 50")
	}
}

func criaNumero() int {
	return rand.Intn(100)
}
