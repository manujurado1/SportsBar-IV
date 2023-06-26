package modelos

import (
	"fmt"
	"time"
)

var (
	ErrorNickVacio = fmt.Errorf("El nick de un amigo no puede ser un string vacío")
)

type Amigo struct {
	identificador string
}

func NewAmigo(nick string, fechaNacimiento time.Time) (Amigo, error) {
	if nick == "" {
		return Amigo{}, ErrorNickVacio
	}

	identificador := GenerarIdentificador(nick, fechaNacimiento)

	return Amigo{
		identificador: identificador,
	}, nil

}

func (a Amigo) ObtenerId() string {
	return a.identificador
}

func GenerarIdentificador(nickAmigo string, fechaNacimiento time.Time) string {
	return nickAmigo + fmt.Sprint(fechaNacimiento.Day()) + fmt.Sprint(fechaNacimiento.Month())
}
