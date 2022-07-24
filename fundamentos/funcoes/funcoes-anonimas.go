package main

import "fmt"

func main() {
	texto := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Olá, mundo!")

	fmt.Println(texto)
}
