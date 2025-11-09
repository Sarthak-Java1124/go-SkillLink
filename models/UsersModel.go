package models

import "go.mongodb.org/mongo-driver/v2/bson"

// UserModel represents a registered user in the system.
// swagger:model UserModel
type UserModel struct {
	ID       bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name,omitempty" bson:"name,omitempty"`
	Email    string        `json:"email,omitempty" bson:"email,omitempty"`
	Password *string       `json:"password,omitempty" bson:"password,omitempty"`
	Role     string        `json:"role,omitempty" bson:"role,omitempty"`
	Country  string        `json:"country,omitempty" bson:"country,omitempty"`
	Rating   string        `json:"rating,omitempty" bson:"rating,omitempty"`
	Balance  int           `json:"balance,omitempty" bson:"balance,omitempty"`
}
