package repositories

import (
	"context"
	"strconv"
	"strings"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_ScheduleRange_List(pg_schedule []models.Pg_ScheduleRange_External, idcarta int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB_External()

	//Rango horarios
	idschedule_pg, idcartamain_pg, idbusinessmain_pg, name_pg, description_pg, minutesperfraction_pg, numberfractions_pg, start_pg, end_pg, maxorders_pg := []int{}, []int{}, []int{}, []string{}, []string{}, []int{}, []int{}, []string{}, []string{}, []int{}

	//Lista de actualizacion de rangos horarios
	idschedulerange_pg, idcarta_pg, idbusiness_pg, startime_pg, endtime_pg, max_orders := []int{}, []int{}, []int{}, []string{}, []string{}, []int{}

	for _, sch := range pg_schedule {

		go func() {
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
		}()

		for i := 0; i < sch.NumberOfFractions; i++ {

			arr := strings.Split(sch.StartTime, ":")
			hora_ini, _ := strconv.Atoi(arr[0] + arr[1][:2])

			//Inicio de bucle para obtener la hora fin
			hora_fin := strconv.Itoa(hora_ini + sch.MinutePerFraction)
			var index int
			if len(hora_fin) > 3 {
				index = 2
			} else {
				index = 1
			}
			hora_fin_toinsert := hora_fin[:index] + ":" + hora_fin[index:]
			//Fin de bucle para obtener la hora fin

			//Insertamos los datos en el modelo
			idschedulerange_pg = append(idschedulerange_pg, sch.IDSchedule)
			idcarta_pg = append(idcarta_pg, idcarta)
			idbusiness_pg = append(idbusiness_pg, idbusiness)
			startime_pg = append(startime_pg, sch.StartTime)
			endtime_pg = append(endtime_pg, hora_fin_toinsert)
			max_orders = append(max_orders, sch.MaxOrders)
		}
	}

	q := `BEGIN;
	DELETE FROM ScheduleRange WHERE idbusiness=$1;
	DELETE FROM ListScheduleRange WHERE idbusiness=$2; 
	INSERT INTO ScheduleRange(idScheduleRange,idbusiness,idCarta,name,description,minuteperfraction,numberfractions,startTime,endTime,maxOrders) (SELECT * FROM unnest($3::int[],$4::int[],$5::int[],$6::string[],$7::string[],$8::int[],$9::int[],$10::varchar(10)[],$11::varchar(10)[],$12::int[]));
	INSERT INTO ListScheduleRange(idcarta,idschedulemain,idbusiness,starttime,endtime,maxorders) (select * from unnest($13::int[],$14::int[],$15::int[],$16::string[],$17::string[],$18::int[]));
	ROLLBACK;`
	if _, err_update := db_external.Exec(context.Background(), q, idbusiness, idbusiness, idschedule_pg, idcartamain_pg, idbusinessmain_pg, name_pg, description_pg, minutesperfraction_pg, numberfractions_pg, start_pg, end_pg, maxorders_pg, idschedulerange_pg, idcarta_pg, idbusiness_pg, startime_pg, endtime_pg, max_orders); err_update != nil {
		return err_update
	}

	return nil
}
