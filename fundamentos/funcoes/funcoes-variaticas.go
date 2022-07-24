package main

import "fmt"

func soma(numeros ...int) (total int) {
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(10, 20, 30, 40, 50))
	fmt.Println("Olá, mundo", 10, 20, 30, 40, 50)
}
