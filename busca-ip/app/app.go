package app

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

type Application struct {
	Name        string
	Description string
}

// vai retornar a aplicação de linha de comando para ser executada
func gerar(a Application) *cli.App {
	application := cli.NewApp()

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	application.Name = a.Name
	application.Usage = a.Description
	application.Commands = []cli.Command{
		{
			Name:   "busca-ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "busca-servidores",
			Usage:  "Busca nomes de servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return application
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Lista de IPs públicos para o endereço: ", host)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}

func (a Application) Iniciar() {
	aplicacao := gerar(a)
	aplicacao.Run(os.Args)
}
