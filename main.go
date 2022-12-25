package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wibowo-id/sms-backend/app/http/controllers/LoginController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/ProductController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/UserController"
	"github.com/wibowo-id/sms-backend/app/http/middlewares"
	"github.com/wibowo-id/sms-backend/app/models"
	"net/http"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Selamat datang di sms-core")
	})

	// Auth
	public.POST("/login", LoginController.Login)
	public.POST("/register", LoginController.Register)
	public.POST("/logout", LoginController.Logout)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/profile", UserController.CurrentUser)

	// Product
	protected.GET("/products", ProductController.Index)
	protected.GET("/products/:id", ProductController.Show)
	protected.POST("/products", ProductController.Create)
	protected.PUT("/products/:id", ProductController.Update)
	protected.DELETE("/products", ProductController.Delete)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
