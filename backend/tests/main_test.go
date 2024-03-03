package tests

import (
	"testing"

	model "github.com/celtic01/Light-Elearning/api/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func TestMain(m *testing.M) {
	DB = model.InitializeDB()
	m.Run()
}

func refreshUserTable() error {
	err := DB.Migrator().DropTable(&model.User{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}

func seedOneUser() (model.User, error) {
	user := model.User{
		Username: "testusser",
		Email:    "one@email.com",
		Password: "password",
	}
	err := DB.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
