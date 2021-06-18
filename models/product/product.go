package model

import (
	"gin-webapi/database"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product structure
type Product struct {
	ID              bson.ObjectId `bson:"_id"`
	CategoryId      int           `bson:"categoryid"`
	ProductName     string        `bson:"productname"`
	UnitPrice       float64       `bson:"unitprice"`
	UnitsInStock    int           `bson:"unitsinstock"`
	QuantityPerUnit string        `bson:"quantityperunit"`
	CreatedAt       time.Time     `bson:"created_at"`
	UpdatedAt       time.Time     `bson:"updated_at"`
}

// Products list
type Products []Product

// ProductInfo model function
func ProductInfo(id bson.ObjectId, productCollection string) (Product, error) {
	db := database.GetMongoDB()
	product := Product{}
	err := db.C(productCollection).Find(bson.M{"_id": &id}).One(&product)
	return product, err
}
