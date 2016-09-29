package repositories

import (
	dtos "userservices/DTOs"
	"userservices/domain/models"
	"userservices/infrastructure/mongodb"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	session *mgo.Session
}

func NewUserRepository() UserRepository {
	return UserRepository{
		session: mongodb.GetSession(),
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
	err := u.session.DB(mongodb.Database).C(mongodb.Users).FindId(id).One(&matchedUser)
	if err != nil {
		panic(err)
	}

	return mapToDTO(matchedUser)
}

func (u *UserRepository) GetUserByName(name string) []dtos.UserDTO {
	var matchedUsers []models.User

	err := u.session.DB(mongodb.Database).C(mongodb.Users).Find(bson.M{"firstName": name}).All(&matchedUsers)
	if err != nil {
		panic(err)
	}

	var rs []dtos.UserDTO
	for _, user := range matchedUsers {
		rs = append(rs, mapToDTO(user))
	}

	return rs
}

func mapToDTO(user models.User) dtos.UserDTO {
	userDto := dtos.UserDTO{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		FBAccessToken: user.FBAccessToken,
	}
	userDto.GetFullName()

	return userDto
}
