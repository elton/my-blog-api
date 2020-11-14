package models

import (
	"time"

	"gorm.io/gorm"
)

// Post post struct
type Post struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Title      string     `gorm:"size:256;not null;index:idx_title" json:"title"`
	Summary    string     `gorm:"not null" json:"summary"`
	Content    string     `gorm:"not null" json:"content"`
	Categories []Category `gorm:"not null;many2many:post_categories" json:"categories"`
	Comments   []Comment  `json:"comments"`
	Likes      []Like     `json:"likes"`
	UserID     uint64     `gorm:"not null;index:idx_user_id" json:"user_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Deleted    gorm.DeletedAt
}
