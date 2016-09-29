package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id             bson.ObjectId `json:"id" bson:"_id"`
	Gender         string        `json:"gender" bson:"gender"`
	FirstName      string        `json:"firstName" bson:"firstName"`
	LastName       string        `json:"lastName" bson:"lastName"`
	DOB            time.Time     `json:"dob" bson:"dob"`
	Email          string        `json:"email" bson:"email"`
	PhoneNumber    string        `json:"phoneNumber" bson:"phoneNumber"`
	HashedPassword string        `json:"hashedPassword" bson:"hashedPassword"`
	// Facebook section
	FBAccessToken string `json:"fbAccessToken" bson:"fbAccessToken"`
	FBId          string `json:"fbId" bson:"fbId"`
	FBFriendIds   []string
}
