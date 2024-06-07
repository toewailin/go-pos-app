package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toewailin/pos-app/middleware"
)

func Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// This is just a mock. In real applications, validate the credentials with the database
	if creds.Username == "admin" && creds.Password == "password" {
		token, err := middleware.GenerateJWT(creds.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
