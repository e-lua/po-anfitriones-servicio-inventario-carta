package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Availability(available bool, idmenu int, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE Menu SET isavailable=$1 WHERE idMenu=$2 AND idbusiness=$3`

	if _, err_update := db.Exec(context.Background(), query, available, idmenu, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
