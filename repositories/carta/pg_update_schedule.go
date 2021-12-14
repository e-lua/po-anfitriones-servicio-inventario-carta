package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_ScheduleRange(pg_schedule []models.Pg_ScheduleRange_External, idcarta int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB_External()

	//Rango horarios
	idschedule_pg, idcartamain_pg, idbusinessmain_pg, name_pg, description_pg, minutesperfraction_pg, numberfractions_pg, start_pg, end_pg, maxorders_pg := []int{}, []int{}, []int{}, []string{}, []string{}, []int{}, []int{}, []string{}, []string{}, []int{}

	for _, sch := range pg_schedule {

		idschedule_pg = append(idschedule_pg, sch.IDSchedule)
		idcartamain_pg = append(idcartamain_pg, idcarta)
		idbusinessmain_pg = append(idbusinessmain_pg, idbusiness)
		name_pg = append(name_pg, sch.Name)
		description_pg = append(description_pg, sch.Description)
		minutesperfraction_pg = append(minutesperfraction_pg, sch.MinutePerFraction)
		numberfractions_pg = append(numberfractions_pg, sch.NumberOfFractions)
		start_pg = append(start_pg, sch.StartTime)
		end_pg = append(end_pg, sch.EndTime)
		maxorders_pg = append(maxorders_pg, sch.MaxOrders)

	}

	q := `INSERT INTO ScheduleRange(idScheduleRange,idbusiness,idCarta,name,description,minuteperfraction,numberfractions,startTime,endTime,maxOrders) (SELECT * FROM unnest($1::int[],$2::int[],$3::int[],$4::varchar(6)[],$5::varchar(6)[],$6::int[],$7::int[],$8::varchar(10)[],$9::varchar(10)[],$10::int[]));`
	if _, err_update := db_external.Exec(context.Background(), q, idschedule_pg, idcartamain_pg, idbusinessmain_pg, name_pg, description_pg, minutesperfraction_pg, numberfractions_pg, start_pg, end_pg, maxorders_pg); err_update != nil {
		return err_update
	}

	return nil
}
