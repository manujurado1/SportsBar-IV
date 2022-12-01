package pkg

type Grupo struct{
  Nombre string
  JugadoresNombres []*string
  JugadoresDisponibles map[string]string
  JugadoresNiveles map[string]int
}
