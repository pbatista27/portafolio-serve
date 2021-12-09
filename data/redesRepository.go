package data

import (
	"github.com/pbatista27/portafolio-serve/db"
	"github.com/pbatista27/portafolio-serve/models"
	"gopkg.in/mgo.v2/bson"
)

type RedesRepository struct{}

func (rp RedesRepository) GetAll() ([]models.SocialNetworks, error) {

	var redes []models.SocialNetworks

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("SocialNetworks")

	err := col.Find(bson.M{"statu": true}).All(&redes)

	if err != nil {
		return redes, err
	}

	return redes, nil
}

func (rp RedesRepository) Show(id string) (models.SocialNetworks, error) {

	var redes models.SocialNetworks
	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("SocialNetworks")

	err := col.FindId(bson.ObjectIdHex(id)).One(&redes)

	if err != nil {
		return redes, err
	}

	return redes, nil

}

func (rp RedesRepository) Create(redes models.SocialNetworks) error {

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("SocialNetworks")
	err := col.Insert(redes)

	if err != nil {
		return err
	}

	return nil
}

func (rp RedesRepository) Update(redes models.SocialNetworks) error {

	var red models.SocialNetworks

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("SocialNetworks")
	err := col.FindId(redes.Id).One(&red)

	if err != nil {
		return err
	}

	red.Name = redes.Name
	red.Url = redes.Url
	red.Logo = redes.Logo
	red.Statu = redes.Statu

	err = col.UpdateId(red.Id, red)

	if err != nil {
		return err
	}

	return nil
}

func (rp RedesRepository) Delete(id string) error {

	var redes models.SocialNetworks

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("SocialNetworks")
	err := col.FindId(bson.ObjectIdHex(id)).One(&redes)

	if err != nil {
		return err
	}

	err = col.RemoveId(redes.Id)

	if err != nil {
		return err
	}

	return nil

}
