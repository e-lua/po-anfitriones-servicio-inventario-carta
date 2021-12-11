package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_AvailableToTrue(idcategory int, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	q := "UPDATE Category SET available=true,updateddate=$1 WHERE idcategory=$2 AND idbusiness=$3"
	if _, err_update := db.Exec(context.Background(), q, time.Now().Date, idcategory, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
