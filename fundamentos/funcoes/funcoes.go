package main

import (
	"fmt"
)

func main() {
	fmt.Println(somar(10, 10))

	var f = func(txt string) string {
		return txt + "Mundo"
	}

	fmt.Println(f("Ola, "))

	fmt.Println(calculos(20, 10))

	resultadoSoma, resultadoSubtracao := calculos(30, 15)

	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma2, _ := calculos(35, 50)
	fmt.Println(resultadoSoma2)

	_, resultadoSubtracao2 := calculos(30, 15)
	fmt.Println(resultadoSubtracao2)
}

func somar(n1 int16, n2 int16) int16 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}
