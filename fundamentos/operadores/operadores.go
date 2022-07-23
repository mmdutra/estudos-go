package main

import "fmt"

func main() {
	// ARITMETICOS

	soma := 1 + 1
	subtracao := 1 - 1
	divisao := 10 / 2
	multiplicacao := 10 * 4
	resto := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	// ATRIBUICAO

	var var1 string = "texto"
	var2 := "texto2"
	fmt.Println(var1, var2)

	// OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// OPERADORES LOGICOS

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 20
	numero *= 3
	fmt.Println(numero)
}
