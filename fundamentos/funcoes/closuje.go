package main

import "fmt"

func closuje() func() {
	texto := "Dentro da funcao closuje"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closuje()
	funcaoNova()
}
