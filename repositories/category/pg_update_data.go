package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Data(idcategory int, idbusiness int, name string) error {

	db := models.Conectar_Pg_DB()

	q := "UPDATE Category SET name=$1,updateddate=$2 WHERE idcategory=$3 AND idbusiness=$4"
	if _, err_update := db.Exec(context.Background(), q, name, time.Now(), idcategory, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
