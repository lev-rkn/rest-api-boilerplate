package domain

import (
	"errors"
	"time"
)

var (
	ErrArticleNotFound   = errors.New("article not found")
)

type Article struct {
	Id          int       `json:"id,omitempty"`
	UserId      int       `json:"user_id,omitempty"`
	Title       string    `json:"title,omitempty" validate:"required,max=140"`
	Description string    `json:"description,omitempty" validate:"required,max=1000"`
	Photos      []string  `json:"photos,omitempty" validate:"required,max=3"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}
