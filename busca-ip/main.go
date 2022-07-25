package main

import (
	"busca-ip/app"
	"fmt"
)

func main() {
	fmt.Println("Iniciando aplicação")

	application := app.Application{"Busca IP", "Busca endereços e nomes de servidores na web"}
	application.Iniciar()
}
