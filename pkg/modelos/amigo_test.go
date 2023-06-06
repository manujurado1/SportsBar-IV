package modelos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNuevoAmigo(t *testing.T) {

	//Caso incorrecto por nombre vacío
	amigo, error := NuevoAmigo("", 5, true)
	assert.Errorf(t, error, "El nombre de un jugador no puede ser un string vacío")
	assert.Equal(t, Amigo{}, amigo)

	//Caso incorrecto por nivel fuera de rango
	amigo2, error2 := NuevoAmigo("Carlos", 23, false)
	assert.Errorf(t, error2, "El nivel de un jugador debe estar en el rango 0 - 10")
	assert.Equal(t, Amigo{}, amigo2)

	//Caso correcto
	amigo3, error3 := NuevoAmigo("Juan", 7, true)
	amigoEsperado := Amigo{"Juan", 7, true}
	assert.Nil(t, error3)
	assert.Equal(t, amigoEsperado, amigo3)

}

func TestCambiarDisponibilidad(t *testing.T) {

	amigo, _ := NuevoAmigo("Carlos", 7, false)
	amigo.CambiarDisponibilidad(true)

	assert.True(t, amigo.EstaDisponible())

}

func TestAumentarNivelAmigo(t *testing.T) {
	// Caso aumentar dentro del rango
	amigo, _ := NuevoAmigo("Carlos", 7, false)
	amigo.AumentarNivel(2)
	assert.Equal(t, Nivel(9), amigo.ObtenerNivel())

	//Caso aumentar fuera del rango
	amigo2, _ := NuevoAmigo("David", 8, false)
	amigo2.AumentarNivel(4)
	assert.Equal(t, NivelMaximo, amigo2.ObtenerNivel())
}

func TestDisminuirNivelAmigo(t *testing.T) {
	// Caso disminuir dentro del rango
	amigo, _ := NuevoAmigo("Carlos", 4, false)
	amigo.DisminuirNivel(1)
	assert.Equal(t, Nivel(3), amigo.ObtenerNivel())

	//Caso disminuir fuera del rango
	amigo2, _ := NuevoAmigo("David", 3, false)
	amigo2.DisminuirNivel(4)
	assert.Equal(t, NivelMinimo, amigo2.ObtenerNivel())
}
