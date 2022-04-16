package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_One_Element(stock int, idelement int, idcarta int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db_external := models.Conectar_Pg_DB_External()

	q := `UPDATE Element SET stock=$1 WHERE idbusiness=$2 AND idcarta=$3 AND idelement=$4`
	if _, err_update := db_external.Exec(ctx, q, stock, idbusiness, idcarta, idelement); err_update != nil {
		return err_update
	}

	return nil
}
