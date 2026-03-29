package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Secret Key
var secretKey = []byte(os.Getenv("JWT_SECRET"))

// auth.go — Register function update:

func Register(c *gin.Context) {
    var user User

    // ❌ Data sahi nahi aaya?
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{
            "error": "Data sahi nahi hai!",
        })
        return
    }

    // ❌ Fields khaali hain?
    if user.Name == "" || user.Email == "" || user.Password == "" {
        c.JSON(400, gin.H{
            "error": "Name, Email aur Password zaroori hain!",
        })
        return
    }

    // Password hash karo
    hashedPassword, err := bcrypt.GenerateFromPassword(
        []byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(500, gin.H{
            "error": "Password process nahi hua!",
        })
        return
    }
    user.Password = string(hashedPassword)

    // Database mein save karo
    if err := db.Create(&user).Error; err != nil {
        c.JSON(400, gin.H{
            "error": "Email already registered hai!",
        })
        return
    }

    c.JSON(201, gin.H{
        "message": "Register ho gaye!",
        "name":    user.Name,
        "email":   user.Email,
    })
}

// Login Handler
func Login(c *gin.Context) {
    var input User
    var user User

    // Data check karo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{
            "error": "Data sahi nahi hai!",
        })
        return
    }

    // Email se user dhundho
    if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(404, gin.H{
            "error": "Email registered nahi hai!",
        })
        return
    }

    // Password check karo
    if err := bcrypt.CompareHashAndPassword(
        []byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(401, gin.H{
            "error": "Password galat hai!",
        })
        return
    }

    // Token banao
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":    user.ID,
        "email": user.Email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        c.JSON(500, gin.H{
            "error": "Token nahi ban paya!",
        })
        return
    }

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
