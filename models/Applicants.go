package models

import "go.mongodb.org/mongo-driver/v2/bson"

// Applicants groups every bid submitted for a given client.
// swagger:model Applicants
type Applicants struct {
	ID             bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ClientId       int           `json:"client_id" bson:"client_id,omitempty"`
	ApplicantArray []Bids        `json:"applicants,omitempty" bson:"applicants,omitempty"`
}
