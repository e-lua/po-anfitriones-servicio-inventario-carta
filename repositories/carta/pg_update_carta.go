package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_AvailableToFalse(available bool, visible bool, idcarta int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB_External()

	q := "UPDATE Carta set availableorders=$1,visible=$2,updateddate=$3 WHERE idcarta=$4 AND idbusiness=$5"
	if _, err_update := db_external.Exec(context.Background(), q, available, visible, time.Now(), idcarta, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
