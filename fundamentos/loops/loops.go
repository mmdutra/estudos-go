package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 3 {
		time.Sleep(time.Second)
		fmt.Println("Valor de i: ", i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Valor de j: ", j)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	// o primeiro parametro sempre eh o indice, quando nao precisar, colocar _
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// na letra ira retornar o codigo da tabela ASC da letra
	for indice, letra := range "Mateus" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	x := 20
	// para fazer um loop infinito
	for {
		if x >= 23 {
			fmt.Println("Loop acabooou")
			break
		}
		fmt.Println("Loop ainda nao acabou")
		time.Sleep(time.Second)
		x++
	}
}
