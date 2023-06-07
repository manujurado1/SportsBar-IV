package modelos

import "fmt"

var (
	ErrorNombreVacio = fmt.Errorf("El nombre de un jugador no puede ser un string vacío")
)

const (
	DisponibilidadPorDefecto bool = true
)

type Amigo struct {
	nombreDentroDelGrupo string
	nivel                Nivel
	disponible           bool
}

func NewAmigo(nombre string) (Amigo, error) {
	if nombre == "" {
		return Amigo{}, ErrorNombreVacio
	}

	return Amigo{
		nombreDentroDelGrupo: nombre,
		nivel:                NewNivel(),
		disponible:           DisponibilidadPorDefecto,
	}, nil
}

func (a Amigo) ObtenerNombreDentroDelGrupo() string {
	return a.nombreDentroDelGrupo
}

func (a Amigo) ObtenerNivel() Nivel {
	return a.nivel
}

func (a Amigo) EstaDisponible() bool {
	return a.disponible
}

func (a Amigo) CambiarDisponibilidad(disponible bool) Amigo {
	return Amigo{
		nombreDentroDelGrupo: a.ObtenerNombreDentroDelGrupo(),
		nivel:                a.ObtenerNivel(),
		disponible:           disponible,
	}

}

func (a Amigo) AumentarNivel() Amigo {
	return Amigo{
		nombreDentroDelGrupo: a.ObtenerNombreDentroDelGrupo(),
		nivel:                a.nivel.AumentarNivel(),
		disponible:           a.EstaDisponible(),
	}
}

func (a Amigo) DisminuirNivel() Amigo {
	return Amigo{
		nombreDentroDelGrupo: a.ObtenerNombreDentroDelGrupo(),
		nivel:                a.nivel.DisminuirNivel(),
		disponible:           a.EstaDisponible(),
	}
}
