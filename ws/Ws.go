package ws1

import (
	"net/http"

	"github.com/fochoac/go-ws1/service"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Iniciar(router1 *gin.Engine) {
	router = router1
	initializeRoutes()

}

func initializeRoutes() {
	router.POST("/api/usuario/listarTodos", ListarPersonasActivasNativo)
	router.POST("/api/usuario/listarNombres", ListarNombrePersonasActivasNativo)
	router.POST("/api/usuario/listarUsuariosORM", ListarPersonasActivas)

}

func ListarPersonasActivasNativo(c *gin.Context) {

	c.JSON(http.StatusOK, service.ListarPersonasActivasNativo)
}

func ListarNombrePersonasActivasNativo(c *gin.Context) {

	c.JSON(http.StatusOK, service.ListarNombrePersonasActivasNativo)
}
func ListarPersonasActivas(c *gin.Context) {

	c.JSON(http.StatusOK, service.ListarPersonasActivas)
}
