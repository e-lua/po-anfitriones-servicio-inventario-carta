package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Delete(idschedulerange int, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	q := "DELETE FROM ScheduleRange WHERE idschedulerange=$1 AND idbusiness=$2"
	if _, err_update := db.Exec(context.Background(), q, idschedulerange, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
