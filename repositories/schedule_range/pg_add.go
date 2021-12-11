package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(idbusiness int, schedulerange models.Pg_ScheduleRange) error {

	db := models.Conectar_Pg_DB()
	query := `INSERT INTO ScheduleRange(idbusiness,name,description,minutePerFraction,numberfractions,starttime,endtime,maxorders,updateddate) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	if _, err := db.Exec(context.Background(), query, idbusiness, schedulerange.Name, schedulerange.Description, schedulerange.MinutePerFraction, schedulerange.NumberOfFractions, schedulerange.StartTime, schedulerange.EndTime, schedulerange.MaxOrders, time.Now()); err != nil {
		return err
	}

	return nil
}
