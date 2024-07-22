package server

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadhafizmm/banking-service/internal/app/comment/handler"
)

type commentRoutes struct{}

var CommentRoutes commentRoutes

func (commentRoutes) NewRoutes(router *gin.RouterGroup) {
	router.GET("", handler.GetComment)
	router.GET(":id", handler.GetCommentByID)
	router.POST("", handler.PostComment)
	router.PUT(":id", handler.PutCommentByID)
	router.PATCH(":id", handler.UpdateCommentByID)
}
