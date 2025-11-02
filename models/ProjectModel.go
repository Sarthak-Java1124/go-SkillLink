package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectModel struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       *string            `json:"title,omitempty" bson:"title,omitempty"`
	Description *string            `json:"description ,omitempty" bson:"description,omitempty"`
	Budget_Min  *int               `json:"budget_min,omitempty" bson:"budget_min,omitempty"`
	Budget_Max  *int               `json:"budget_max,omitempty" bson:"budget_max,omitempty"`
	Client_Id   *int               `json:"client_id ,omitempty" bson:"client_id,omitempty"`
	Status      *string            `json:"status,omitempty" bson:"status,omitempty"`
	Tags        *[]string          `json:"balance ,omitempty" bson:"balance,omitempty"`
}
