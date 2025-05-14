package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ylmzemre/ders-api/config"
	"github.com/ylmzemre/ders-api/models"
	"net/http"
)

func CreateDers(c *gin.Context) {
	var ders models.Ders
	if err := c.ShouldBindJSON(&ders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&ders)
	c.JSON(http.StatusCreated, ders)
}

func GetDersler(c *gin.Context) {
	var dersler []models.Ders
	config.DB.Find(&dersler)
	c.JSON(http.StatusOK, dersler)
}

func GetDers(c *gin.Context) {
	id := c.Param("id")
	var ders models.Ders
	if err := config.DB.First(&ders, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ders bulunamadı"})
		return
	}
	c.JSON(http.StatusOK, ders)
}

func UpdateDers(c *gin.Context) {
	id := c.Param("id")
	var ders models.Ders
	if err := config.DB.First(&ders, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ders bulunamadı"})
		return
	}
	var input models.Ders
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&ders).Updates(input)
	c.JSON(http.StatusOK, ders)
}

func DeleteDers(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Ders{}, id)
	c.JSON(http.StatusNoContent, nil)
}
