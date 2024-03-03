package model

import (
	gorm "gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string `gorm:"size:100; not null;"`
	Email    string `gorm:"size:100; not null; unique"`
	Username string `gorm:"size:255; not null; unique"`
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GetUser(db *gorm.DB, uid uint) (*User, error) {
	err := db.Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
