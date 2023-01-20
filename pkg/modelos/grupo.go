package modelos

type Grupo struct {
	Nombre               string
	JugadoresDisponibles []string
	JugadoresNiveles     map[string]int
}
