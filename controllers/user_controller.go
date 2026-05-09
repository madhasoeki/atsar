package controllers

import (
	"net/http"
	"regexp"

	"atsar/config"
	"atsar/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	NamaLengkap   string `json:"nama_lengkap" binding:"required"`
	NamaPanggilan string `json:"nama_panggilan" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	RoleID        uint   `json:"role_id" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Password validation (regex): lowercase, uppercase, and number
	passwordRegex := regexp.MustCompile(`[a-z]`)
	if !passwordRegex.MatchString(input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format password tidak valid (harus mengandung huruf kecil)"})
		return
	}
	passwordRegex = regexp.MustCompile(`[A-Z]`)
	if !passwordRegex.MatchString(input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format password tidak valid (harus mengandung huruf besar)"})
		return
	}
	passwordRegex = regexp.MustCompile(`[0-9]`)
	if !passwordRegex.MatchString(input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format password tidak valid (harus mengandung angka)"})
		return
	}

	// Check if email already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email sudah terdaftar"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}

	// Create user
	user := models.User{
		NamaLengkap:   input.NamaLengkap,
		NamaPanggilan: input.NamaPanggilan,
		Email:         input.Email,
		Password:      string(hashedPassword),
		RoleID:        input.RoleID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan user ke database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User berhasil ditambahkan"})
}
