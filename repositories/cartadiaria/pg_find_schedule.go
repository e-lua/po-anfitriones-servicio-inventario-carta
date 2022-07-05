package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_ScheduleRanges(idcarta int, idbusiness int) ([]models.Pg_ScheduleRange_External, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB(1)
	q := "SELECT idschedulerange,name,description,minuteperfraction,numberfractions,starttime,endtime,maxorders,timezone FROM schedulerange WHERE idcarta=$1 AND idbusiness=$2"
	rows, error_shown := db.Query(ctx, q, idcarta, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListSchedule []models.Pg_ScheduleRange_External

	if error_shown != nil {

		return oListSchedule, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oSchedule models.Pg_ScheduleRange_External
		rows.Scan(&oSchedule.IDSchedule, &oSchedule.Name, &oSchedule.Description, &oSchedule.MinutePerFraction, &oSchedule.NumberOfFractions, &oSchedule.StartTime, &oSchedule.EndTime, &oSchedule.MaxOrders, &oSchedule.TimeZone)
		oListSchedule = append(oListSchedule, oSchedule)
	}

	//Si todo esta bien
	return oListSchedule, nil
}
