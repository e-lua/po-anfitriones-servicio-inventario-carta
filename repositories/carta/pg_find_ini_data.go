package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_IniData(date string, idbusiness int) (models.Pg_Carta_External, error) {

	var carta_ini_data models.Pg_Carta_External

	db := models.Conectar_Pg_DB_External()
	q := "SELECT c.idcarta,c.date,c.availableorders,c.visible,COUNT(e.idelement),COUNT(s.idschedulerange) FROM carta c LEFT JOIN Element e on c.idcarta=e.idcarta LEFT JOIN schedulerange s on c.idcarta=s.idcarta WHERE c.date=$1 AND c.idbusiness=$2 GROUP BY c.idcarta"
	error_shown := db.QueryRow(context.Background(), q, date, idbusiness).Scan(&carta_ini_data.IDCarta, &carta_ini_data.Date, &carta_ini_data.AvailableForOrders, &carta_ini_data.Visible, &carta_ini_data.Elements, &carta_ini_data.ScheduleRanges)

	if error_shown != nil {
		return carta_ini_data, error_shown
	}

	//Si todo esta bien
	return carta_ini_data, nil
}
