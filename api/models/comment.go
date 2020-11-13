package models

import (
	"time"

	"gorm.io/gorm"
)

// Comment comments for a post.
type Comment struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:256;not null;index:idx_title" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	PostID    uint      `gorm:"not null;index:idx_post_id" json:"post_id"`
	UserID    uint      `gorm:"not null;index:idx_user_id" json:"user_id"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   gorm.DeletedAt
}
