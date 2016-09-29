package models

import "gopkg.in/mgo.v2/bson"

type Zone struct {
	Id        bson.ObjectId
	CountryId bson.ObjectId
	Code      string
	Name      string
}
