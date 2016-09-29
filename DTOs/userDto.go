package dtos

type UserDTO struct {
	FirstName     string `json:"firstName" bson:"firstName"`
	LastName      string `json:"lastName" bson:"lastName"`
	FullName      string `json:"fullName" bson:"fullName"`
	FBAccessToken string `json:"fbAccessToken" bson:"fbAccessToken"`
}

func (u *UserDTO) GetFullName() {
	u.FullName = u.FirstName + u.LastName
}
