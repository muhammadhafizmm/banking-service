package rest

import (
	"github.com/gin-gonic/gin"
	commentServer "github.com/muhammadhafizmm/banking-service/internal/app/comment/server"
)

func StartServer() {
	router := gin.Default()
	InitRouter(router)

	router.Run("localhost:8080")
}

func InitRouter(router *gin.Engine) {
	apiRouter := router.Group("/v1/api")

	commentServer.CommentRoutes.NewRoutes(apiRouter.Group("/comment"))
}
