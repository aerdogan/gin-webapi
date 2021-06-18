package category

import (
	"errors"
	"gin-webapi/database"
	"net/http"
	"time"

	category "gin-webapi/models/category"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

const CategoryCollection = "category"

var (
	errNotExist        = errors.New("Categories are not exist")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidBody     = errors.New("Invalid request body")
	errInsertionFailed = errors.New("Error in the category insertion")
	errUpdationFailed  = errors.New("Error in the category updation")
	errDeletionFailed  = errors.New("Error in the category deletion")
)

// GetAllCategory Endpoint
func GetAllCategory(c *gin.Context) {
	db := database.GetMongoDB()
	categories := category.Categories{}
	err := db.C(CategoryCollection).Find(bson.M{}).All(&categories)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "categories": &categories})
}

// GetCategory Endpoint
func GetCategory(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
	category, err := category.CategoryInfo(id, CategoryCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "category": &category})
}

// CreateCategory Endpoint
func CreateCategory(c *gin.Context) {
	db := database.GetMongoDB()
	category := category.Category{}
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	category.ID = bson.NewObjectId()
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	err = db.C(CategoryCollection).Insert(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "category": &category})
}

// UpdateCategory Endpoint
func UpdateCategory(c *gin.Context) {
	db := database.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
	existingCategory, err := category.CategoryInfo(id, CategoryCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}

	err = c.Bind(&existingCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	existingCategory.ID = id
	existingCategory.UpdatedAt = time.Now()
	err = db.C(CategoryCollection).Update(bson.M{"_id": &id}, existingCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdationFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "category": &existingCategory})
}

// DeleteCategory Endpoint
func DeleteCategory(c *gin.Context) {
	db := database.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	err := db.C(CategoryCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Category deleted successfully"})
}
