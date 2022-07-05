package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(idbusiness int, schedulerange models.Pg_ScheduleRange) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `INSERT INTO ScheduleRange(idbusiness,name,description,minutePerFraction,numberfractions,starttime,endtime,maxorders,updateddate,timezone) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	if _, err := db.Exec(ctx, query, idbusiness, schedulerange.Name, schedulerange.Description, schedulerange.MinutePerFraction, schedulerange.NumberOfFractions, schedulerange.StartTime, schedulerange.EndTime, schedulerange.MaxOrders, time.Now(), "-5"); err != nil {
		return err
	}

	return nil
}
