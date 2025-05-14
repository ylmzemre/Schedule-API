package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ylmz/ders-api/models"
	"github.com/ylmzemre/ders-api/config"
	"github.com/ylmzemre/ders-api/handlers"
	"os"
)

func main() {
	// 1. DB Bağlantısı
	config.ConnectDatabase()

	// 2. Migrasyon
	config.DB.AutoMigrate(&models.Ders{})

	// 3. Router
	r := gin.Default()

	// 4. Route’lar
	ders := r.Group("/ders")
	{
		ders.POST("", handlers.CreateDers)
		ders.GET("", handlers.GetDersler)
		ders.GET("/:id", handlers.GetDers)
		ders.PUT("/:id", handlers.UpdateDers)
		ders.DELETE("/:id", handlers.DeleteDers)
	}

	// 5. Çalıştır
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
