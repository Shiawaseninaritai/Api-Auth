package initializers

import "github.com/Shiawaseninaritai/Api-Auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
