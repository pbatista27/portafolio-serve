package data

import (
	"errors"

	"github.com/pbatista27/portafolio-serve/db"
	"github.com/pbatista27/portafolio-serve/models"
	"gopkg.in/mgo.v2/bson"
)

type StudyRepository struct{}

func (sr StudyRepository) GelAll() ([]models.Study, error) {

	var estudios []models.Study
	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("studys")
	err := col.Find(bson.M{}).All(&estudios)

	if err != nil {
		return estudios, err
	}

	return estudios, nil
}

func (sr StudyRepository) Show(id string) (models.Study, error) {

	var study models.Study

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("studys")
	err := col.FindId(bson.ObjectIdHex(id)).One(&study)

	if err != nil {
		return study, err
	}

	return study, nil
}

func (sr StudyRepository) Create(study models.Study) (models.Study, error) {

	study.Id = bson.NewObjectId()

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("studys")
	err := col.Insert(study)

	if err != nil {
		return study, err
	}

	return study, nil
}

func (sr StudyRepository) Update(study models.Study) error {

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("studys")

	cont, err := col.FindId(study.Id).Count()

	if err != nil {
		return err
	}

	if cont <= 0 {
		return errors.New("estudio no encontrado")
	}

	err = col.UpdateId(study.Id, study)
	if err != nil {
		return err
	}

	return nil
}

func (sr StudyRepository) Delete(id string) error {

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("studys")
	err := col.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return err
	}

	return nil

}
