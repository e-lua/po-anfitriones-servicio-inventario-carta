package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_One_Element(stock int, idelement int, idcarta int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB_External()

	q := `UPDATE Element SET stock=$1 WHERE idbusiness=$2 AND idcarta=$3 AND idelement=$4`
	if _, err_update := db_external.Exec(context.Background(), q, stock, idbusiness, idcarta, idelement); err_update != nil {
		return err_update
	}

	return nil
}
