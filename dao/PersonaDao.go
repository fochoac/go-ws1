package dao

import (
	"log"

	model "github.com/fochoac/go-model-ws/persona"

	"github.com/fochoac/go-ws1/resources"
	"github.com/jinzhu/gorm"
)

const (
	SQL_LISTAR_PERSONAS = "select cast(p.persona_pk as text) id,p.cedula,p.nombres,	p.apellidos,to_char(p.fechanacimiento, 'dd/MM/yyyy') fechaNacimiento, " +
		" p.genero, to_char(p.fechacreacion, 'dd/MM/yyyy HH24:mm:ss') fechaCreacion, p.estado::text estado " +
		" from 	public.personas p where p.estado = ?"
	SQL_LISTAR_NOMBRE_PERSONAS = "select cast(p.persona_pk as text) id,p.cedula,p.nombres	" +
		" from 	public.personas p where p.estado = ?"
)
const HABILITAR_DEBUG = false

var db *gorm.DB

func ListarTodasPersonas(estado bool) []model.Persona {
	db = resources.GetConexion()

	defer db.Close()
	db.LogMode(HABILITAR_DEBUG)

	var personas []model.Persona
	db.Model(&model.Persona{}).Select("*").Where("estado = ?", estado).Find(&personas)

	return personas
}

func ListarTodasPersonasNativo(estado bool) []model.PersonaNativa {
	db = resources.GetConexion()

	defer db.Close()
	db.LogMode(HABILITAR_DEBUG)
	var personas []model.PersonaNativa //= make([]model.PersonaNativa, 15)
	rows, err := db.Raw(SQL_LISTAR_PERSONAS, true).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		persona := new(model.PersonaNativa)
		errr := rows.Scan(&persona.PersonaPk, &persona.Cedula, &persona.Nombres,
			&persona.Apellidos, &persona.FechaNacimiento, &persona.Genero,
			&persona.FechaCreacion, &persona.Estado)
		if errr != nil {
			log.Fatal(err)
			continue
		}
		personas = append(personas, *persona)

	}

	return personas
}

func ListarTodasPersonasNombresNativo(estado bool) []model.PersonaNativa {
	db = resources.GetConexion()

	defer db.Close()
	db.LogMode(HABILITAR_DEBUG)
	var personas []model.PersonaNativa //= make([]model.PersonaNativa, 15)
	rows, err := db.Raw(SQL_LISTAR_NOMBRE_PERSONAS, true).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		persona := new(model.PersonaNativa)
		errr := rows.Scan(&persona.PersonaPk, &persona.Cedula, &persona.Nombres)
		if errr != nil {
			log.Fatal(err)
			continue
		}
		personas = append(personas, *persona)

	}

	return personas
}
