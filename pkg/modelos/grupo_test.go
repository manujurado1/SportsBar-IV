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
	grupoEsperado := GrupoAmigos{"Prueba", []Amigo{}, map[string]Nivel{}, map[string]bool{}}

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

func TestModificarDisponibilidadAmigo(t *testing.T) {

	//Caso incorrecto intentando cambiar disponibilidad de un amigo que no está en el grupo
	grupo, _ := NewGrupoAmigos("Prueba")
	amigo, _ := NewAmigo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))

	success, errorDisponibilidad := grupo.CambiarDisponibilidadAmigo(amigo, false)

	assert.False(t, success)
	assert.EqualError(t, errorDisponibilidad, ErrorAmigoInexistente.Error())

	//Caso correcto
	_, _ = grupo.CrearAmigoYAniadirAlGrupo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	amigo = grupo.ListaAmigos[0]
	success, errorDisponibilidad = grupo.CambiarDisponibilidadAmigo(amigo, false)

	assert.True(t, success)
	assert.Nil(t, errorDisponibilidad)
	assert.False(t, grupo.DisponibilidadAmigos[amigo.ObtenerId()])

}

func TestAumentarNivelAmigo(t *testing.T) {
	//Caso incorrecto intentando cambiar disponibilidad de un amigo que no está en el grupo
	grupo, _ := NewGrupoAmigos("Prueba")
	amigo, _ := NewAmigo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))

	success, errorNivel := grupo.AumentarNivelAmigo(amigo)

	assert.False(t, success)
	assert.EqualError(t, errorNivel, ErrorAmigoInexistente.Error())

	//Caso correcto sin alcanzar el límite
	_, _ = grupo.CrearAmigoYAniadirAlGrupo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	amigo = grupo.ListaAmigos[0]
	success, errorNivel = grupo.AumentarNivelAmigo(amigo)

	assert.True(t, success)
	assert.Nil(t, errorNivel)
	assert.Greater(t, grupo.NivelAmigos[amigo.ObtenerId()], NivelPorOmision)

	//Caso correcto alcanzando el límite

	grupo.NivelAmigos[amigo.ObtenerId()] = NivelMaximo
	success, errorNivel = grupo.AumentarNivelAmigo(amigo)
	assert.True(t, success)
	assert.Nil(t, errorNivel)
	assert.Equal(t, grupo.NivelAmigos[amigo.ObtenerId()], NivelMaximo)
}

func TestDisminuirNivelAmigo(t *testing.T) {
	//Caso incorrecto intentando cambiar disponibilidad de un amigo que no está en el grupo
	grupo, _ := NewGrupoAmigos("Prueba")
	amigo, _ := NewAmigo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))

	success, errorNivel := grupo.DisminuirNivelAmigo(amigo)

	assert.False(t, success)
	assert.EqualError(t, errorNivel, ErrorAmigoInexistente.Error())

	//Caso correcto sin alcanzar el límite
	_, _ = grupo.CrearAmigoYAniadirAlGrupo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	amigo = grupo.ListaAmigos[0]
	success, errorNivel = grupo.DisminuirNivelAmigo(amigo)

	assert.True(t, success)
	assert.Nil(t, errorNivel)
	assert.Less(t, grupo.NivelAmigos[amigo.ObtenerId()], NivelPorOmision)

	//Caso correcto alcanzando el límite

	grupo.NivelAmigos[amigo.ObtenerId()] = NivelMinimo
	success, errorNivel = grupo.DisminuirNivelAmigo(amigo)
	assert.True(t, success)
	assert.Nil(t, errorNivel)
	assert.Equal(t, grupo.NivelAmigos[amigo.ObtenerId()], NivelMinimo)

}
