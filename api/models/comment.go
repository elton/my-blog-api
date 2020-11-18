package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Comment comments for a post.
type Comment struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:256;not null;index:idx_title" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	PostID    uint64    `gorm:"not null;index:idx_post_id" json:"post_id"`
	UserID    uint64    `gorm:"not null;index:idx_user_id" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   gorm.DeletedAt
}

// SaveComment Create a comment.
// curl -i -X POST \
//   http://127.0.0.1:8080/api/v1/comments/\?pid\=1\&uid\=1 \
//   -H 'cache-control: no-cache' \
//   -H 'content-type: application/json' \
//   -d '{
//         "title":"comment title II","content":"comment content","post_id":1,"user_id":1
// }'
func (c *Comment) SaveComment(db *gorm.DB) (*Comment, error) {
	if err := db.Create(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// FindCommentByID returns a comment by id.
func (c *Comment) FindCommentByID(db *gorm.DB) (*Comment, error) {
	err := db.Where("id=?", c.ID).Take(&c).Error
	if err == gorm.ErrRecordNotFound {
		return &Comment{}, errors.New("Comment not found")
	} else if err != nil {
		return nil, err
	}
	return c, nil
}

// FindCommentsBy returns a list of comments by criterias.
func (c *Comment) FindCommentsBy(db *gorm.DB) (*[]Comment, error) {
	var (
		comments []Comment
		err      error
	)
	if c.UserID != 0 {
		err = db.Where("user_id=?", c.UserID).Find(&comments).Error
	} else if c.PostID != 0 {
		err = db.Where("post_id=?", c.PostID).Find(&comments).Error
	}
	if err == gorm.ErrRecordNotFound {
		return &[]Comment{}, errors.New("Comment not found")
	} else if err != nil {
		return nil, err
	}

	return &comments, nil
}
