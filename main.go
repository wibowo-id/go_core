package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/wibowo-id/sms-backend/app/http/controllers/LoginController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/ProductController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/UserController"
	"github.com/wibowo-id/sms-backend/app/http/middlewares"
	initModel "github.com/wibowo-id/sms-backend/app/models"
	cfg "github.com/wibowo-id/sms-backend/config"
)

type App struct {
	config cfg.Config
}

var app App

func init() {
	config := cfg.NewViperConfig()
	app = App{config: config}

	if config.GetBool(`app.debug`) {
		fmt.Println("Service RUN on DEBUG mode. Port: " + config.GetString("app.host"))
	}
}

func main() {
	initModel.ConnectDatabase()

	r := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("APP_SECRET")))
	r.Use(sessions.Sessions("coresession", store))

	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}
	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Selamat datang di sms-core")
	})

	// Auth
	public.POST("/login", LoginController.Login)
	public.POST("/register", LoginController.Register)
	public.POST("/logout", LoginController.Logout)

	// as user with jwt auth middleware
	userLogged := r.Group("/api/user")
	userLogged.Use(middlewares.JwtAuthMiddleware())
	userLogged.GET("/profile", UserController.Profile)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/profile", UserController.CurrentUser)

	// Product
	protected.GET("/products", ProductController.Index)
	protected.GET("/products/:id", ProductController.Show)
	protected.POST("/products", ProductController.Create)
	protected.PUT("/products/:id", ProductController.Update)
	protected.DELETE("/products", ProductController.Delete)

	err = r.Run(app.config.GetString("app.host"))
	if err != nil {
		return
	}
}
