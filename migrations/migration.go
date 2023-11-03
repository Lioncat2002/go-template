package main

import (
	"backend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}
	dburi := os.Getenv("DB_URI") //used a cockroachdb database but postgres is fine
	db, err := gorm.Open(postgres.Open(dburi), &gorm.Config{})
	if err != nil {
		log.Println("Error coonection to db", err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Post{})
	db.Exec("DROP TABLE user_likes")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserLikes{})
	db.AutoMigrate(&models.Post{})
	log.Println("Successfully migrated")
}
