package model

import (
	"gin-webapi/database"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Category structure
type Category struct {
	ID           bson.ObjectId `bson:"_id"`
	CategoryName string        `bson:"categoryname"`
	CreatedAt    time.Time     `bson:"created_at"`
	UpdatedAt    time.Time     `bson:"updated_at"`
}

// Category list
type Categories []Category

// CategoryInfo model function
func CategoryInfo(id bson.ObjectId, categoryCollection string) (Category, error) {
	db := database.GetMongoDB()
	category := Category{}
	err := db.C(categoryCollection).Find(bson.M{"_id": &id}).One(&category)
	return category, err
}
