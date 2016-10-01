package repositories

import (
	"fmt"
	"userservices/infrastructure/mongodb"

	mgo "gopkg.in/mgo.v2"
)

type BaseRepository struct {
	session *mgo.Session
	// collection name
	name string
}

func NewSession(name string) BaseRepository {
	return BaseRepository{
		session: mongodb.GetSession(),
		name:    name,
	}
}

/*
	Return Mongodb Query object
	id: is a bson object id as a string.
*/
func (b *BaseRepository) GetById(id string) *mgo.Query {
	return b.session.DB(mongodb.Database).C(b.name).FindId(id)
}

/*
	Returns Mongodb Query object
	query: is a criteria to get data as many.
*/
func (b *BaseRepository) Find(query interface{}) *mgo.Query {
	fmt.Println("---- collection -----")
	fmt.Println(b.name)
	return b.session.DB(mongodb.Database).C(b.name).Find(query)
}
