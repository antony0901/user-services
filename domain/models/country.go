package models

import "gopkg.in/mgo.v2/bson"

type Country struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Name            string        `json:"namae" bson:"name"`
	ISOCode2        string        `json:"isoCode2" bson:"isoCode2"`
	ISOCode3        string        `json:"isoCode3" bson:"isoCode3"`
	AddressFormatId bson.ObjectId `json:"addressFormatId" bson:"addressFormatId"`
}
