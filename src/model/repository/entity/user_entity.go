package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Age      int8               `bson:"age"`
}
