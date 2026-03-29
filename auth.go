package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Secret Key
var secretKey = []byte("reyansh-secret-key-2024")

// Register Handler
func Register(c *gin.Context) {
	var user User
	c.Bind(&user)

	// Password Hash karo
	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// Database mein save karo
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Email already hai!"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Register ho gaye Reyansh bhai!",
		"name":    user.Name,
		"email":   user.Email,
	})
}

// Login Handler
func Login(c *gin.Context) {
	var input User
	var user User

	c.Bind(&input)

	// Email se user dhundho
	result := db.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Email nahi mila!"})
		return
	}

	// Password check karo
	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": "Password galat hai!"})
		return
	}

	// Token banao
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(secretKey)

	c.JSON(200, gin.H{
		"message": "Login ho gaye!",
		"token":   tokenString,
	})

}

func GetProfile(c *gin.Context) {
	email := c.MustGet("email").(string)

	c.JSON(200, gin.H{
		"message": "Tumhara Profile!",
		"email":   email,
	})
}
