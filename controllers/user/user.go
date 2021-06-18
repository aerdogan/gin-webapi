package user

import (
	"errors"
	"gin-webapi/database"
	"net/http"
	"time"

	user "gin-webapi/models/user"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

const UserCollection = "user"

var (
	errNotExist        = errors.New("Users are not exist")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidBody     = errors.New("Invalid request body")
	errInsertionFailed = errors.New("Error in the user insertion")
	errUpdationFailed  = errors.New("Error in the user updation")
	errDeletionFailed  = errors.New("Error in the user deletion")
)

// GetAllUser Endpoint
func GetAllUser(c *gin.Context) {
	db := database.GetMongoDB()
	users := user.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "users": &users})
}

// GetUser Endpoint
func GetUser(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
	user, err := user.UserInfo(id, UserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

// CreateUser Endpoint
func CreateUser(c *gin.Context) {
	db := database.GetMongoDB()
	user := user.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = db.C(UserCollection).Insert(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

// UpdateUser Endpoint
func UpdateUser(c *gin.Context) {
	db := database.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	existingUser, err := user.UserInfo(id, UserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	err = c.Bind(&existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	existingUser.ID = id
	existingUser.UpdatedAt = time.Now()
	err = db.C(UserCollection).Update(bson.M{"_id": &id}, existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdationFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &existingUser})
}

// DeleteUser Endpoint
func DeleteUser(c *gin.Context) {
	db := database.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	err := db.C(UserCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted successfully"})
}
