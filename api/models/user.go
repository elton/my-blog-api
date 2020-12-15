package models

import (
	"errors"
	"time"

	"github.com/elton/my-blog-api/api/utils"
	"gorm.io/gorm"
)

// User represents a user.
type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:256;not null;uniqueIndex:idx_userame;" json:"username"`
	Password  string    `gorm:"size:256;not null" json:"password"`
	Nickname  string    `gorm:"size:256;uniqueIndex:idx_nickname;" json:"nickname"`
	Type      uint8     `gorm:"not null;index:idx_type;default:1;comment:1:user, 2:admin" json:"type"`
	Mobile    string    `gorm:"size:16;not null;uniqueIndex:idx_mobile;" json:"mobile"`
	Email     string    `gorm:"size:128;not null;uniqueIndex:idx_email;" json:"email"`
	Posts     []Post    `json:"posts"`
	Likes     []Like    `json:"likes"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   gorm.DeletedAt
}

// Validate validates the user's struct.
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("Username field for user is required")
	}
	if u.Password == "" {
		return errors.New("Password field for user is required")
	}
	if u.Mobile == "" {
		return errors.New("Mobile field for user is required")
	}
	if u.Email == "" {
		return errors.New("Email field for user is required")
	}
	if u.Type != 0 && u.Type != 1 && u.Type != 2 {
		return errors.New("Type for user is invalid")
	}
	return nil
}

// SaveUser create a new user.
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	hashedpwd, err := utils.HashAndSalt([]byte(u.Password))
	if err != nil {
		return nil, err
	}

	u.Password = string(hashedpwd)
	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// FindUserByID returns a user matching specific ID.
func (u *User) FindUserByID(db *gorm.DB) (*User, error) {
	err := db.Where("id=?", u.ID).Take(&u).Error
	if err == gorm.ErrRecordNotFound {
		return &User{}, errors.New("User not found")
	} else if err != nil {
		return nil, err
	}
	return u, nil
}

// FindUsers returns a list of top 100 users.
func (u *User) FindUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	if err := db.Limit(100).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

// FindUsersBy returns a list of users matching the specific conditions.
func (u *User) FindUsersBy(db *gorm.DB) (*[]User, error) {
	users := []User{}
	var err error

	switch {
	case u.Username != "":
		err = db.Where("username=?", u.Username).Find(&users).Error
	case u.Nickname != "":
		err = db.Where("nickname like ?", "%"+u.Nickname+"%").Find(&users).Error
	case u.Type == 1 || u.Type == 2:
		err = db.Where("type = ?", u.Type).Find(&users).Error
	case u.Mobile != "":
		err = db.Where("mobile = ?", u.Mobile).Find(&users).Error
	case u.Email != "":
		err = db.Where("email = ?", u.Email).Find(&users).Error
	}

	if err == gorm.ErrRecordNotFound || len(users) <= 0 {
		return &[]User{}, gorm.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}
	return &users, nil
}

// UpdateUser updates a user.
func (u *User) UpdateUser(db *gorm.DB) error {
	if u.Password != "" {
		hashedpwd, err := utils.HashAndSalt([]byte(u.Password))
		if err != nil {
			return err
		}

		u.Password = string(hashedpwd)
	}
	// Updates只更新不为空的字段
	if err := db.Updates(&u).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user.
func (u *User) DeleteUser(db *gorm.DB) error {
	if err := db.Delete(&User{}, u.ID).Error; err != nil {
		return err
	}
	return nil
}
