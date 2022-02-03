package repositories

import (
	"context"
	"strconv"
	"strings"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Copy_Carta(pg_schedule []models.Pg_ScheduleRange_External, pg_element_external []models.Pg_Element_With_Stock_External, idbusiness int, date string, idcarta int) (int, error) {
	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db_external := models.Conectar_Pg_DB_External()

	//Elementos
	idelement_pg, idcarta_pg, idcategory_pg, namecategory_pg, urlphotocategory_pg, name_pg, price_pg, description_pg, urlphot_pg, typem_pg, stock_pg, idbusiness_pg := []int{}, []int{}, []int{}, []string{}, []string{}, []string{}, []float32{}, []string{}, []string{}, []int{}, []int{}, []int{}
	for _, e := range pg_element_external {
		idelement_pg = append(idelement_pg, e.IDElement)
		idcarta_pg = append(idcarta_pg, idcarta)
		idcategory_pg = append(idcategory_pg, e.IDCategory)
		namecategory_pg = append(namecategory_pg, e.NameCategory)
		urlphotocategory_pg = append(urlphotocategory_pg, e.UrlPhotoCategory)
		name_pg = append(name_pg, e.Name)
		price_pg = append(price_pg, e.Price)
		description_pg = append(description_pg, e.Description)
		urlphot_pg = append(urlphot_pg, e.UrlPhoto)
		typem_pg = append(typem_pg, e.TypeMoney)
		stock_pg = append(stock_pg, e.Stock)
		idbusiness_pg = append(idbusiness_pg, idbusiness)
	}

	//Rango horarios
	idschedule_pg_2, idcartamain_pg_2, idbusinessmain_pg_2, name_pg_2, description_pg_2, minutesperfraction_pg_2, numberfractions_pg_2, start_pg_2, end_pg_2, maxorders_pg_2 := []int64{}, []int{}, []int{}, []string{}, []string{}, []int{}, []int{}, []string{}, []string{}, []int{}

	for _, sch := range pg_schedule {

		idschedule_pg_2 = append(idschedule_pg_2, sch.IDSchedule)
		idcartamain_pg_2 = append(idcartamain_pg_2, idcarta)
		idbusinessmain_pg_2 = append(idbusinessmain_pg_2, idbusiness)
		name_pg_2 = append(name_pg_2, sch.Name)
		description_pg_2 = append(description_pg_2, sch.Description)
		minutesperfraction_pg_2 = append(minutesperfraction_pg_2, sch.MinutePerFraction)
		numberfractions_pg_2 = append(numberfractions_pg_2, sch.NumberOfFractions)
		start_pg_2 = append(start_pg_2, sch.StartTime)
		end_pg_2 = append(end_pg_2, sch.EndTime)
		maxorders_pg_2 = append(maxorders_pg_2, sch.MaxOrders)

	}

	//Lista de actualizacion de rangos horarios
	idschedulerange_pg_3, idcarta_pg_3, idbusiness_pg_3, startime_pg_3, endtime_pg_3, max_orders_3 := []int64{}, []int{}, []int{}, []string{}, []string{}, []int{}

	for _, sch := range pg_schedule {

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
			idschedulerange_pg_3 = append(idschedulerange_pg_3, sch.IDSchedule)
			idcarta_pg_3 = append(idcarta_pg_3, idcarta)
			idbusiness_pg_3 = append(idbusiness_pg_3, idbusiness)
			startime_pg_3 = append(startime_pg_3, hora_ini_string)
			endtime_pg_3 = append(endtime_pg_3, hora_fin_toinsert)
			max_orders_3 = append(max_orders_3, sch.MaxOrders)

			//Nuevo valor de hora de inicio
			new_hora_ini, _ := strconv.Atoi(strconv.Itoa(horas) + minutos_string)
			hora_ini = new_hora_ini
			hora_ini_string = hora_fin_toinsert
		}
	}

	//BEGIN
	tx, error_tx := db_external.Begin(ctx)
	if error_tx != nil {
		return 0, error_tx
	}

	//INSERTAR ELEMENTO
	q_element := `INSERT INTO element(idelement,idcarta,idcategory,namecategory,urlphotcategory,name,price,description,urlphoto,typemoney,stock,idbusiness) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(100)[],$5::varchar(230)[],$6::varchar(100)[],$7::decimal(8,2)[],$8::varchar(250)[],$9::varchar(230)[],$10::int[],$11::int[],$12::int[]));`
	if _, err_insert_element := tx.Exec(ctx, q_element, idelement_pg, idcarta_pg, idcategory_pg, namecategory_pg, urlphotocategory_pg, name_pg, price_pg, description_pg, urlphot_pg, typem_pg, stock_pg, idbusiness_pg); err_insert_element != nil {
		tx.Rollback(ctx)
		return 0, err_insert_element
	}

	//INSERTAR RANGO HORARIO
	q_schedulerange := `INSERT INTO ScheduleRange(idScheduleRange,idbusiness,idcarta,name,description,minuteperfraction,numberfractions,startTime,endTime,maxOrders) (SELECT * FROM unnest($1::int[],$2::int[],$3::int[],$4::varchar(12)[],$5::varchar(60)[],$6::int[],$7::int[],$8::varchar(10)[],$9::varchar(10)[],$10::int[]));`
	if _, err_insert_schedulerange := tx.Exec(ctx, q_schedulerange, idschedule_pg_2, idbusinessmain_pg_2, idcartamain_pg_2, name_pg_2, description_pg_2, minutesperfraction_pg_2, numberfractions_pg_2, start_pg_2, end_pg_2, maxorders_pg_2); err_insert_schedulerange != nil {
		tx.Rollback(ctx)
		return 0, err_insert_schedulerange
	}

	//INSERTAR LISTAS DE RANGOS HORARIOS
	q_listschedulerange := `INSERT INTO ListScheduleRange(idcarta,idschedulemain,idbusiness,starttime,endtime,maxorders) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(6)[],$5::varchar(6)[],$6::int[]))`
	if _, err_listschedulerange := tx.Exec(ctx, q_listschedulerange, idcarta_pg_3, idschedulerange_pg_3, idbusiness_pg_3, startime_pg_3, endtime_pg_3, max_orders_3); err_listschedulerange != nil {
		tx.Rollback(ctx)
		return 0, err_listschedulerange
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return 0, err_commit
	}

	return idcarta, nil
}
