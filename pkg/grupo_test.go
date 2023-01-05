package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearJugador(t *testing.T) {
	grupo := Grupo{}
	grupo.JugadoresNiveles = make(map[string]int)
	assert.Errorf(t, grupo.crearJugador("Manuel", 20, false), "El nivel del jugador debe estar entre 0 y 10")
	assert.Equal(t, nil, grupo.crearJugador("Manuel", 5, false))
	assert.Errorf(t, grupo.crearJugador("Manuel", 6, false), "Ya existe un jugador con ese nombre")

}
