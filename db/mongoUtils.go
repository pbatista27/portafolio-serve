package db

import (
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {

	if session == nil {
		var err error

		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{os.Getenv("SERVER_DB")},
			Username: "",
			Password: "",
			Database: os.Getenv("DB_NAME"),
			Timeout:  60 * time.Second,
		})

		if err != nil {
			log.Fatalf("[Getdbsession] %s", err)
		}

	} else {
		createDbSession()
	}
	return session

}

func createDbSession() {

	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{os.Getenv("SERVER_DB")},
		Username: "",
		Password: "",
		Timeout:  60 * time.Second,
		Database: os.Getenv("DB_NAME"),
	})

	if err != nil {
		log.Fatalf("[createdbSession] %s", err)
	}

}

//add indexes into mongoDB
// func addIndex() {

// 	var err error

// 	userIndex := mgo.Index{
// 		Key:        []string{"id"},
// 		Unique:     true,
// 		Background: true,
// 		Sparse:     true,
// 	}

// 	projectIndex := mgo.Index{
// 		Key:        []string{"userid"},
// 		Unique:     false,
// 		Background: true,
// 		Sparse:     true,
// 	}

// 	session = GetSession().Copy()
// 	defer session.Close()

// 	userCol := session.DB(os.Getenv("DB_NAME")).C("users")
// 	projectCol := session.DB(os.Getenv("DB_NAME")).C("projects")

// 	err = userCol.EnsureIndex(userIndex)

// 	if err != nil {
// 		log.Fatalf("[addIndex] %s", err)
// 	}

// 	err = projectCol.EnsureIndex(projectIndex)

// 	if err != nil {
// 		log.Fatalf("[addIndex] %s", err)
// 	}

// }
