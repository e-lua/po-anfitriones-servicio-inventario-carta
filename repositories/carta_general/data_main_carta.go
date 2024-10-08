package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Main_Data(idbsuiness int) (models.Pg_Category_Element_ScheduleRange_AutomaticDiscount, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var oCES models.Pg_Category_Element_ScheduleRange_AutomaticDiscount

	db := models.Conectar_Pg_DB()

	q := "SELECT (SELECT count(*) FROM category WHERE idbusiness=$1 AND available=true AND isdeleted=false), (SELECT count(*) FROM element e JOIN category c on e.idcategory=c.idcategory WHERE c.idbusiness=$2 AND e.available=true AND e.isdeleted=false),(SELECT count(*) FROM ScheduleRange WHERE idbusiness=$3 AND available=true),(SELECT count(*) FROM AutomaticDiscount WHERE idbusiness=$4)"
	error_showname := db.QueryRow(ctx, q, idbsuiness, idbsuiness, idbsuiness, idbsuiness).Scan(&oCES.Category, &oCES.Element, &oCES.Schedule, &oCES.AutomaticDiscount)

	if error_showname != nil {

		return oCES, error_showname
	}

	//Si todo esta bien
	return oCES, nil
}
