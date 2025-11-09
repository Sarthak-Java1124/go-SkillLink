package models

import "go.mongodb.org/mongo-driver/v2/bson"

// Bids represents a bid made by a freelancer on a project.
// swagger:model Bid
type Bids struct {
	ID           bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Pitch        string        `json:"pitch,omitempty" bson:"pitch,omitempty"`
	BidAmount    int           `json:"bid_amount,omitempty" bson:"bid_amount,omitempty"`
	FreelancerId int           `json:"freelancer_id,omitempty" bson:"freelancer_id,omitempty"`
}
