package models

import (
	"errors"
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

// FindLikeByID returns a Like by id
func (l *Like) FindLikeByID(db *gorm.DB) (*Like, error) {
	err := db.Where("id=?", l.ID).Take(&l).Error
	if err == gorm.ErrRecordNotFound {
		return &Like{}, errors.New("Like not found")
	} else if err != nil {
		return nil, err
	}
	return l, nil
}

// FindLikesBy returns a list of Like by user id or post id.
func (l *Like) FindLikesBy(db *gorm.DB) (*[]Like, error) {
	var (
		likes []Like
		err   error
	)
	if l.UserID != 0 {
		err = db.Where("user_id=?", l.UserID).Find(&likes).Error
	} else if l.PostID != 0 {
		err = db.Where("post_id=?", l.PostID).Find(&likes).Error
	}
	if err == gorm.ErrRecordNotFound {
		return &[]Like{}, errors.New("Like not found")
	} else if err != nil {
		return nil, err
	}
	return &likes, nil
}
