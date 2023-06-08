package modelos

import (
	"fmt"
	"strings"
)

var (
	ErrorNombreVacio = fmt.Errorf("El nombre de un jugador no puede ser un string vacío")
)

const (
	DisponibilidadPorDefecto bool = true
)

type Amigo struct {
	nombre        string
	diaNacimiento int
	mesNacimiento string
}

func NewAmigo(nombre string, diaNacimiento int, mesNacimiento string) (Amigo, error) {
	if nombre == "" {
		return Amigo{}, ErrorNombreVacio
	}

	return Amigo{
		nombre:        nombre,
		diaNacimiento: diaNacimiento,
		mesNacimiento: mesNacimiento,
	}, nil
}

func (a Amigo) ObtenerId() string {
	identificativo := a.nombre + fmt.Sprint(a.diaNacimiento) + strings.ToLower(a.mesNacimiento)
	return identificativo
}
