package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Delete_Vencidas() error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db_external := models.Conectar_Pg_DB_External()

	//BEGIN
	tx, error_tx := db_external.Begin(ctx)
	if error_tx != nil {
		return error_tx
	}

	//ELIMINAMOS EL ELEMENTO
	query_element := `DELETE FROM element AS ele JOIN carta AS car ON ele.idcarta=car.idcarta WHERE car.deleteddate<=NOW()`
	_, error_element := tx.Exec(ctx, query_element)
	if error_element != nil {
		tx.Rollback(ctx)
		return error_element
	}

	//ELIMINAMOS LA LISTA DE RANGO HORARIOS
	query_listschedulerange := `DELETE FROM listschedulerange AS lsch JOIN carta AS car ON lsch.idcarta=car.idcarta WHERE car.deleteddate<=NOW()`
	_, error_listschedulerange := tx.Exec(ctx, query_listschedulerange)
	if error_listschedulerange != nil {
		tx.Rollback(ctx)
		return error_listschedulerange
	}

	//ELIMINAMOS EL RANGO HORARIO
	query_schedulerange := `DELETE FROM schedulerange AS sch JOIN carta AS car ON sch.idcarta=car.idcarta WHERE car.deleteddate<=NOW()`
	_, error_schedule := tx.Exec(ctx, query_schedulerange)
	if error_schedule != nil {
		tx.Rollback(ctx)
		return error_schedule
	}

	//ELIMINAMOS LA CARTA
	query_carta := `DELETE FROM carta WHERE carta.deleteddate<=NOW()`
	_, error_carta := tx.Exec(ctx, query_carta)
	if error_carta != nil {
		tx.Rollback(ctx)
		return error_carta
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	return nil
}
