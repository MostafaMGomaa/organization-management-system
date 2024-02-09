package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Members struct {
	ID primitive.ObjectID `bson "_id,omitempty" json:"id"`
	Name string  `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Password string`bson:"password" json:"password"`
	AccessLevel string`bson:"access_level" json:"access_level"`
}

type Organiztion struct {
	ID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string  `bson:"name" json:"name"`
    Description string `bson:"description" json:"description"`
    Members []Members `bson:"members" json:"members"`
}