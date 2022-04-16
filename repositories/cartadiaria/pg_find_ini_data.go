package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_IniData(date string, idbusiness int) (models.Pg_Carta_External, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var carta_ini_data models.Pg_Carta_External

	db := models.Conectar_Pg_DB_External()
	q := "SELECT idcarta,date,availableorders,visible,(SELECT COUNT(*) FROM Element e JOIN Carta c on e.idcarta=c.idcarta WHERE c.date=$1 AND c.idbusiness=$2),(SELECT COUNT(*) FROM ScheduleRange sr JOIN Carta c on sr.idcarta=c.idcarta WHERE c.date=$1 AND c.idbusiness=$2) FROM Carta WHERE date=$1 AND idbusiness=$2"
	error_shown := db.QueryRow(ctx, q, date, idbusiness).Scan(&carta_ini_data.IDCarta, &carta_ini_data.Date, &carta_ini_data.AvailableForOrders, &carta_ini_data.Visible, &carta_ini_data.Elements, &carta_ini_data.ScheduleRanges)

	if error_shown != nil {
		return carta_ini_data, error_shown
	}

	//Si todo esta bien
	return carta_ini_data, nil
}
