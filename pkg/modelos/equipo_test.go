package modelos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEquipo(t *testing.T) {

	Equipo := NewEquipo()
	fechaEsperada := time.Now().Format("02-01-2006")
	assert.Equal(t, []string{}, Equipo.ObtenerEquipo())
	assert.Equal(t, fechaEsperada, Equipo.ObtenerFechaCreacion())

}

func TestRellenarEquipo(t *testing.T) {

	Equipo := NewEquipo()
	ListaAmigosEquipo1 := []string{"Amigo1", "Amigo2", "Amigo3", "Amigo4", "Amigo5"}
	Equipo = Equipo.RellenarEquipo(ListaAmigosEquipo1)
	assert.Equal(t, ListaAmigosEquipo1, Equipo.ObtenerEquipo())
}
