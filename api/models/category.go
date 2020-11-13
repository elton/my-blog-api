package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Category the category for post.
type Category struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:256;uniqueIndex:idx_name;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   gorm.DeletedAt
}

// Validate validates the category
func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.New("Required Name For Category")
	}
	return nil
}

// SaveCategory save category information to database.
func (c *Category) SaveCategory(db *gorm.DB) (*Category, error) {
	if err := db.Create(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
