package modelos

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
	assert.Zero(t, len(grupo.JugadoresDisponibilidad))
	assert.Zero(t, len(grupo.JugadoresNiveles))
}

func TestCrearJugador(t *testing.T) {

	JugadoresNiveles := map[string]int{}
	JugadoresDisponibilidad := map[string]bool{}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibilidad: JugadoresDisponibilidad, JugadoresNiveles: JugadoresNiveles}

	success, error := grupo.crearJugador("Manuel", 120, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "El nivel del jugador debe estar entre 0 y 100")

	success, error = grupo.crearJugador("Manuel", 50, false)
	assert.Equal(t, true, success)
	assert.Equal(t, 1, len(grupo.JugadoresNiveles))
	assert.Equal(t, 1, len(grupo.JugadoresDisponibilidad))
	assert.Equal(t, 50, grupo.JugadoresNiveles["Manuel"])
	assert.False(t, grupo.JugadoresDisponibilidad["Manuel"])

	success, error = grupo.crearJugador("Jorge", 70, true)
	assert.Equal(t, true, success)
	assert.Equal(t, 2, len(grupo.JugadoresNiveles))
	assert.Equal(t, 2, len(grupo.JugadoresDisponibilidad))
	assert.Equal(t, 70, grupo.JugadoresNiveles["Jorge"])
	assert.True(t, grupo.JugadoresDisponibilidad["Jorge"])

	success, error = grupo.crearJugador("Manuel", 75, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "Ya existe un jugador con ese nombre")

}

func TestCambiarDisponibilidadJugador(t *testing.T) {

	JugadoresNiveles := map[string]int{"Manuel": 20}
	JugadoresDisponibilidad := map[string]bool{"Manuel": true}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibilidad: JugadoresDisponibilidad, JugadoresNiveles: JugadoresNiveles}

	success, error := grupo.cambiarDisponibilidadJugador("Manuel", false)
	assert.True(t, success)
	assert.Nil(t, error)
	assert.False(t, grupo.JugadoresDisponibilidad["Manuel"])

	success, error = grupo.cambiarDisponibilidadJugador("Jorge", false)
	assert.False(t, success)
	assert.Errorf(t, error, "No existe un jugador con ese nombre")

}

func TestConseguirListaJugadoresDisponibles(t *testing.T) {
	ListaEsperada := []string{"Javier"}

	JugadoresNiveles := map[string]int{"Manuel": 20, "Javier": 50, "Jorge": 80}
	JugadoresDisponibilidad := map[string]bool{"Manuel": false, "Javier": true, "Jorge": false}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibilidad: JugadoresDisponibilidad, JugadoresNiveles: JugadoresNiveles}

	lista, error := grupo.conseguirListaJugadoresDisponibles(grupo.JugadoresDisponibilidad)
	assert.Nil(t, error)
	assert.Equal(t, ListaEsperada, lista)

	grupo.JugadoresDisponibilidad = make(map[string]bool)
	grupo.JugadoresNiveles = make(map[string]int)
	lista, error = grupo.conseguirListaJugadoresDisponibles(grupo.JugadoresDisponibilidad)
	assert.Nil(t, lista)
	assert.Errorf(t, error, "El map de JugadoresDisponibilidad no puede estar vacío")

}

func TestValidarListaJugadoresDisponiblesParaPartido(t *testing.T) {
	ListaMala1 := []string{}
	ListaMala2 := []string{"Jugador1", "Jugador2"}
	ListaMala3 := []string{"Jugador1", "Jugador2", "Jugador3", "Jugador4", "Jugador5", "Jugador6", "Jugador7",
		"Jugador8", "Jugador9", "Jugador10", "Jugador11"}
	ListaBuena := []string{"Jugador1", "Jugador2", "Jugador3", "Jugador4", "Jugador5", "Jugador6", "Jugador7",
		"Jugador8", "Jugador9", "Jugador10", "Jugador11", "Jugador12"}

	grupo := Grupo{Nombre: "GrupoTest"}

	success, error := grupo.validarListaJugadoresDisponiblesParaPartido(ListaMala1)
	assert.False(t, success)
	assert.Errorf(t, error, "La lista de jugadores disponibles no puede estar vacía")

	success, error = grupo.validarListaJugadoresDisponiblesParaPartido(ListaMala2)
	assert.False(t, success)
	assert.Errorf(t, error, "Debe haber almenos 10 jugadores disponibles para jugar un partido")

	success, error = grupo.validarListaJugadoresDisponiblesParaPartido(ListaMala3)
	assert.False(t, success)
	assert.Errorf(t, error, "El número de jugadores debe ser un número par")

	success, error = grupo.validarListaJugadoresDisponiblesParaPartido(ListaBuena)
	assert.True(t, success)
	assert.Nil(t, error)
}
