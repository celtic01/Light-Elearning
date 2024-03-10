package model

import (
	"html"
	"strings"
	"time"

	"github.com/celtic01/Light-Elearning/api/security"
	gorm "gorm.io/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Password  string    `gorm:"size:100; not null;"`
	Email     string    `gorm:"size:100; not null; unique"`
	Username  string    `gorm:"size:255; not null; unique"`
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GetUser(db *gorm.DB, uid uint) (*User, error) {
	err := db.Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateUser(db *gorm.DB, uid uint, updateUser *User) (*User, error) {
	hashedPassword, err := security.Hash(updateUser.Password)

	if err != nil {
		return &User{}, err
	}

	updateUser.Password = string(hashedPassword)

	err = db.Model(&User{}).Where("id = ?", uid).Updates(updateUser).Error
	if err != nil {
		return &User{}, err
	}

	err = db.First(u, uid).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint) (int64, error) {
	db = db.Delete(&User{}, uid)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
