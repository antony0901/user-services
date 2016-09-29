package models

import "gopkg.in/mgo.v2/bson"

type AddressBook struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	UserId    bson.ObjectId `json:"id" bson:"userId"`
	Company   string        `json:"company" bson:"company"`
	Street    string        `json:"street" bson:"street"`
	Suburb    string        `json:"suburb" bson:"suburb"`
	Postcode  string        `json:"postcode" bson:"postcode"`
	City      string        `json:"city" bson:"city"`
	State     string        `json:"state" bson:"state"`
	CountryId bson.ObjectId `json:"countryId" bson:"countryId"`
	ZoneId    bson.ObjectId `json:"zoneId" bson:"zoneId"`
}
