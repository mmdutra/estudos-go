package main

import (
	"errors"
	"fmt"
)

func main() {
	// TIPOS INTEIROS: int, int8, int16, int32, int64

	// Usando somente 'int', o go vai usar de acordo com a arquitetura do computador (32 ou 64bits)
	var numero int = 100000000000
	numero2 := 100000000000 // tmb vai seguir a arquitetura do computador
	fmt.Println(numero, numero2)

	var numero3 rune = 123456 // rune == int32
	fmt.Println(numero3)

	var numero4 byte = 123 // byte = int8
	fmt.Println(numero4)

	// TIPOS REAIS: float32, float64

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000000000000.45
	fmt.Println(numeroReal2)

	// STRINGS: string

	var str string = "sdjnfksjndfkjn"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char) // irá printar o valor correspondente a esse caracter na tabela ASC (66)

	var valorInicialInt int16
	fmt.Println(valorInicialInt)

	// BOOLEAN

	var boolean bool = true
	fmt.Println(boolean)

	// ERRORS

	var erro error // valor inicial <nil>
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno do servidor")
	fmt.Println(erro2)
}
