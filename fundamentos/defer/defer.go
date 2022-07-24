package main

import "fmt"

func registrarLog() {
	fmt.Println("Executando log")
}

func salvarDado() {
	fmt.Println("Dado salvo")
}

func main() {
	defer registrarLog()

	salvarDado()
	salvarDado()
}
