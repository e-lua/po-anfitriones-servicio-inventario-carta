package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Data(shcedule_range models.Pg_ScheduleRange, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "UPDATE ScheduleRange SET name=$1,description=$2,minutePerFraction=$3,startTime=$4,endTime=$5,maxOrders=$6,updateddate=$7,numberfractions=$8,timezone=$9 WHERE idScheduleRange=$10 AND idbusiness=$11"
	if _, err_update := db.Exec(ctx, q, shcedule_range.Name, shcedule_range.Description, shcedule_range.MinutePerFraction, shcedule_range.StartTime, shcedule_range.EndTime, shcedule_range.MaxOrders, time.Now(), shcedule_range.NumberOfFractions, shcedule_range.TimeZone, shcedule_range.IDSchedule, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
