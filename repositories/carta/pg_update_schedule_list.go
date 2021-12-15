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

		arr := strings.Split(sch.StartTime, ":")
		hora_ini := 0

		for i := 0; i < sch.NumberOfFractions; i++ {

			if i == 0 {
				//TODO SOBRE LA HORA DE INICIO
				hora_ini_c, _ := strconv.Atoi(arr[0] + arr[1][:2])
				hora_ini = hora_ini_c
			}

			//TODO SOBRE LA HORA PRE FIN
			hora_pre_fin := strconv.Itoa(hora_ini + sch.MinutePerFraction)

			var index_pre_fin int
			if len(hora_pre_fin) > 3 {
				index_pre_fin = 2
			} else {
				index_pre_fin = 1
			}

			//Minutos y Horas
			minutos, _ := strconv.Atoi(hora_pre_fin[index_pre_fin:])
			horas, _ := strconv.Atoi(hora_pre_fin[:index_pre_fin])

			//Validamos que no sobrepase los 60 minutos
			var minutos_string string
			if minutos > 59 {
				minutos = 60 - minutos
				if minutos < 10 {
					minutos_string = "0" + strconv.Itoa(minutos)
				} else {
					minutos_string = strconv.Itoa(minutos)
				}
				horas = horas + 1
			} else {
				minutos_string = hora_pre_fin[index_pre_fin:]
			}
			hora_finaliza := strconv.Itoa(horas) + minutos_string

			//TODO SOBRE LA HORA FIN
			var index_fin int
			if len(hora_finaliza) > 3 {
				index_fin = 2
			} else {
				index_fin = 1
			}
			hora_fin_toinsert := hora_finaliza[:index_fin] + ":" + hora_finaliza[index_fin:]

			//Fin de bucle para obtener la hora fin

			//Insertamos los datos en el modelo
			idschedulerange_pg = append(idschedulerange_pg, sch.IDSchedule)
			idcarta_pg = append(idcarta_pg, idcarta)
			idbusiness_pg = append(idbusiness_pg, idbusiness)
			startime_pg = append(startime_pg, sch.StartTime)
			endtime_pg = append(endtime_pg, hora_fin_toinsert)
			max_orders = append(max_orders, sch.MaxOrders)

			//Nuevo valor de hora de inicio
			new_hora_ini, _ := strconv.Atoi(strconv.Itoa(horas) + minutos_string)
			hora_ini = new_hora_ini
		}
	}

	q := `INSERT INTO ListScheduleRange(idcarta,idschedulemain,idbusiness,starttime,endtime,maxorders) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(6)[],$5::varchar(6)[],$6::int[]))`
	if _, err_update := db_external.Exec(context.Background(), q, idcarta_pg, idschedulerange_pg, idbusiness_pg, startime_pg, endtime_pg, max_orders); err_update != nil {
		return err_update
	}

	return nil
}
