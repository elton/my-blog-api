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

// SaveLikes creates a new Like.
func (l *Like) SaveLikes(db *gorm.DB) (*Like, error) {
	if err := db.Create(&l).Error; err != nil {
		return nil, err
	}
	return l, nil
}
