package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toewailin/pos-app/models"
	"gorm.io/gorm"
)

func GetItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var items []models.Item
		db.Find(&items)
		c.JSON(http.StatusOK, items)
	}
}

func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newItem models.Item
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newItem)
		c.JSON(http.StatusCreated, newItem)
	}
}
