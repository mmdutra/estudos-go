package entity

import (
	"time"
)

type Pessoa struct {
	Name     string
	DataNasc time.Time
}

func (p Pessoa) EhMaiorDeIdade() bool {
	dataAtual := time.Now()
	return (dataAtual.Year() - p.DataNasc.Year()) >= 18
}
