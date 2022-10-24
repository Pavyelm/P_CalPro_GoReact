package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entry struct {
	ID       primitive.ObjectID `bson:"id"`
	Quantity *int               `json:"quantity"`
	Dish     *string            `json:"dish"`
	Calories *int               `json:"calories"`
	Protein  *int               `json:"protein"`
}
