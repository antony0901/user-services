package repositories

import (
	dtos "userservices/DTOs"
	"userservices/domain/models"

	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	BaseRepository
}

// NewUserRepository return instant of UserRepository
func NewUserRepository() UserRepository {
	bs := NewSession("users")
	return UserRepository{
		BaseRepository: bs,
	}
}

func (u *UserRepository) CreateUserWithFB(userDto dtos.UserDTO) bool {
	newUser := models.User{}
	newUser.MapFromFBUser(userDto)
	newUser.Id = bson.NewObjectId()

	return u.create(newUser)
}

func (u *UserRepository) GetUserById(id string) dtos.UserDTO {
	matchedUser := models.User{}
	u.GetById(id).One(&matchedUser)

	return mapToDTO(matchedUser)
}

func (u *UserRepository) GetUserByName(name string) []dtos.UserDTO {
	var matchedUsers []models.User
	query := bson.M{
		"firstName": name,
	}
	u.find(query).All(&matchedUsers)

	var rs []dtos.UserDTO
	for _, user := range matchedUsers {
		rs = append(rs, mapToDTO(user))
	}

	return rs
}

func (u *UserRepository) GetBy(queries ...interface{}) []dtos.UserDTO {
	var matchedUsers []models.User
	u.find(queries).All(&matchedUsers)

	var rs []dtos.UserDTO
	for _, user := range matchedUsers {
		rs = append(rs, mapToDTO(user))
	}

	return rs
}

func (u *UserRepository) GetOneBy(selector string, searchBy string) models.User {
	user := models.User{}
	u.getOneBy(selector, searchBy).One(&user)

	return user
}

// mapToDTO returns object as DTO or payload object to API.
func mapToDTO(user models.User) dtos.UserDTO {
	userDto := dtos.UserDTO{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		FBAccessToken: user.FBAccessToken,
	}

	return userDto
}
