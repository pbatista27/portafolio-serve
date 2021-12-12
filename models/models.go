package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Email       string        `bson:"email,omitempty" json:"email" `
	Password    string        `bson:"password" json:"passoword"`
	Name        string        `bson:"name" json:"name"`
	LastName    string        `bson:"last_name" json:"last_name"`
	Description string        `bson:"description" json:"description"`
	Img         string        `bson:"img" json:"img"`
}

// type Profile struct {
// 	Id          bson.ObjectId `bson:"_id" json:"id"`
// 	Name        string        `bson:"name" json:"name"`
// 	LastName    string        `bson:"last_name" json:"last_name"`
// 	Description string        `bson:"description" json:"description"`
// 	Img         string        `bson:"img" json:"img"`
// }

type Study struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title string        `bson:"title" json:"title"`
	Year  time.Time     `bson:"year" json:"year"`
}

type Project struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Imgs        []string      `bson:"imgs" json:"imgs"`
}

type SocialNetworks struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Logo  string        `bson:"logo" json:"logo"`
	Url   string        `bson:"url" json:"url"`
	Name  string        `bson:"name" json:"name"`
	Statu bool          `bson:"statu" json:"statu"`
}

type Languaje struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Logo  string        `bson:"logo" json:"logo"`
	Statu bool          `bson:"statu" json:"statu"`
}

type Login struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}
