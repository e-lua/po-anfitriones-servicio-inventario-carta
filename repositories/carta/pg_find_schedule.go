package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_ScheduleRanges(idcarta int, idbusiness int) ([]models.Pg_ScheduleRange_External, error) {

	db := models.Conectar_Pg_DB_External()

	q := "SELECT idschedulerange,name,description,minuteperfraction,numberfractions,starttime,endtime,maxorders FROM schedulerange WHERE idcarta=$1 AND idbusiness=$2"
	rows, error_shown := db.Query(context.Background(), q, idcarta, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListSchedule []models.Pg_ScheduleRange_External

	if error_shown != nil {

		return oListSchedule, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oSchedule models.Pg_ScheduleRange_External
		rows.Scan(&oSchedule.IDSchedule, &oSchedule.Name, &oSchedule.Description, &oSchedule.MinutePerFraction, &oSchedule.NumberOfFractions, &oSchedule.StartTime, &oSchedule.EndTime, &oSchedule.MaxOrders)
		oListSchedule = append(oListSchedule, oSchedule)
	}

	//Si todo esta bien
	return oListSchedule, nil
}
