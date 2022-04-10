package cmd

import (
	_ "github.com/berkayersoyy/go-products-example-ddd/docs"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/application"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/infrastructure/mysql"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/infrastructure/redis"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/presentation/http"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/presentation/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func setup(db *gorm.DB) *gin.Engine {

	productRepository := mysql.ProvideProductRepository(db)
	productService := application.ProvideProductService(productRepository)
	productApi := http.ProvideProductAPI(productService)

	userRepository := mysql.ProvideUserRepository(db)
	userService := application.ProvideUserService(userRepository)
	userApi := http.ProvideUserAPI(userService)

	r := redis.ProvideRedisClient()
	authService := application.ProvideAuthService(r.GetClient())
	authApi := http.ProvideAuthAPI(authService, userService)

	router := gin.Default()

	//TODO Middleware for validation
	//router.Use(validators.ProductValidator())

	//TODO Error handler can be add as a middleware

	//products
	products := router.Group("/v1")

	products.Use(middleware.AuthorizeJWTMiddleware(authService))

	products.GET("/products", productApi.GetAllProducts)
	products.POST("/products", productApi.AddProduct)
	products.GET("/products/:id", productApi.GetProductByID)
	products.DELETE("/products/:id", productApi.DeleteProduct)
	products.PUT("/products/:id", productApi.UpdateProduct)

	//users
	users := router.Group("/v1")
	users.GET("/users", userApi.GetAllUsers)
	users.POST("/users", userApi.AddUser)
	users.GET("/users/:id", userApi.GetUserByID)
	users.DELETE("/users/:id", userApi.DeleteUser)
	users.PUT("/users/:id", userApi.UpdateUser)

	//auth
	auth := router.Group("/v1")
	auth.POST("/login", authApi.Login)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	dbClient := mysql.ProvideMysqlClient("./")
	db := dbClient.GetClient()
	defer db.Close()
	r := setup(db)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
