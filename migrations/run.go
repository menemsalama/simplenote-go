package migrations

import (
	"github.com/menemsalama/simplenote-go/internal/database"
	"github.com/menemsalama/simplenote-go/models"
)

// Run runs all necessary migrations
// TODO: handle migrating fields, cause according to https://gorm.io/docs/migration.html
// migration is only about creating tables
func Run() {

	// database.PG.DropTable(&models.User{}, &models.Note{})
	database.PG.AutoMigrate(&models.User{}, &models.Note{})

	database.PG.Model(&models.Note{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
}
