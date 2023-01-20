package modelos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearGrupo(t *testing.T) {
	grupo, error := CrearGrupo("")
	assert.Nil(t, grupo)
	assert.Errorf(t, error, "El nombre del grupo no puede ser un string vacío")

	grupo, error = CrearGrupo("Grupo1")
	assert.Nil(t, error)
	assert.Equal(t, "Grupo1", grupo.Nombre)
	assert.Zero(t, len(grupo.JugadoresDisponibilidad))
	assert.Zero(t, len(grupo.JugadoresNiveles))
}
