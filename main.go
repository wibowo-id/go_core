package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wibowo-id/sms-backend/app/http/controllers/BalanceController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/BankController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/LoginController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/PaymentCallbackController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/PaymentController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/PermissionController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/PesantrenController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/ProductController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/RoleController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/RoleHasPermissionController"
	"github.com/wibowo-id/sms-backend/app/http/controllers/TransactionController"
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
	public.POST("payments/midtrans-notification", PaymentCallbackController.Receive)

	// Auth
	public.POST("/login", LoginController.Login)
	public.POST("/register", LoginController.Register)
	public.POST("/logout", LoginController.Logout)

	// as user role_id = 2
	userLogged := r.Group("/api/user")
	userLogged.Use(middlewares.JwtAuthMiddleware())
	userLogged.GET("/daftar-pembayaran", PaymentController.List)
	userLogged.GET("/cart", PaymentController.Cart)
	userLogged.POST("/cart", PaymentController.AddToCart)
	userLogged.POST("/update-cart", PaymentController.DeleteFromCart)
	userLogged.GET("/checkout", PaymentController.Checkout)
	userLogged.POST("/checkout", PaymentController.Store)

	userLogged.GET("/transaction", TransactionController.Index)
	userLogged.GET("/transaction/{orderId}", TransactionController.Detail)
	userLogged.POST("/transaction/payment", TransactionController.Payment)

	userLogged.GET("/profile", UserController.Profile)
	userLogged.POST("/profile", UserController.SaveProfile)

	// as user role_id = 1
	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/profile", UserController.CurrentUser)

	// Product
	protected.GET("/products", ProductController.Index)
	protected.GET("/products/:id", ProductController.Show)
	protected.POST("/products", ProductController.Create)
	protected.PUT("/products/:id", ProductController.Update)
	protected.DELETE("/products", ProductController.Delete)

	core := protected.Group("/core")
	product := core.Group("/produk")
	product.GET("/", ProductController.Index)
	product.POST("/create", ProductController.Store)
	product.GET("/update/{id}", ProductController.Edit)
	product.PUT("/update/{id}", ProductController.Update)
	product.DELETE("/delete/{id}", ProductController.Delete)

	role := core.Group("/role")
	role.GET("/", RoleController.Index)
	role.POST("/create", RoleController.Store)
	role.GET("/update/{id}", RoleController.Edit)
	role.PUT("/update/{id}", RoleController.Update)
	role.DELETE("/delete/{id}", RoleController.Destroy)

	permission := core.Group("/permissions")
	permission.GET("/index", PermissionController.Index)
	permission.POST("/create", PermissionController.Store)
	permission.GET("/update/{id}", PermissionController.Edit)
	permission.PUT("/update/{id}", PermissionController.Update)
	permission.DELETE("/update/{id}", PermissionController.Delete)

	roleHasPermission := core.Group("/role_has_permissions")
	roleHasPermission.GET("/index", RoleHasPermissionController.Index)
	roleHasPermission.POST("/create", RoleHasPermissionController.Store)
	roleHasPermission.GET("/update/{id}", RoleHasPermissionController.Edit)
	roleHasPermission.PUT("/update/{id}", RoleHasPermissionController.Update)
	roleHasPermission.DELETE("/update/{id}", RoleHasPermissionController.Delete)

	bank := core.Group("/bank")
	bank.GET("/index", BankController.Index)
	bank.POST("/create", BankController.Store)
	bank.GET("/update/{id}", BankController.Edit)
	bank.PUT("/update/{id}", BankController.Update)
	bank.DELETE("/update/{id}", BankController.Delete)

	pesantren := core.Group("/pesantren")
	pesantren.GET("/index", PesantrenController.Index)
	pesantren.POST("/create", PesantrenController.Store)
	pesantren.GET("/update/{id}", PesantrenController.Edit)
	pesantren.PUT("/update/{id}", PesantrenController.Update)
	pesantren.DELETE("/update/{id}", PesantrenController.Delete)

	balance := core.Group("/balance")
	balance.GET("/index", BalanceController.Index)
	balance.POST("/create", BalanceController.Store)
	balance.GET("/update/{id}", BalanceController.Edit)
	balance.PUT("/update/{id}", BalanceController.Update)
	balance.DELETE("/update/{id}", BalanceController.Delete)

	payment := core.Group("/payment")
	payment.GET("/", PaymentController.Index)
	payment.POST("/create", PaymentController.Store)
	payment.GET("/update/{id}", PaymentController.Edit)
	payment.PUT("/update/{id}", PaymentController.Update)
	payment.DELETE("/update/{id}", PaymentController.Delete)

	master := core.Group("/master")
	user := master.Group("/user")
	user.GET("/", UserController.Index)
	user.GET("/detail/{userId}/{type}", UserController.Show)
	user.POST("/create", UserController.Store)
	user.GET("/update/{userId}", UserController.Edit)
	user.PUT("/update/{userId}", UserController.Update)
	user.DELETE("/delete/{userId}", UserController.Delete)
	user.POST("/update-status", UserController.UpdateStatus)
	user.POST("/import", UserController.Import)
	user.GET("/export-csv", UserController.ExportCsv)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
