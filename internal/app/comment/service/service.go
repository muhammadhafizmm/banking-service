package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/muhammadhafizmm/banking-service/internal/app/comment/model"
)

func BindComment(c *gin.Context) (*model.Comment, error) {
	var comment model.Comment

	if err := c.BindJSON(&comment); err != nil {
		return nil, err
	}

	return &comment, nil
}

func FindCommentByID(c *gin.Context, comments []model.Comment) (*model.Comment, error) {
	id := c.Param("id")

	for _, v := range comments {
		if *v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("NOT FOUND")
}

func ReplaceCommentByID(c *gin.Context, comments []model.Comment, comment model.Comment) (*model.Comment, error) {
	id := c.Param("id")

	for i, v := range comments {
		if *v.ID == id {
			comments[i] = comment
			return &comment, nil
		}
	}

	return nil, errors.New("NOT FOUND")
}

func UpdateCommentByID(c *gin.Context, comments []model.Comment, comment model.Comment) (*model.Comment, error) {
	id := c.Param("id")

	for i, v := range comments {
		if *v.ID == id {
			if comment.Title != nil {
				comments[i].Title = comment.Title
			}
			if comment.Message != nil {
				comments[i].Message = comment.Message
			}
			if comment.Like != nil {
				comments[i].Like = comment.Like
			}
			return &comments[i], nil
		}
	}

	return nil, errors.New("NOT FOUND")
}
