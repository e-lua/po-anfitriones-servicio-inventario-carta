package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_ScheduleRange_ToCreate(date string, idbusiness int) ([]models.Pg_Schedule_ToCreate, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB_External()

	q := "SELECT ls.idschedule,c.date::date::varchar(12),ls.starttime::time::varchar(5),ls.endtime::time::varchar(5),ls.timezone,ls.maxorders,CONCAT(ls.starttime,' - ',ls.endtime) FROM listschedulerange ls LEFT JOIN carta c ON ls.idcarta=c.idcarta WHERE c.date::date=$1::date AND ls.idbusiness=$2 AND ls.maxorders>0 AND concat(ls.starttime,'-',(ls.timezone::integer*-1)::varchar(3))::time with time zone > NOW()::time at time zone CONCAT('UTC',(ls.timezone::integer*-1)::varchar(3)) ORDER BY REPLACE(ls.starttime,':','')::int  ASC"
	rows, error_shown := db.Query(ctx, q, date, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oScheduleList []models.Pg_Schedule_ToCreate

	if error_shown != nil {

		return oScheduleList, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oSchedule models.Pg_Schedule_ToCreate
		rows.Scan(&oSchedule.IDSchedule, &oSchedule.Date, &oSchedule.Starttime, &oSchedule.Endtime, &oSchedule.TimeZone, &oSchedule.MaxOrders, &oSchedule.ShowToComensal)
		oScheduleList = append(oScheduleList, oSchedule)
	}

	//Si todo esta bien
	return oScheduleList, nil
}
