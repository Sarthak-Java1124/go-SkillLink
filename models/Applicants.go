package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Applicants struct {
	ID             bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicantArray []Bids        `json:"applicants,omitempty" bson:"applicants,omitempty"`
}
