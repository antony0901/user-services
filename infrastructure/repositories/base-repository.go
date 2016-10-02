package repositories

import (
	"fmt"
	"userservices/infrastructure/common"
	"userservices/infrastructure/mongodb"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
func (b *BaseRepository) find(query interface{}) *mgo.Query {
	fmt.Println("---- collection -----")
	fmt.Println(b.name)
	return b.session.DB(mongodb.Database).C(b.name).Find(query)
}

func (b *BaseRepository) Update(selector interface{}, update interface{}) {
	b.session.DB(mongodb.Database).C(b.name).Update(selector, update)
}

func (b *BaseRepository) UpdateById(id bson.ObjectId, update interface{}) {
	b.session.DB(mongodb.Database).C(b.name).UpdateId(id, update)
}

// As a private method that used in inheritant objects.
func (b *BaseRepository) getOneBy(selector string, searchBy string) *mgo.Query {
	query := bson.M{
		selector: searchBy,
	}

	return b.find(query)
}

// As a private method that used in inheritant objects.
func (b *BaseRepository) create(entity interface{}) bool {
	err := b.session.DB(mongodb.Database).C(b.name).Insert(entity)
	common.Check(err)

	return true
}
