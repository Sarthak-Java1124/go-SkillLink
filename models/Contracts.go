package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Contracts struct {
	ID            bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Project_Id    bson.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty"`
	Client_Id     bson.ObjectID `json:"client_id ,omitempty" bson:"client_id,omitempty"`
	Freelancer_Id bson.ObjectID `json:"freelancer_id,omitempty" bson:"freelancer_id,omitempty"`
	Agreed_Budget *int          `json:"agreed_budget,omitempty" bson:"agreed_budget,omitempty"`
	Status        *string       `json:"status,omitempty" bson:"status,omitempty"`
	Deadline      bson.DateTime `json:"deadline ,omitempty" bson:"deadline,omitempty"`
}
