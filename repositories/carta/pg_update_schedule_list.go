package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_ScheduleRange_List(pg_schedule []models.Pg_ScheduleRange_External, idcarta int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB_External()

	//Lista de actualizacion de rangos horarios
	idschedulerange_pg, idcarta_pg, idbusiness_pg, startime_pg, endtime_pg, max_orders := []int{}, []int{}, []int{}, []time.Time{}, []time.Time{}, []int{}

	for _, sch := range pg_schedule {

		for i := 0; i < sch.NumberOfFractions; i++ {

			tiempo_inicial, _ := time.Parse("15:04:00", sch.StartTime)

			//Añadimos el tiempo
			tiempo_fin := tiempo_inicial.Add(time.Minute * time.Duration(sch.MinutePerFraction))

			/*var minutos_string string

			tiempo,_:=time.Parse("15:04",sch.StartTime)

			t:=tiempo.Add(time.Minute*30)
			t.Clock()
			//Minutos y Horas
			minutos, _ := strconv.Atoi(arr[1][:2])
			horas, _ := strconv.Atoi(arr[0])

			//Validamos que no sobrepase los 60 minutos
			if minutos > 60 {
				minutos_string = "00"
				horas = horas + 1
			} else {
				minutos_string = arr[1][:2]
			}

			//Horas con minutos juntos
			hora_ini, _ := strconv.Atoi(strconv.Itoa(horas) + minutos_string)

			//Inicio de bucle para obtener la hora fin
			hora_fin := strconv.Itoa(hora_ini + sch.MinutePerFraction)
			var index int
			if len(hora_fin) > 3 {
				index = 2
			} else {
				index = 1
			}
			hora_fin_toinsert := hora_fin[:index] + ":" + hora_fin[index:]
			sch.StartTime = hora_fin_toinsert*/
			//Fin de bucle para obtener la hora fin

			//Insertamos los datos en el modelo
			idschedulerange_pg = append(idschedulerange_pg, sch.IDSchedule)
			idcarta_pg = append(idcarta_pg, idcarta)
			idbusiness_pg = append(idbusiness_pg, idbusiness)
			startime_pg = append(startime_pg, tiempo_inicial)
			endtime_pg = append(endtime_pg, tiempo_fin)
			max_orders = append(max_orders, sch.MaxOrders)

			tiempo_inicial = tiempo_fin

		}
	}

	q := `INSERT INTO ListScheduleRange(idcarta,idschedulemain,idbusiness,starttime,endtime,maxorders) (select * from unnest($1::int[],$2::int[],$3::int[],$4::time[],$5::time[],$6::int[]))`
	if _, err_update := db_external.Exec(context.Background(), q, idcarta_pg, idschedulerange_pg, idbusiness_pg, startime_pg, endtime_pg, max_orders); err_update != nil {
		return err_update
	}

	return nil
}
