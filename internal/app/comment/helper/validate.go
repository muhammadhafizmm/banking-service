package helper

import (
	"errors"

	"github.com/muhammadhafizmm/banking-service/internal/app/comment/model"
)

func ValidateComment(comment model.Comment) error {
	if comment.ID == nil || comment.Title == nil || comment.Message == nil || comment.Like == nil {
		return errors.New("one or more required fields are null")
	}
	return nil
}
