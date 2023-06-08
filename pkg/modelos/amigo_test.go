package modelos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNuevoAmigo(t *testing.T) {

	//Caso incorrecto por nombre vacío
	amigo, error := NewAmigo("", 01, "Enero")
	assert.Errorf(t, error, "El nombre de un jugador no puede ser un string vacío")
	assert.Equal(t, Amigo{}, amigo)

	//Caso correcto
	amigo2, error2 := NewAmigo("Juan", 01, "Enero")
	identificativoEsperado := "Juan1enero"
	assert.Nil(t, error2)
	assert.Equal(t, identificativoEsperado, amigo2.ObtenerId())

}
