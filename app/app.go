package app

import (
	"fmt"
	"net/http"

	"github.com/florgm/webplanner_api/src/api/controllers/eventos"
	"github.com/florgm/webplanner_api/src/api/controllers/tareas"
	"github.com/florgm/webplanner_api/src/api/controllers/usuarios"
	"github.com/florgm/webplanner_api/src/api/services/sessions"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = SetupRouter()
)

func Run() {
	if err := router.Run(:8080); err != nil {
		fmt.Println(err.Error())
	}
}

func SetupRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/login", usuarios.Login)
	router.POST("/usuarios", usuarios.CreateUsuario)

	auth := router.Group("/auth")
	auth.Use(AuthRequired)
	{
		auth.GET("/productos", productos.AddProdcuto)
		auth.GET("/code", controller.GetToken)
	}

	return router
}
