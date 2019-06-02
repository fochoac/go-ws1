package main

import (
	"fmt"

	model "github.com/fochoac/go-model-ws/persona"
	"github.com/fochoac/go-ws1/service"
)

func main() {
	fmt.Println("************************ consulta orm 1 ***************")

	var personas []model.Persona = service.ListarPersonasActivas()
	for _, persona := range personas {
		fmt.Println(persona)
	}

	fmt.Println("************************ consulta nativa 1 ***************")
	var personasNativas []model.PersonaNativa = service.ListarPersonasActivasNativo()
	for _, persona := range personasNativas {
		fmt.Println(persona)
	}

	fmt.Println("************************ consulta nativa 2 ***************")
	var nombrePersonasNativas []model.PersonaNativa = service.ListarNombrePersonasActivasNativo()
	for _, persona := range nombrePersonasNativas {
		fmt.Println(persona)
	}
}
