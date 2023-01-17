package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearGrupo(t *testing.T) {

	grupo, error := CrearGrupo("")
	assert.Nil(t, grupo)
	assert.Errorf(t, error, "El nombre del grupo no puede ser un string vacío")

	grupo, error = CrearGrupo("Grupo1")
	assert.Nil(t, error)
	assert.Equal(t, "Grupo1", grupo.Nombre)
	assert.Zero(t, len(grupo.JugadoresDisponibles))
	assert.Zero(t, len(grupo.JugadoresNiveles))
}

func TestCrearJugador(t *testing.T) {

	JugadoresNiveles := map[string]int{}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresNiveles: JugadoresNiveles}

	success, error := grupo.crearJugador("Manuel", 120, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "El nivel del jugador debe estar entre 0 y 100")

	success, error = grupo.crearJugador("Manuel", 50, false)
	assert.Equal(t, true, success)
	assert.Equal(t, 1, len(grupo.JugadoresNiveles))
	assert.Equal(t, 0, len(grupo.JugadoresDisponibles))

	success, error = grupo.crearJugador("Jorge", 70, true)
	assert.Equal(t, true, success)
	assert.Equal(t, 2, len(grupo.JugadoresNiveles))
	assert.Equal(t, 1, len(grupo.JugadoresDisponibles))

	success, error = grupo.crearJugador("Manuel", 75, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "Ya existe un jugador con ese nombre")

}

func TestExisteEnArreglo(t *testing.T) {
	ArregloVacio := []string{}
	Arreglo := []string{"prueba1", "prueba2", "prueba3"}

	Existe, error := existeEnArreglo(ArregloVacio, "prueba1")
	assert.False(t, Existe)
	assert.Errorf(t, error, "El arreglo está vacío")

	Existe, error = existeEnArreglo(Arreglo, "prueba1")
	assert.True(t, Existe)
	assert.Nil(t, error)

	Existe, error = existeEnArreglo(Arreglo, "prueba4")
	assert.False(t, Existe)
	assert.Nil(t, error)
}
