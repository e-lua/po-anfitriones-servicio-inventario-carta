package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Delete_Elements(idcarta int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB_External()

	q := `DELETE FROM element WHERE idbusiness=$1 AND idcarta=$2`
	if _, err_update := db_external.Exec(context.Background(), q, idbusiness, idcarta); err_update != nil {
		return err_update
	}

	return nil
}
