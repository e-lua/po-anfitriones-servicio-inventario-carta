package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_All(idbusiness int) ([]models.Pg_ScheduleRange, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "SELECT idScheduleRange,name,description,minutePerFraction,numberfractions,startTime,endTime,maxOrders,available,timezone FROM ScheduleRange WHERE idbusiness=$1 AND available=true"
	rows, error_shown := db.Query(ctx, q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListScheduleRange := []models.Pg_ScheduleRange{}

	if error_shown != nil {

		return oListScheduleRange, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oScheduleRange := models.Pg_ScheduleRange{}
		rows.Scan(&oScheduleRange.IDSchedule, &oScheduleRange.Name, &oScheduleRange.Description, &oScheduleRange.MinutePerFraction, &oScheduleRange.NumberOfFractions, &oScheduleRange.StartTime, &oScheduleRange.EndTime, &oScheduleRange.MaxOrders, &oScheduleRange.Available, &oScheduleRange.TimeZone)
		oListScheduleRange = append(oListScheduleRange, oScheduleRange)
	}

	//Si todo esta bien
	return oListScheduleRange, nil

}
