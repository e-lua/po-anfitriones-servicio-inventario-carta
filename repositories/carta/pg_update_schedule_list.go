package repositories

import (
	"context"
	"strconv"
	"strings"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_ScheduleRange_List(pg_schedule []models.Pg_ScheduleRange_External, idcarta int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB_External()

	//Lista de actualizacion de rangos horarios
	idschedulerange_pg, idcarta_pg, idbusiness_pg, startime_pg, endtime_pg, max_orders := []int{}, []int{}, []int{}, []string{}, []string{}, []int{}

	for _, sch := range pg_schedule {

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

	q := `INSERT INTO ListScheduleRange(idcarta,idschedulemain,idbusiness,starttime,endtime,maxorders) (select * from unnest($11::int[],$12::int[],$13::int[],$14::string[],$15::string[],$16::int[]))`
	if _, err_update := db_external.Exec(context.Background(), q, idschedulerange_pg, idcarta_pg, idbusiness_pg, startime_pg, endtime_pg, max_orders); err_update != nil {
		return err_update
	}

	return nil
}
