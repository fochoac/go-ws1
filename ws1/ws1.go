package ws1

import (
	"net/http"

	"github.com/fochoac/go-util-ws/constantes"

	"github.com/fochoac/go-model-ws/prueba"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: username`
	Password string `json: password`
}

var router *gin.Engine

func Iniciar(router1 *gin.Engine) {
	router = router1
	initializeRoutes()

}

func initializeRoutes() {
	router.POST("/api/testPost", listarTag)
	router.GET("/api", handleGet)
}

func listarTag(c *gin.Context) {
	tag := prueba.Tag{
		Name:        "Ecuador",
		Description: "Pais",
		ExtraData: prueba.TagExtra{
			Extra: "Informacion adicional",
		},
	}
	datos := make([]prueba.Tag, 10)
	for index := 0; index < len(datos); index++ {
		datos[index] = tag
	}
	respuesta := prueba.RetornoTag{
		CodigoRetorno: constantes.Ok(),
		Tag:           datos,
	}
	c.JSON(http.StatusOK, respuesta)
}

func handleGet(c *gin.Context) {
	c.String(http.StatusOK, "ping")
}
