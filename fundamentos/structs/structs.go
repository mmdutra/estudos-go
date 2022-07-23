package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	telefone telefone
}

type telefone struct {
	ddd    string
	numero string
}

func main() {
	var u usuario
	var t telefone = telefone{"31", "9999-9999"}

	fmt.Println(u)

	u.nome = "Mateus Dutra"
	u.idade = 21
	fmt.Println(u)

	u2 := usuario{"Mateus Dutra", 21, t}
	fmt.Println(u2)

	u3 := usuario{nome: "Mateus Dutra", telefone: t}
	fmt.Println(u3)
}
