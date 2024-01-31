package initializers

import "github.com/robbyklein/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
