package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Bids struct {
	ID        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Pitch     string        `json:"pitch,omitempty" bson:"pitch,omitempty"`
	BidAmount int           `json:"bid_amount,omitempty" bson:"bid_amount,omitempty"`
}
