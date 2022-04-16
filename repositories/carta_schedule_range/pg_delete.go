package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Delete(idschedulerange int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "DELETE FROM ScheduleRange WHERE idschedulerange=$1 AND idbusiness=$2"
	if _, err_update := db.Exec(ctx, q, idschedulerange, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
