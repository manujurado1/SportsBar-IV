package modelos

type Grupo struct {
	Nombre                  string
	JugadoresDisponibilidad map[string]bool
	JugadoresNiveles        map[string]int
}
