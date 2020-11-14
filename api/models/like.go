package models

import (
	"time"

	"gorm.io/gorm"
)

// Like A user like a post
type Like struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	PostID    uint64    `gorm:"not null;index:idx_post_id" json:"post_id"`
	UserID    uint64    `gorm:"not null;index:idx_user_id" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   gorm.DeletedAt
}
