package routes

import (
	category "gin-webapi/controllers/category"
	product "gin-webapi/controllers/product"
	user "gin-webapi/controllers/user"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func StartGin() {

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	api := router.Group("/api")
	{
		api.GET("/users", user.GetAllUser)
		api.POST("/users", user.CreateUser)
		api.GET("/users/:id", user.GetUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.DELETE("/users/:id", user.DeleteUser)

		api.GET("/products", product.GetAllProduct)
		api.POST("/products", product.CreateProduct)
		api.GET("/products/:id", product.GetProduct)
		api.PUT("/products/:id", product.UpdateProduct)
		api.DELETE("/products/:id", product.DeleteProduct)

		api.GET("/categories", category.GetAllCategory)
		api.POST("/categories", category.CreateCategory)
		api.GET("/categories/:id", category.GetCategory)
		api.PUT("/categories/:id", category.UpdateCategory)
		api.DELETE("/categories/:id", category.DeleteCategory)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":8000")
}
