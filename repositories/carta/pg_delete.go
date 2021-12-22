package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Delete(idbusiness int, idcarta int) error {

	db_external := models.Conectar_Pg_DB_External()

	//BEGIN
	tx, error_tx := db_external.Begin(context.Background())
	if error_tx != nil {
		return error_tx
	}

	//ELIMINAMOS LA CARTA
	query_carta := `DELETE FROM carta WHERE carta.idbusiness=$1 AND carta.idcarta=$2`
	_, error_carta := tx.Query(context.Background(), query_carta, idbusiness, idcarta)
	if error_carta != nil {
		return error_carta
	}

	//ELIMINAMOS EL ELEMENTO
	query_element := `DELETE FROM element USING carta WHERE element.idcarta = (SELECT distinct (e.idcarta) FROM element as e JOIN carta as ca ON e.idbusiness=ca.idbusiness WHERE ca.idbusiness=$1 AND e.idcarta=$2)`
	_, error_element := tx.Query(context.Background(), query_element, idbusiness, idcarta)
	if error_element != nil {
		return error_element
	}

	//ELIMINAMOS LA LISTA DE RANGO HORARIOS
	query_listschedulerange := `DELETE FROM listschedulerange WHERE listschedulerange.idcarta=(SELECT distinct (ls.idcarta) FROM listschedulerange as ls JOIN carta as ca ON ls.idbusiness=ca.idbusiness WHERE ca.idbusiness=$1 AND ls.idcarta=$2)`
	_, error_listschedulerange := tx.Query(context.Background(), query_listschedulerange, idbusiness, idcarta)
	if error_listschedulerange != nil {
		return error_listschedulerange
	}

	//ELIMINAMOS EL RANGO HORARIO
	query_schedulerange := `DELETE FROM schedulerange WHERE schedulerange.idcarta = (SELECT distinct (s.idcarta) FROM schedulerange as s JOIN carta as ca ON s.idbusiness=ca.idbusiness WHERE ca.idbusiness=$1 AND s.idcarta=$2)`
	_, error_schedule := tx.Query(context.Background(), query_schedulerange, idbusiness, idcarta)
	if error_schedule != nil {
		return error_schedule
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(context.Background())
	if err_commit != nil {
		return err_commit
	}

	return nil
}
