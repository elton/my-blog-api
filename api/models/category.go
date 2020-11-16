package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Category the category for post.
type Category struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:256;uniqueIndex:idx_name;not null" json:"name"`
	Posts     []Post    `gorm:"many2many:post_categories" json:"posts"`
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

// FindCategoryByID find a specific category by ID.
func (c *Category) FindCategoryByID(db *gorm.DB) (*Category, error) {
	err := db.Where("id = ?", c.ID).Take(&c).Error
	if err == gorm.ErrRecordNotFound {
		return &Category{}, errors.New("Category Not Found")
	} else if err != nil {
		return &Category{}, err
	}
	return c, nil
}

// FindCategories returns a list of categories in first 100 results.
func (c *Category) FindCategories(db *gorm.DB) (*[]Category, error) {
	categories := []Category{}
	if err := db.Limit(100).Find(&categories).Error; err != nil {
		return &[]Category{}, err
	}
	return &categories, nil
}

// FindCategoriesByName returns a list of categories by name.
func (c *Category) FindCategoriesByName(db *gorm.DB, name string) (*[]Category, error) {
	categories := []Category{}
	err := db.Where("name like ?", "%"+name+"%").Find(&categories).Error
	if err == gorm.ErrRecordNotFound || len(categories) <= 0 {
		return &[]Category{}, errors.New("Category Not Found")
	} else if err != nil {
		return &[]Category{}, err
	}
	return &categories, nil
}

// UpdateCategory updates a category.
func (c *Category) UpdateCategory(db *gorm.DB) error {

	if err := db.Updates(&c).Error; err != nil {
		return err
	}
	return nil
}

// Delete a category.
func (c *Category) Delete(db *gorm.DB) error {
	if err := db.Delete(&Category{}, c.ID).Error; err != nil {
		return err
	}
	return nil
}
