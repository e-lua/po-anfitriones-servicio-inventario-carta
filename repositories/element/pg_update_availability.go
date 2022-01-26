package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Available(status bool, idelement int, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	q := "UPDATE Element SET available=$1,updateddate=$2 FROM Category WHERE idelement=$3 AND Category.idbusiness=$4"
	if _, err_update := db.Exec(context.TODO(), q, status, time.Now(), idelement, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
