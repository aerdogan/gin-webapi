package model

import (
	"gin-webapi/database"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User structure
type User struct {
	ID        bson.ObjectId `bson:"_id"`
	FirstName string        `bson:"firstname"`
	LastName  string        `bson:"lastname"`
	Address   string        `bson:"address"`
	Age       int           `bson:"age"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

// Users list
type Users []User

// UserInfo model function
func UserInfo(id bson.ObjectId, userCollection string) (User, error) {
	db := database.GetMongoDB()
	user := User{}
	err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)
	return user, err
}
