package modelos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNuevoAmigo(t *testing.T) {

	//Caso incorrecto por nombre vacío
	amigo, error := NewAmigo("")
	assert.Errorf(t, error, "El nombre de un jugador no puede ser un string vacío")
	assert.Equal(t, Amigo{}, amigo)

	//Caso correcto
	amigo3, error3 := NewAmigo("Juan")
	amigoEsperado := Amigo{"Juan", NivelPorOmision, true}
	assert.Nil(t, error3)
	assert.Equal(t, amigoEsperado, amigo3)

}

func TestCambiarDisponibilidad(t *testing.T) {

	amigo, _ := NewAmigo("Carlos")
	amigo = amigo.CambiarDisponibilidad(false)

	assert.False(t, amigo.EstaDisponible())

}

func TestAumentarNivelAmigo(t *testing.T) {

	//Caso aumentar dentro del rango permitido
	amigo, _ := NewAmigo("Carlos")
	amigo = amigo.AumentarNivel()

	amigoEsperado := Amigo{"Carlos", Nivel(6), true}

	assert.Equal(t, amigoEsperado, amigo)

	//Caso aumentar superando el rango permitido
	amigo2, _ := NewAmigo("Carlos")
	amigo2.nivel = NivelMaximo
	amigo2.AumentarNivel()

	amigoEsperado2 := Amigo{"Carlos", NivelMaximo, true}

	assert.Equal(t, amigoEsperado2, amigo2)
}

func TestDisminuirNivelAmigo(t *testing.T) {

	//Caso disminuir dentro del rango permitido
	amigo, _ := NewAmigo("Carlos")
	amigo = amigo.DisminuirNivel()

	amigoEsperado := Amigo{"Carlos", Nivel(4), true}

	assert.Equal(t, amigoEsperado, amigo)

	//Caso disminuir superando el rango permitido
	amigo2, _ := NewAmigo("Carlos")
	amigo2.nivel = NivelMinimo
	amigo2.AumentarNivel()

	amigoEsperado2 := Amigo{"Carlos", NivelMinimo, true}

	assert.Equal(t, amigoEsperado2, amigo2)
}
