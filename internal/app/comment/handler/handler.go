package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadhafizmm/banking-service/internal/app/comment/helper"
	"github.com/muhammadhafizmm/banking-service/internal/app/comment/model"
	"github.com/muhammadhafizmm/banking-service/internal/app/comment/service"
)

func strPtr(s string) *string { return &s }
func intPtr(i int) *int       { return &i }

var comments = []model.Comment{
	{ID: strPtr("1"), Title: strPtr("Hello World"), Message: strPtr("Lorem Ipsum"), Like: intPtr(0)},
	{ID: strPtr("2"), Title: strPtr("Hello World"), Message: strPtr("Lorem Ipsum"), Like: intPtr(0)},
	{ID: strPtr("3"), Title: strPtr("Hello World"), Message: strPtr("Lorem Ipsum"), Like: intPtr(0)},
	{ID: strPtr("4"), Title: strPtr("Hello World"), Message: strPtr("Lorem Ipsum"), Like: intPtr(0)},
}

func GetComment(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, comments)
}

func GetCommentByID(c *gin.Context) {
	comment, err := service.FindCommentByID(c, comments)

	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "GetCommentByID: 'ID' not found!"},
		)
		return
	}

	c.IndentedJSON(http.StatusOK, &comment)
}

func PostComment(c *gin.Context) {
	comment, err := service.BindComment(c)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "PostComment: 'param' error"})
		return
	}

	err = helper.ValidateComment(*comment)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("PostComment: %q", err)})
		return
	}

	comments = append(comments, *comment)
	c.IndentedJSON(http.StatusCreated, &comment)
}

func PutCommentByID(c *gin.Context) {
	comment, err := service.BindComment(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "UpdateCommentByID: 'param' error"})
		return
	}

	err = helper.ValidateComment(*comment)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("PostComment: %q", err)})
		return
	}

	replaceComment, err := service.ReplaceCommentByID(c, comments, *comment)
	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "UpdateCommentByID: 'ID' not found!"},
		)
		return
	}

	c.IndentedJSON(http.StatusAccepted, &replaceComment)
}

func UpdateCommentByID(c *gin.Context) {
	comment, err := service.BindComment(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "UpdateCommentByID: 'param' error"})
		return
	}

	updateComment, err := service.UpdateCommentByID(c, comments, *comment)
	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "UpdateCommentByID: 'ID' not found!"},
		)
		return
	}

	c.IndentedJSON(http.StatusAccepted, &updateComment)

}
