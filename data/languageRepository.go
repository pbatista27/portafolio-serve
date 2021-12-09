package data

import (
	"github.com/pbatista27/portafolio-serve/db"
	"github.com/pbatista27/portafolio-serve/models"
	"gopkg.in/mgo.v2/bson"
)

type LanguajeRespository struct{}

func (lr LanguajeRespository) GetAll() ([]models.Languaje, error) {

	var langs []models.Languaje

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("languajes")
	err := col.Find(bson.M{"statu": true}).All(&langs)

	if err != nil {

		return nil, err

	}

	return langs, nil

}

func (lr LanguajeRespository) Create(lang models.Languaje) (models.Languaje, error) {

	con := db.NewDbSession()
	col := con.DbCollection("languajes")
	err := col.Insert(lang)
	if err != nil {
		return lang, err
	}

	return lang, nil
}

func (lr LanguajeRespository) Show(id string) (models.Languaje, error) {

	var lang models.Languaje

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("languajes")
	err := col.FindId(bson.ObjectIdHex(id)).One(&lang)

	if err != nil {
		return lang, err
	}

	return lang, err
}

func (lr LanguajeRespository) Update(lang models.Languaje) (models.Languaje, error) {

	var interlag models.Languaje
	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("languajes")
	err := col.FindId(lang.Id).One(&interlag)

	if err != nil {
		return interlag, err
	}

	interlag.Name = lang.Name
	interlag.Statu = lang.Statu
	interlag.Logo = lang.Logo

	err = col.UpdateId(interlag.Id, interlag)

	if err != nil {
		return interlag, err
	}

	return interlag, nil

}

func (lr LanguajeRespository) Delete(id string) (bool, error) {

	var lang models.Languaje

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("languajes")

	err := col.FindId(bson.ObjectIdHex(id)).One(&lang)

	if err != nil {
		return false, err
	}

	err = col.RemoveId(lang.Id)

	if err != nil {
		return false, err
	}

	return true, nil

}
