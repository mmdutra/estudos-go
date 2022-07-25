package main

import (
	"fmt"
	"testes/entity"
	"time"
)

func main() {
	dataNasc := time.Date(2001, time.Month(1), 14, 0, 0, 0, 0, time.UTC)

	p := entity.Pessoa{"Mateus Dutra", dataNasc}

	fmt.Println(p.EhMaiorDeIdade())
}
