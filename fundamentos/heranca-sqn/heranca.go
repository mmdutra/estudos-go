package main

import "fmt"

// Pqp que gambiarra

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // pega todos os campos que tem em pessoa e adiciona nessa struct
	custo     string
	faculdade string
}

func main() {
	p1 := pessoa{"Joao", "Pedro da Silva", 20, 180}
	fmt.Println(p1)

	e1 := estudante{p1, "Computação", "UFMG"}
	fmt.Println(e1.nome)
}
