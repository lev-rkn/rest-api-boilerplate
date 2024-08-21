package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidPageNumber = errors.New("invalid page number")
	ErrInvalidUserId     = errors.New("invalid user id")
	ErrInvalidDateSort   = errors.New("invalid date sort parameter")
	ErrInvalidArticleId  = errors.New("invalid id of the article")
	ErrInvalidToken      = errors.New("invalid access token")
)

func errorResponse(c *gin.Context, code int, err error) {
	var errMessage string
	unwrapped := errors.Unwrap(err)
	if unwrapped == nil {
		errMessage = err.Error()
	} else {
		errMessage = unwrapped.Error()
	}

	c.JSON(code, gin.H{
		"error": errMessage,
	})
}
