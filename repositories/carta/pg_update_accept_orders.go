package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_AceeptOrders(acceptOrders bool, idmenu int, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE Menu SET acceptOrders=$1 WHERE idMenu=$2 AND idbusiness=$3`

	if _, err_update := db.Exec(context.Background(), query, acceptOrders, idmenu, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
