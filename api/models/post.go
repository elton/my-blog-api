package models

import (
	"errors"
	"strconv"
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

// Validate validates the post data.
func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("Title is required for Post")
	}

	if p.Summary == "" {
		return errors.New("Summary is required for Post")
	}

	if p.Content == "" {
		return errors.New("Content is required for Post")
	}

	if strconv.Itoa(int(p.UserID)) == "" {
		return errors.New("UserID is required for Post")
	}
	return nil
}

// SavePost save a post data in database.
func (p *Post) SavePost(db *gorm.DB, uid uint64, cids []uint64) (*Post, error) {
	p.UserID = uid
	for _, cid := range cids {
		var category Category
		category.ID = cid
		categoryGotten, err := category.FindCategoryByID(db)
		if err != nil {
			return nil, err
		}
		p.Categories = append(p.Categories, *categoryGotten)
	}

	if err := db.Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// FindPostByID find posts by post id.
func (p *Post) FindPostByID(db *gorm.DB) (*Post, error) {
	err := db.Preload("Categories").Find(&p).Error
	if err == gorm.ErrRecordNotFound {
		return &Post{}, errors.New("Post not found")
	} else if err != nil {
		return nil, err
	}
	return p, nil
}

// FindPostsByUser find posts by user.
func (p *Post) FindPostsByUser(db *gorm.DB, uid uint64) (*[]Post, error) {
	var (
		posts []Post
	)

	if err := db.Where("user_id= ?", uid).Order("updated_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

// FindPostsByCategory find posts by specific category.
func (p *Post) FindPostsByCategory(db *gorm.DB, cid uint64) (*[]Post, error) {
	var (
		category Category
		posts    []Post
	)
	if err := db.Where("id=?", cid).Find(&category).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&category).Order("updated_at desc").Preload("Categories").Association("Posts").Find(&posts); err != nil {
		return nil, err
	}
	return &posts, nil
}

// FindPostsByTitle return posts by title.
func (p *Post) FindPostsByTitle(db *gorm.DB, title string) (*[]Post, error) {
	var posts []Post
	err := db.Where("title like ?", "%"+title+"%").Find(&posts).Error
	if err == gorm.ErrRecordNotFound || len(posts) <= 0 {
		return &[]Post{}, errors.New("Posts not found")
	} else if err != nil {
		return nil, err
	}
	return &posts, nil
}

// FindPosts returns the top 100 posts.
func (p *Post) FindPosts(db *gorm.DB) (*[]Post, error) {
	var posts []Post
	if err := db.Limit(100).Order("updated_at desc").Preload("Categories").Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}
