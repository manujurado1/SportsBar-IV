package modelos

import "fmt"

var (
	ErrorNombreVacio = fmt.Errorf("El nombre de un jugador no puede ser un string vacío")
)

type Amigo struct {
	nombre     string //Nombre o apodo por el que se le conoce a ese amigo dentro del grupo de amigos
	nivel      Nivel  //Nivel que posee ese amigo dentro del grupo de amigos
	disponible bool   //Atributo que indica si está disponible en ese momento para jugar un partido
}

func NuevoAmigo(nombre string, nivel Nivel, disponible bool) (Amigo, error) {
	if nombre == "" {
		return Amigo{}, ErrorNombreVacio
	}

	nivel_amigo, error := NuevoNivel(nivel)

	if error != nil {
		return Amigo{}, error
	}

	return Amigo{
		nombre:     nombre,
		nivel:      nivel_amigo,
		disponible: disponible,
	}, nil
}

func (a Amigo) ObtenerNombre() string {
	return a.nombre
}

func (a Amigo) ObtenerNivel() Nivel {
	return a.nivel
}

func (a Amigo) EstaDisponible() bool {
	return a.disponible
}
