package models

import (
	"time"

	"gorm.io/gorm"
)

// Category the category for post.
type Category struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:256;index:idx_name;not null" json:"name"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   gorm.DeletedAt
}
