package product

import (
	"errors"
	"gin-webapi/database"
	"net/http"
	"time"

	product "gin-webapi/models/product"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

const ProductCollection = "product"

var (
	errNotExist        = errors.New("Products are not exist")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidBody     = errors.New("Invalid request body")
	errInsertionFailed = errors.New("Error in the product insertion")
	errUpdationFailed  = errors.New("Error in the product updation")
	errDeletionFailed  = errors.New("Error in the product deletion")
)

// GetAllProduct Endpoint
func GetAllProduct(c *gin.Context) {
	db := database.GetMongoDB()
	products := product.Products{}
	err := db.C(ProductCollection).Find(bson.M{}).All(&products)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "products": &products})
}

// GetProduct Endpoint
func GetProduct(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	product, err := product.ProductInfo(id, ProductCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "product": &product})
}

// CreateProduct Endpoint
func CreateProduct(c *gin.Context) {
	db := database.GetMongoDB()
	product := product.Product{}
	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	product.ID = bson.NewObjectId()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	err = db.C(ProductCollection).Insert(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "product": &product})
}

// UpdateProduct Endpoint
func UpdateProduct(c *gin.Context) {
	db := database.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	existingProduct, err := product.ProductInfo(id, ProductCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}

	err = c.Bind(&existingProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	existingProduct.ID = id
	existingProduct.UpdatedAt = time.Now()
	err = db.C(ProductCollection).Update(bson.M{"_id": &id}, existingProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdationFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "product": &existingProduct})
}

// DeleteProduct Endpoint
func DeleteProduct(c *gin.Context) {
	db := database.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	err := db.C(ProductCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Product deleted successfully"})
}
