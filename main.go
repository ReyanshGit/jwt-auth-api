package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func connectDB() {
	// .env load karo
	godotenv.Load()

	// .env se values lo
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connect nahi hua!")
	}
	db = database
	db.AutoMigrate(&User{})
}

func main() {
	connectDB()

	r := gin.Default()

	r.POST("/register", Register)
	r.POST("/login", Login)

	protected := r.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/profile", GetProfile)
	}

	r.Run(":8080")
}
