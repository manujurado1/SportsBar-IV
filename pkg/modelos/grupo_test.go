package modelos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewGrupoAmigos(t *testing.T) {

	//Caso incorrecto: Nombre vacío
	grupo, errorGrupo := NewGrupoAmigos("")

	assert.EqualError(t, errorGrupo, "Un grupo no puede tener como nombre un string vacío")
	assert.Nil(t, grupo)

	//Caso correcto
	grupo2, errorGrupo2 := NewGrupoAmigos("Prueba")
	grupoEsperado := GrupoAmigos{"Prueba", []Amigo{}}

	assert.Nil(t, errorGrupo2)
	assert.Equal(t, grupoEsperado, *grupo2)
}

func TestCrearAmigoYAniadirAlGrupo(t *testing.T) {

	//Caso incorrecto por nombre amigo vacío
	grupo, _ := NewGrupoAmigos("Prueba")
	success, errorAmigo := grupo.CrearAmigoYAniadirAlGrupo("", time.Date(1990, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	assert.False(t, success)
	assert.EqualError(t, errorAmigo, "El nick de un amigo no puede ser un string vacío")

	//Caso correcto añadiendo 2 amigos con el mismo nombre
	grupo2, _ := NewGrupoAmigos("Prueba2")

	success, errorAmigo = grupo2.CrearAmigoYAniadirAlGrupo("Jose", time.Date(1990, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	assert.True(t, success)
	assert.Nil(t, errorAmigo)

	success, errorAmigo = grupo2.CrearAmigoYAniadirAlGrupo("Jose", time.Date(1990, time.August, 19, 0, 0, 0, 0, time.Now().Location()))
	assert.True(t, success)
	assert.Nil(t, errorAmigo)
	assert.Equal(t, 2, len(grupo2.ListaAmigos))
}
