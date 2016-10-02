package dtos

import "time"

type UserDTO struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	FullName      string    `json:"fullName"`
	FBAccessToken string    `json:"fbAccessToken"`
	FBId          string    `json:"fbId"`
	Dob           time.Time `json:"dob"`
	Email         string    `json:"email"`
	Picture       string    `json:"picture"`
}
