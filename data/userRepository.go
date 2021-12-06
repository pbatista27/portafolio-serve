package data

import (
	"github.com/pbatista27/portafolio-serve/common"
	"github.com/pbatista27/portafolio-serve/db"
	"github.com/pbatista27/portafolio-serve/models"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct{}

func (ur UserRepository) CrearUsuario(u *models.User) error {

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("users")

	u.Id = bson.NewObjectId()

	err := col.Insert(&u)

	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) Login(email, password string) (models.User, bool, error) {

	var mu models.User
	var validar bool

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("users")
	err := col.Find(bson.M{"email": email}).One(&mu)

	if err != nil {
		return mu, false, err
	}

	validar, err = common.NewBcrypt().Descriptar(password, mu.Password)

	if err != nil {
		return mu, false, err
	}

	if !validar {
		return mu, false, err
	}

	return mu, true, nil
}
