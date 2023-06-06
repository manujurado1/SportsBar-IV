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
