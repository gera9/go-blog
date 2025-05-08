package sharedmodels

import (
	"github.com/google/uuid"
)

type Post struct {
	Id       uuid.UUID
	Title    string
	Extract  string
	Text     string
	ImgPath  string
	AuthorId uuid.UUID
}
