package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
	idade     int8
	notas     [2]float64
}

// nao ira fazer alteracoes no estado da struct, portanto sera passada uma copia
func (u usuario) estaAprovado() bool {
	media := (u.notas[0] + u.notas[1]) / 2

	return media >= 6
}

// como ira alterar o estado da struct, eh necessario passar via ponteiro
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	notas := [2]float64{9, 8}
	usuario1 := usuario{"Mateus", "Dutra", 21, notas}

	fmt.Println(usuario1.estaAprovado())
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
