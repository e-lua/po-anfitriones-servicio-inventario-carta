package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Delete(idbusiness int, idcarta int) error {

	db_external := models.Conectar_Pg_DB_External()

	query := `BEGIN;DELETE FROM carta WHERE carta.idbusiness=$1 AND carta.idcarta=$2 ;DELETE FROM element USING carta WHERE element.idcarta = (SELECT distinct (e.idcarta) FROM element as e JOIN carta as ca ON e.idbusiness=ca.idbusiness WHERE ca.idbusiness=$1 AND e.idcarta=$2);DELETE FROM listschedulerange WHERE listschedulerange.idcarta=(SELECT distinct (ls.idcarta) FROM listschedulerange as ls JOIN carta as ca ON ls.idbusiness=ca.idbusiness WHERE ca.idbusiness=$1 AND ls.idcarta=$2);DELETE FROM schedulerange WHERE schedulerange.idcarta = (SELECT distinct (s.idcarta) FROM schedulerange as s JOIN carta as ca ON s.idbusiness=ca.idbusiness WHERE ca.idbusiness=$1 AND s.idcarta=$2);END;`
	_, err := db_external.Query(context.Background(), query, idbusiness, idcarta)

	if err != nil {
		return err
	}

	return nil
}
