package comment

import "github.com/jinzhu/gorm"

// Service - the srtruct for the comment service
type Service struct {
	DB *gorm.DB
}

// Comment - defines the comment structure
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

// CommentService - the interface for the comment serevice
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comment serevice
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
