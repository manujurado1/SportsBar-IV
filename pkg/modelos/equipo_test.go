package modelos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEquipo(t *testing.T) {

	Equipo := NewEquipo()
	assert.Equal(t, []string{}, Equipo.ObtenerEquipo())

}

func TestRellenarEquipo(t *testing.T) {

	Equipo := NewEquipo()
	ListaAmigosEquipo1 := []string{"Amigo1", "Amigo2", "Amigo3", "Amigo4", "Amigo5"}
	Equipo = Equipo.RellenarEquipo(ListaAmigosEquipo1)
	assert.Equal(t, ListaAmigosEquipo1, Equipo.ObtenerEquipo())
}
