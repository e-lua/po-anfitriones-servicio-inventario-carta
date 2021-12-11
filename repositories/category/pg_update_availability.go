package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_AvailableToFalse(idcategory int, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	q := "UPDATE Category SET available=$1,updateddate=$2 WHERE idcategory=$3 AND idbusiness=$4"
	if _, err_update := db.Exec(context.Background(), q, false, time.Now().Date, idcategory, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
