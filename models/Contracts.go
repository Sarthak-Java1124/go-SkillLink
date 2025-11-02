package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contracts struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Project_Id    primitive.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty"`
	Client_Id     primitive.ObjectID `json:"client_id ,omitempty" bson:"client_id,omitempty"`
	Freelancer_Id primitive.ObjectID `json:"freelancer_id,omitempty" bson:"freelancer_id,omitempty"`
	Agreed_Budget *int               `json:"agreed_budget,omitempty" bson:"agreed_budget,omitempty"`
	Status        *string            `json:"status,omitempty" bson:"status,omitempty"`
	Deadline      primitive.DateTime `json:"deadline ,omitempty" bson:"deadline,omitempty"`
}
