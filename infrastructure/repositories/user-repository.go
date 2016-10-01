package repositories

import (
	dtos "userservices/DTOs"
	"userservices/domain/models"
	"userservices/infrastructure/mongodb"

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

func (u *UserRepository) CreateUser(userDto dtos.UserDTO) {
	newUser := models.User{
		Id:            bson.NewObjectId(),
		FirstName:     userDto.FirstName,
		LastName:      userDto.LastName,
		FBAccessToken: userDto.FBAccessToken,
	}

	u.session.DB(mongodb.Database).C(mongodb.Users).Insert(newUser)
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
	// u.GetMany(query, matchedUsers)
	u.Find(query).All(&matchedUsers)

	var rs []dtos.UserDTO
	for _, user := range matchedUsers {
		rs = append(rs, mapToDTO(user))
	}

	return rs
}

// mapToDTO returns object as DTO or payload object to API.
func mapToDTO(user models.User) dtos.UserDTO {
	userDto := dtos.UserDTO{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		FBAccessToken: user.FBAccessToken,
	}
	userDto.GetFullName()

	return userDto
}
