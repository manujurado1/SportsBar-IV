package modelos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNivel(t *testing.T) {

	nivel := NewNivel()
	assert.Equal(t, NivelPorOmision, nivel)

}

func TestAumentarNivel(t *testing.T) {

	//Aumentar dentro de los límites
	nivel := NewNivel()
	nivel = nivel.AumentarNivel()
	assert.Greater(t, nivel, NivelPorOmision)

	//Aumentamos más allá de los límites
	nivel2 := NivelMaximo
	nivel2 = nivel2.AumentarNivel()
	assert.Equal(t, NivelMaximo, nivel2)
}

func TestDisminuirNivel(t *testing.T) {

	//Disminuir dentro de los límites
	nivel := NewNivel()
	nivel = nivel.DisminuirNivel()
	assert.Less(t, nivel, NivelPorOmision)

	//Disminuir más allá de los límites
	nivel2 := NivelMinimo
	nivel2 = nivel2.DisminuirNivel()
	assert.Equal(t, NivelMinimo, nivel2)
}
