package modelos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNuevoAmigo(t *testing.T) {

	//Caso incorrecto por nombre vacío
	amigo, error := NewAmigo("", time.Date(1990, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	assert.EqualError(t, error, "El nick de un amigo no puede ser un string vacío")
	assert.Equal(t, Amigo{}, amigo)

	//Caso correcto, comprobando que 2 personas con el mismo nick tienen identificadores diferentes
	amigo2, error2 := NewAmigo("Juanito", time.Date(1999, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo3, error3 := NewAmigo("Juanito", time.Date(2001, time.January, 22, 0, 0, 0, 0, time.Now().Location()))

	assert.Nil(t, error2)
	assert.Nil(t, error3)

	assert.NotEqual(t, amigo2.ObtenerId(), amigo3.ObtenerId())

}
