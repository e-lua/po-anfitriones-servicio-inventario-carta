package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Data(shcedule_range models.Pg_ScheduleRange, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	q := "UPDATE ScheduleRange SET name=$1,description=$2,minutePerFraction=$3,startTime=$4,endTime=$5,maxOrders=$6,updateddate=$7,numberfractions=$8 WHERE idScheduleRange=$8 AND idbusiness=$9"
	if _, err_update := db.Exec(context.Background(), q, shcedule_range.Name, shcedule_range.Description, shcedule_range.MinutePerFraction, shcedule_range.StartTime, shcedule_range.EndTime, shcedule_range.MaxOrders, time.Now(), shcedule_range.IDSchedule, idbusiness, shcedule_range.NumberOfFractions); err_update != nil {
		return err_update
	}

	return nil
}
