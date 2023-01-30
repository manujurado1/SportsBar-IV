package modelos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearJugador(t *testing.T) {

	grupo, error := CrearGrupo("GrupoTest")
	assert.Nil(t, error)

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

func TestCrearEquiposIgualadosParaPartido(t *testing.T) {

	// Comprobamos que si no hay los suficientes jugadores disponibles no se pueden crear equipos igualados
	JugadoresNiveles := map[string]int{"Manuel": 30, "Jorge": 50, "Edu": 10, "Clara": 90, "Migue": 100, "Alberto": 70,
		"Javi": 20, "Lorena": 80, "Maria": 60, "Sergio": 40}
	JugadoresDisponibilidad := map[string]bool{"Manuel": false, "Jorge": false, "Edu": false, "Migue": true, "Clara": true, "Alberto": true,
		"Javi": true, "Lorena": true, "Maria": true, "Sergio": false}

	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibilidad: JugadoresDisponibilidad, JugadoresNiveles: JugadoresNiveles}

	Equipo1, Equipo2, error := grupo.crearEquiposIgualadosParaPartido(grupo.JugadoresDisponibilidad)
	assert.Errorf(t, error, "Debe haber almenos 10 jugadores disponibles para jugar un partido")

	// Comprobamos la posibilidad de que aunque todo haya ido bien, sea imposible conseguir 2 equipos igualados

	JugadoresDisponibilidad = map[string]bool{"Manuel": true, "Jorge": true, "Edu": true, "Migue": true, "Clara": true, "Alberto": true,
		"Javi": true, "Lorena": true, "Maria": true, "Sergio": true}
	JugadoresNivelesImposible := map[string]int{"Manuel": 3, "Jorge": 5, "Edu": 1, "Clara": 9, "Migue": 100, "Alberto": 7,
		"Javi": 2, "Lorena": 8, "Maria": 6, "Sergio": 4}
	grupo.JugadoresDisponibilidad = JugadoresDisponibilidad
	grupo.JugadoresNiveles = JugadoresNivelesImposible

	Equipo1, Equipo2, error = grupo.crearEquiposIgualadosParaPartido(grupo.JugadoresDisponibilidad)
	assert.Errorf(t, error, "No se han podido crear 2 equipos igualados (La diferencia entre ambos equipos es > al 30 por ciento)")
	igualados, error := grupo.estanIgualados(Equipo1, Equipo2)
	assert.False(t, igualados)

	// Comprobamos caso correcto

	JugadoresNiveles = map[string]int{"Manuel": 30, "Jorge": 50, "Edu": 10, "Clara": 90, "Migue": 100, "Alberto": 70,
		"Javi": 20, "Lorena": 80, "Maria": 60, "Sergio": 40}

	grupo.JugadoresNiveles = JugadoresNiveles

	Equipo1Esperado := []string{"Migue", "Alberto", "Jorge", "Sergio", "Edu"}
	Equipo2Esperado := []string{"Clara", "Lorena", "Maria", "Manuel", "Javi"}

	Equipo1, Equipo2, error = grupo.crearEquiposIgualadosParaPartido(grupo.JugadoresDisponibilidad)
	assert.Equal(t, Equipo1Esperado, Equipo1)
	assert.Equal(t, Equipo2Esperado, Equipo2)
	assert.Nil(t, error)
	igualados, error = grupo.estanIgualados(Equipo1, Equipo2)
	assert.True(t, igualados)

}
