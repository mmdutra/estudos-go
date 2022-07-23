package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	validarEmail("teste@teste.com")
	validarEmail("teste")
}

func validarEmail(email string) {
	erro := checkmail.ValidateFormat(email)

	fmt.Println(erro)
}
