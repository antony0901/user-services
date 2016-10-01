package dtos

import "time"

type UserDTO struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	FullName      string    `json:"fullName"`
	FBAccessToken string    `json:"fbAccessToken"`
	FBId          string    `json:"fbId"`
	Dob           time.Time `json:"dob"`
}

func (u *UserDTO) GetFullName() {
	u.FullName = u.FirstName + u.LastName
}
