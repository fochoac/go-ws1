package service

import (
	model "github.com/fochoac/go-model-ws/persona"
	"github.com/fochoac/go-util-ws/constantes"
	dao "github.com/fochoac/go-ws1/dao"
)

func ListarPersonasActivas() model.PersonaRespuesta {

	var personas []model.Persona = dao.ListarTodasPersonas(true)

	if len(personas) > 0 {
		return model.PersonaRespuesta{
			CodigoRetorno:   constantes.Ok(),
			Personas:        personas,
			PersonasNativas: nil,
		}
	} else {
		return model.PersonaRespuesta{
			CodigoRetorno:   constantes.NoData(),
			Personas:        nil,
			PersonasNativas: nil,
		}
	}
}
func ListarPersonasActivasNativo() model.PersonaRespuesta {

	var personas []model.PersonaNativa = dao.ListarTodasPersonasNativo(true)
	if len(personas) > 0 {
		return model.PersonaRespuesta{
			CodigoRetorno:   constantes.Ok(),
			Personas:        nil,
			PersonasNativas: personas,
		}
	} else {
		return model.PersonaRespuesta{
			CodigoRetorno:   constantes.NoData(),
			Personas:        nil,
			PersonasNativas: nil,
		}
	}

}
func ListarNombrePersonasActivasNativo() model.PersonaRespuesta {

	var personas []model.PersonaNativa = dao.ListarTodasPersonasNombresNativo(true)

	if len(personas) > 0 {
		return model.PersonaRespuesta{
			CodigoRetorno:   constantes.Ok(),
			Personas:        nil,
			PersonasNativas: personas,
		}
	} else {
		return model.PersonaRespuesta{
			CodigoRetorno:   constantes.NoData(),
			Personas:        nil,
			PersonasNativas: nil,
		}
	}
}
