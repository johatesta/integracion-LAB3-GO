package app

import (
	"fmt"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"integracion-LAB3-GO/controller"

)

func Run(){
	router := gin.Default()
	router.GET("/auth", controller.GetToken)
	router.GET("/items/all", controller.GetItem)
	router.Run()
}
