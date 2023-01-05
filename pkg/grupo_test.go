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

func TestCambiarDisponibilidadJugador(t *testing.T) {
	grupo := Grupo{}
	grupo.JugadoresNiveles = make(map[string]int)
	assert.Errorf(t, grupo.cambiarDisponibilidadJugador("Manuel"), "No existe un jugador con ese nombre")
	grupo.JugadoresNiveles["Manuel"] = 6
	assert.Equal(t, 0, len(grupo.JugadoresDisponibles))
	grupo.cambiarDisponibilidadJugador("Manuel")
	assert.Equal(t, 1, len(grupo.JugadoresDisponibles))
	assert.Equal(t, "Manuel", grupo.JugadoresDisponibles[0])
	grupo.cambiarDisponibilidadJugador("Manuel")
	assert.Equal(t, 0, len(grupo.JugadoresDisponibles))

}
