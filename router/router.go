package router

import (
	"coba-gin/controller"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	controller.NewCarController().CarRoutes(r)

	return r
}
