package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// User Model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// Database Connect
func connectDB() {
	dsn := "host=localhost user=postgres password=Megha@420 dbname=productdb port=5432 sslmode=disable"
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
	// Public Routes — Sab access kar sakte hain
	r.POST("/register", Register)
	r.POST("/login", Login)

	// Protected Routes — Sirf login users!
	protected := r.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/profile", GetProfile)
	}

	r.Run(":8080")
}
