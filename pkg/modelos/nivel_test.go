package modelos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetNivel(t *testing.T) {

	//Caso asignación incorrecta
	nivel, error := NuevoNivel(12)
	assert.Errorf(t, error, "El nivel de un jugador debe estar en el rango 0 - 10")
	assert.Equal(t, NivelMedio, nivel)

	//Caso asignación correcta
	nivel2, error2 := NuevoNivel(5)
	assert.Nil(t, error2)
	assert.Equal(t, Nivel(5), nivel2)

}

func TestAumentarNivel(t *testing.T) {

	//Aumentar dentro de los límites
	nivel, error := NuevoNivel(5)
	nivel = nivel.AumentarNivel(3)
	assert.Nil(t, error)
	assert.Equal(t, Nivel(8), nivel)

	//Aumentamos más allá de los límites
	nivel2, error2 := NuevoNivel(8)
	nivel2 = nivel2.AumentarNivel(5)
	assert.Nil(t, error2)
	assert.Equal(t, NivelMaximo, nivel2)
}

func TestDisminuirNivel(t *testing.T) {

	//Disminuir dentro de los límites
	nivel, error := NuevoNivel(5)
	nivel = nivel.DisminuirNivel(3)
	assert.Nil(t, error)
	assert.Equal(t, Nivel(2), nivel)

	//Aumentamos más allá de los límites
	nivel2, error2 := NuevoNivel(3)
	nivel2 = nivel2.DisminuirNivel(5)
	assert.Nil(t, error2)
	assert.Equal(t, NivelMinimo, nivel2)
}
