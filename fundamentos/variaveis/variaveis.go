package main

import "fmt"

func main() {
	var v1 string = "String declarada"
	v2 := "String implicita"

	// str_2 = 2 -> vai dar ruim! :D

	fmt.Println(v1, v2)

	var (
		v3 string = "whatever"
		v4 int    = 10
	)

	fmt.Println(v3, v4)

	v5, v6 := "Teste", "Teste 2"

	fmt.Println(v5, v6)

	const const1 string = "Constante"

	// const1 = "snfdsjnd" -> vai dar ruim! :D

	// inverter valores :p
	v5, v6 = v6, v5
	fmt.Println(v5, v6)
}
