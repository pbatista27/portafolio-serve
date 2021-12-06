package helpers

import (
	"github.com/pbatista27/portafolio-serve/db"
	"gopkg.in/mgo.v2/bson"
)

func ValidarCorreoExiste(email string) (bool, error) {

	con := db.NewDbSession()
	defer con.Close()

	col := con.DbCollection("users")

	n, err := col.Find(bson.M{"email": email}).Count()

	if err != nil {
		return true, err

	}

	if n >= 1 {
		return true, nil
	}

	return false, nil

}
