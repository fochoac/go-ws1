package ws

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
	personas := service.ListarPersonasActivasNativo()
	c.JSON(http.StatusOK, personas)
}

func ListarNombrePersonasActivasNativo(c *gin.Context) {
	personas := service.ListarNombrePersonasActivasNativo()
	c.JSON(http.StatusOK, personas)
}
func ListarPersonasActivas(c *gin.Context) {
	personas := service.ListarPersonasActivas()
	c.JSON(http.StatusOK, personas)
}
