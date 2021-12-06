package db

import (
	"os"

	"gopkg.in/mgo.v2"
)

type DbSession struct {
	MongoSession *mgo.Session
}

func (db *DbSession) Close() {
	db.MongoSession.Close()
}

func (db *DbSession) DbCollection(name string) *mgo.Collection {

	return db.MongoSession.DB(os.Getenv("DB_NAME")).C(name)

}

func NewDbSession() *DbSession {

	session := GetSession().Copy()
	dbsession := &DbSession{
		MongoSession: session,
	}

	return dbsession

}
