package modelos

import "time"

type Equipo struct {
	listaNombreAmigoDentoDelGrupo []string
	fechaCreacion                 time.Time
}

func NewEquipo() Equipo {
	return Equipo{
		listaNombreAmigoDentoDelGrupo: []string{},
		fechaCreacion:                 time.Now(),
	}
}

func (e Equipo) ObtenerEquipo() []string {
	return e.listaNombreAmigoDentoDelGrupo
}

func (e Equipo) ObtenerFechaCreacion() string {
	return e.fechaCreacion.Format("02-01-2006")
}

func (e Equipo) RellenarEquipo(nombreAmigoDentroDelGrupo []string) Equipo {
	return Equipo{
		listaNombreAmigoDentoDelGrupo: nombreAmigoDentroDelGrupo,
		fechaCreacion:                 e.fechaCreacion,
	}
}
