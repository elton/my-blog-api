package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user.
type User struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:256;not null;uniqueIndex:idx_nickname;" json:"nickname"`
	Type      uint8     `gorm:"not null;index:idx_type" json:"type"`
	Mobile    string    `gorm:"size:16;not null;uniqueIndex:idx_mobile;" json:"mobile"`
	Email     string    `gorm:"size:128;not null;uniqueIndex:idx_email;" json:"email"`
	Posts     []Post    `json:"posts"`
	Likes     []Like    `json:"likes"`
	Comments  []Comment `json:"comments"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   gorm.DeletedAt
}
