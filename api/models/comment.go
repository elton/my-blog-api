package models

import (
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
