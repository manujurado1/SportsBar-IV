package modelos

import (
	"fmt"
	"sort"
	"time"
)

const (
	DisponibilidadPorDefecto             bool = true
	CantidadAmigosMinimaParaCrearEquipos      = 10
	DiferenciaDeNivelMaximaEntreEquipos       = 1
)

var (
	ErrorNombreGrupoVacio                            = fmt.Errorf("Un grupo no puede tener como nombre un string vacío")
	ErrorAmigoDuplicado                              = fmt.Errorf("Ya existe un amigo con ese identificador")
	ErrorAmigoInexistente                            = fmt.Errorf("No existe el amigo indicado dentro del grupo de amigos")
	ErrorListaAmigosVacia                            = fmt.Errorf("No se puede crear un grupo de amigos con una lista de amigos vacía")
	ErrorJugadoresDisponiblesNoAptosParaCrearEquipos = fmt.Errorf("No se pueden crear 2 equipos igualados si la cantidad de amigos disponibles no es un número par mayor o igual a 10")
	ErrorImposibilidadCrearEquiposIgualados          = fmt.Errorf("Ha sido imposible crear 2 equipos igualados con la lista de jugadores disponibles")
)

func FormatearError(e error, identificador string) error {
	return fmt.Errorf(e.Error() + ": " + identificador)
}

var EstadoAmigoPorDefecto EstadoAmigo = EstadoAmigo{
	Nivel:      NivelPorOmision,
	Disponible: DisponibilidadPorDefecto,
}

type EstadoAmigo struct {
	Nivel      Nivel
	Disponible bool
}
type GrupoAmigos struct {
	NombreDelGrupo             string
	ListaAmigos                []Amigo
	NivelYDisponibilidadAmigos map[string]EstadoAmigo
}

func NewGrupoAmigos(nombre string, listaAmigos []Amigo) (*GrupoAmigos, error) {
	if nombre == "" {
		return nil, ErrorNombreGrupoVacio
	}
	if len(listaAmigos) == 0 {
		return nil, ErrorListaAmigosVacia
	}

	estadosAmigos := map[string]EstadoAmigo{}

	for _, amigo := range listaAmigos {
		estadosAmigos[amigo.ObtenerId()] = EstadoAmigoPorDefecto
	}

	return &GrupoAmigos{
		NombreDelGrupo:             nombre,
		ListaAmigos:                listaAmigos,
		NivelYDisponibilidadAmigos: estadosAmigos,
	}, nil
}

func (g *GrupoAmigos) CrearAmigoYAniadirAlGrupo(nickAmigo string, fechaNacimiento time.Time) error {
	amigo, errorAmigo := NewAmigo(nickAmigo, fechaNacimiento)

	if errorAmigo != nil {
		return errorAmigo
	}

	_, existe := g.NivelYDisponibilidadAmigos[amigo.ObtenerId()]

	if existe {
		return FormatearError(ErrorAmigoDuplicado, amigo.ObtenerId())
	}

	g.ListaAmigos = append(g.ListaAmigos, amigo)
	g.NivelYDisponibilidadAmigos[amigo.ObtenerId()] = EstadoAmigoPorDefecto

	return nil
}

func (g *GrupoAmigos) CambiarDisponibilidadAmigo(amigo Amigo, disponibilidad bool) error {
	estado, existe := g.NivelYDisponibilidadAmigos[amigo.ObtenerId()]

	if !existe {
		return FormatearError(ErrorAmigoInexistente, amigo.ObtenerId())
	}

	estado.Disponible = disponibilidad
	g.NivelYDisponibilidadAmigos[amigo.ObtenerId()] = estado

	return nil
}

func (g *GrupoAmigos) AumentarNivelAmigo(amigo Amigo) error {
	estado, existe := g.NivelYDisponibilidadAmigos[amigo.ObtenerId()]

	if !existe {
		return FormatearError(ErrorAmigoInexistente, amigo.ObtenerId())
	}

	estado.Nivel = estado.Nivel.AumentarNivel()
	g.NivelYDisponibilidadAmigos[amigo.ObtenerId()] = estado

	return nil
}

func (g *GrupoAmigos) DisminuirNivelAmigo(amigo Amigo) error {
	estado, existe := g.NivelYDisponibilidadAmigos[amigo.ObtenerId()]

	if !existe {
		return FormatearError(ErrorAmigoInexistente, amigo.ObtenerId())
	}

	estado.Nivel = estado.Nivel.DisminuirNivel()
	g.NivelYDisponibilidadAmigos[amigo.ObtenerId()] = estado

	return nil
}

func (g *GrupoAmigos) GrupoAmigosListoParaCrearEquipos(estadosAmigos map[string]EstadoAmigo) (bool, []string, error) {
	amigosDisponibles := g.ObtenerListaAmigosDisponibles(estadosAmigos)
	if (len(amigosDisponibles) < CantidadAmigosMinimaParaCrearEquipos) || (len(amigosDisponibles)%2 != 0) {
		return false, nil, ErrorJugadoresDisponiblesNoAptosParaCrearEquipos
	}

	return true, amigosDisponibles, nil

}

func (g *GrupoAmigos) ObtenerListaAmigosDisponibles(estadosAmigos map[string]EstadoAmigo) []string {
	amigosDisponibles := []string{}

	for nombre, estado := range estadosAmigos {
		if estado.Disponible {
			amigosDisponibles = append(amigosDisponibles, nombre)
		}
	}

	sort.SliceStable(amigosDisponibles, func(i, j int) bool {
		return estadosAmigos[amigosDisponibles[i]].Nivel > estadosAmigos[amigosDisponibles[j]].Nivel
	})

	return amigosDisponibles
}

func (g *GrupoAmigos) CrearDosEquiposIgualados(estadosAmigos map[string]EstadoAmigo) (Equipo, Equipo, error) {
	listo, listaAmigosDisponiblesOrdenados, errorListaAmigos := g.GrupoAmigosListoParaCrearEquipos(estadosAmigos)

	if !listo {
		return Equipo{}, Equipo{}, errorListaAmigos
	}

	equipo1, equipo2 := g.RepartirJugadoresDisponiblesEnDosEquiposSegunNivel(listaAmigosDisponiblesOrdenados)

	igualados := g.EstanIgualados(equipo1, equipo2)

	if !igualados {
		return Equipo{}, Equipo{}, ErrorImposibilidadCrearEquiposIgualados
	}

	return equipo1, equipo2, nil

}

func (g *GrupoAmigos) RepartirJugadoresDisponiblesEnDosEquiposSegunNivel(ListaAmigosDisponiblesOrdenados []string) (Equipo, Equipo) {
	var EquipoA []string
	var EquipoB []string
	var NivelTotalEquipoA uint
	var NivelTotalEquipoB uint

	for _, Amigo := range ListaAmigosDisponiblesOrdenados {

		if len(EquipoA) == 0 && len(EquipoB) == 0 {
			EquipoA = append(EquipoA, Amigo)
			NivelTotalEquipoA += uint(g.NivelYDisponibilidadAmigos[Amigo].Nivel)
		} else if len(EquipoB) == 0 {
			EquipoB = append(EquipoB, Amigo)
			NivelTotalEquipoB += uint(g.NivelYDisponibilidadAmigos[Amigo].Nivel)
		} else if (NivelTotalEquipoA < NivelTotalEquipoB) && (len(EquipoA) < (len(ListaAmigosDisponiblesOrdenados) / 2)) {
			EquipoA = append(EquipoA, Amigo)
			NivelTotalEquipoA += uint(g.NivelYDisponibilidadAmigos[Amigo].Nivel)
		} else if len(EquipoB) < (len(ListaAmigosDisponiblesOrdenados) / 2) {
			EquipoB = append(EquipoB, Amigo)
			NivelTotalEquipoB += uint(g.NivelYDisponibilidadAmigos[Amigo].Nivel)
		} else {
			EquipoA = append(EquipoA, Amigo)
			NivelTotalEquipoA += uint(g.NivelYDisponibilidadAmigos[Amigo].Nivel)
		}
	}

	Equipo := NewEquipo()
	PrimerEquipo := Equipo.RellenarEquipo(EquipoA)
	SegundoEquipo := Equipo.RellenarEquipo(EquipoB)

	return PrimerEquipo, SegundoEquipo

}

func (g *GrupoAmigos) EstanIgualados(Equipo1 Equipo, Equipo2 Equipo) bool {
	Igualados := false
	var NivelTotalEquipo1 uint = 0
	var NivelTotalEquipo2 uint = 0

	if len(Equipo1.listaNombreAmigoDentoDelGrupo) == len(Equipo2.listaNombreAmigoDentoDelGrupo) && len(Equipo1.listaNombreAmigoDentoDelGrupo) >= 5 {

		for _, jugador := range Equipo1.ObtenerEquipo() {
			NivelTotalEquipo1 += uint(g.NivelYDisponibilidadAmigos[jugador].Nivel)
		}

		for _, jugador := range Equipo2.ObtenerEquipo() {
			NivelTotalEquipo2 += uint(g.NivelYDisponibilidadAmigos[jugador].Nivel)
		}

		NivelEquipo1 := float64(NivelTotalEquipo1) / float64(len(Equipo1.ObtenerEquipo()))
		NivelEquipo2 := float64(NivelTotalEquipo2) / float64(len(Equipo2.ObtenerEquipo()))

		if (NivelEquipo1+DiferenciaDeNivelMaximaEntreEquipos > NivelEquipo2) && (NivelEquipo2+DiferenciaDeNivelMaximaEntreEquipos > NivelEquipo1) {
			Igualados = true
		}

	}

	return Igualados
}
