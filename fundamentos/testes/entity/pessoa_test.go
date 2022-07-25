package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func TestEhMaiorDeIdade(t *testing.T) {

	t.Run("Pessoa com menos de 18 anos", func(t *testing.T) {
		menorDeIdade := Pessoa{"Teste", newDate(time.Now().Year(), 10, 10)}
		assert.False(t, menorDeIdade.EhMaiorDeIdade())
	})

	t.Run("Pessoa com 18 anos", func(t *testing.T) {
		maiorDeIdade := Pessoa{"Teste", newDate((time.Now().Year() - 18), 10, 10)}
		assert.True(t, maiorDeIdade.EhMaiorDeIdade())
	})

	t.Run("Pessoa maior de 18 anos", func(t *testing.T) {
		maiorDeIdade := Pessoa{"Teste", newDate(2000, 10, 10)}
		assert.True(t, maiorDeIdade.EhMaiorDeIdade())
	})
}
