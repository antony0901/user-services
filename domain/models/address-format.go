package models

import "gopkg.in/mgo.v2/bson"

type AddressFormat struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Format  string        `json:"format" bson:"format"`
	Summary string        `json:"summary" bson:"summary"`
}
