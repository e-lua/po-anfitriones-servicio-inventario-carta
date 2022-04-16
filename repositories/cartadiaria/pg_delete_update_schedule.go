package repositories

import (
	"context"
	"strconv"
	"strings"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Delete_Update_ScheduleRange(pg_schedule []models.Pg_ScheduleRange_External, idcarta int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()
	db_external := models.Conectar_Pg_DB_External()

	//Rango horarios
	idschedule_pg, idcartamain_pg, idbusinessmain_pg, name_pg, description_pg, minutesperfraction_pg, numberfractions_pg, start_pg, end_pg, maxorders_pg, timezone_pg := []int64{}, []int{}, []int{}, []string{}, []string{}, []int{}, []int{}, []string{}, []string{}, []int{}, []string{}
	for _, sch := range pg_schedule {

		if "-5" != "" {
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
			timezone_pg = append(timezone_pg, "-5")
		}
	}

	//Lista de actualizacion de rangos horarios
	idschedulerange_pg_2, idcarta_pg_2, idbusiness_pg_2, startime_pg_2, endtime_pg_2, max_orders_2, timezone_2 := []int64{}, []int{}, []int{}, []string{}, []string{}, []int{}, []string{}

	for _, sch := range pg_schedule {

		if "-5" != "" {
			arr := strings.Split(sch.StartTime, ":")
			hora_ini := 0
			hora_ini_string := sch.StartTime

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
				idschedulerange_pg_2 = append(idschedulerange_pg_2, sch.IDSchedule)
				idcarta_pg_2 = append(idcarta_pg_2, idcarta)
				idbusiness_pg_2 = append(idbusiness_pg_2, idbusiness)
				startime_pg_2 = append(startime_pg_2, hora_ini_string)
				endtime_pg_2 = append(endtime_pg_2, hora_fin_toinsert)
				max_orders_2 = append(max_orders_2, sch.MaxOrders)
				timezone_2 = append(timezone_2, "-5")

				//Nuevo valor de hora de inicio
				new_hora_ini, _ := strconv.Atoi(strconv.Itoa(horas) + minutos_string)
				hora_ini = new_hora_ini
				hora_ini_string = hora_fin_toinsert
			}
		}
	}

	//BEGIN
	tx, error_tx := db_external.Begin(ctx)
	if error_tx != nil {
		return error_tx
	}

	//ELIMINAR LISTA DE RANGOS HORARIOS
	q_delete_list := `DELETE FROM ListScheduleRange WHERE idbusiness=$1 AND idcarta=$2`
	if _, err_update := tx.Exec(ctx, q_delete_list, idbusiness, idcarta); err_update != nil {
		tx.Rollback(ctx)
		return err_update
	}

	//ELIMINAR RANGO HORARIO
	q_2 := `DELETE FROM ScheduleRange WHERE idbusiness=$1 AND idcarta=$2`
	_, err_update := tx.Exec(ctx, q_2, idbusiness, idcarta)
	if err_update != nil {
		tx.Rollback(ctx)
		return err_update
	}

	//RANGO HORARIO
	q_schedulerange := `INSERT INTO ScheduleRange(idScheduleRange,idbusiness,idcarta,name,description,minuteperfraction,numberfractions,startTime,endTime,maxOrders,timezone) (SELECT * FROM unnest($1::int[],$2::int[],$3::int[],$4::varchar(12)[],$5::varchar(60)[],$6::int[],$7::int[],$8::varchar(10)[],$9::varchar(10)[],$10::int[],$11::varchar(3)[]));`
	if _, err_schedule := tx.Exec(ctx, q_schedulerange, idschedule_pg, idbusinessmain_pg, idcartamain_pg, name_pg, description_pg, minutesperfraction_pg, numberfractions_pg, start_pg, end_pg, maxorders_pg, timezone_pg); err_schedule != nil {
		tx.Rollback(ctx)
		return err_schedule
	}

	//LISTA RANGOS HORARIOS
	q_listschedule := `INSERT INTO ListScheduleRange(idcarta,idschedulemain,idbusiness,starttime,endtime,maxorders,timezone) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(6)[],$5::varchar(6)[],$6::int[],$7::varchar(3)[]))`
	if _, err_listschedule := tx.Exec(ctx, q_listschedule, idcarta_pg_2, idschedulerange_pg_2, idbusiness_pg_2, startime_pg_2, endtime_pg_2, max_orders_2, timezone_2); err_listschedule != nil {
		tx.Rollback(ctx)
		return err_listschedule
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	return nil
}
