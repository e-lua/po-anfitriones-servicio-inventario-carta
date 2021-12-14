package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Delete_ScheduleRange_List(idcarta int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB_External()

	q := `BEGIN;
	DELETE FROM ScheduleRange WHERE idbusiness=$1 AND idcarta=$2;
	DELETE FROM ListScheduleRange WHERE idbusiness=$3 AND idcarta=$4; 
	ROLLBACK;`
	if _, err_update := db_external.Exec(context.Background(), q, idbusiness, idcarta, idbusiness, idcarta); err_update != nil {
		return err_update
	}

	return nil
}
