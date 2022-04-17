package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Available_Visible(available bool, visible bool, idcarta int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db_external := models.Conectar_Pg_DB_External()

	//BEGIN
	tx, error_tx := db_external.Begin(ctx)
	if error_tx != nil {
		tx.Rollback(ctx)
		return error_tx
	}

	q_carta := "UPDATE Carta set availableorders=$1,visible=$2,updateddate=$3 WHERE idcarta=$4 AND idbusiness=$5"
	if _, err_update := tx.Exec(ctx, q_carta, available, visible, time.Now(), idcarta, idbusiness); err_update != nil {
		return err_update
	}

	q := "UPDATE Element set availableorders=$1 WHERE idcarta=$2 AND idbusiness=$3"
	if _, err_update := tx.Exec(ctx, q, available, idcarta, idbusiness); err_update != nil {
		return err_update
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	return nil
}
