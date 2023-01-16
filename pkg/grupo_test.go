package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearJugador(t *testing.T) {

	JugadoresNiveles := map[string]int{}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresNiveles: JugadoresNiveles}

	success, error := grupo.crearJugador("Manuel", 20, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "El nivel del jugador debe estar entre 0 y 10")

	success, error = grupo.crearJugador("Manuel", 5, false)
	assert.Equal(t, true, success)
	assert.Equal(t, 1, len(grupo.JugadoresNiveles))

	success, error = grupo.crearJugador("Jorge", 7, true)
	assert.Equal(t, true, success)
	assert.Equal(t, 2, len(grupo.JugadoresNiveles))
	assert.Equal(t, 1, len(grupo.JugadoresDisponibles))

	success, error = grupo.crearJugador("Manuel", 7, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "Ya existe un jugador con ese nombre")

}
