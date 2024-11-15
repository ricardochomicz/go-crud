package entity

type UserEntity struct {
	Id       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Age      int8   `bson:"age"`
}
