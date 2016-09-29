package mongodb

import mgo "gopkg.in/mgo.v2"

const (
	dialUrl  = "localhost"
	Database = "userservices"
	Users    = "users"
)

func GetSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://" + dialUrl)

	// Check if connection existed
	if err != nil {
		panic(err)
	}

	return s
}
