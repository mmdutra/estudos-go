package main

import "fmt"

// nao sera necessario declarar as variaveis na funcao, no final so dar o return vazio
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(20, 10)
	fmt.Println(soma, subtracao)
}
