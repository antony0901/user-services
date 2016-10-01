package models

import (
	"strings"
	"time"
	dtos "userservices/DTOs"

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

func NewUser(email string, phoneNumber string) User {
	return User{
		Id:          bson.NewObjectId(),
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

func (u *User) SetName(fullName string) {
	s := strings.Split(fullName, " ")
	u.FirstName = s[0]
	u.LastName = s[1]
}

func MapToUser(dto dtos.UserDTO, user *User) {
	user.SetName(dto.FullName)
	user.DOB = dto.Dob
	user.FBAccessToken = dto.FBAccessToken
	user.FBId = dto.FBId
}
