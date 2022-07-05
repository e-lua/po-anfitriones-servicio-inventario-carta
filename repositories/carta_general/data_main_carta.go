package repositories

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Main_Data(idbsuiness int) (models.Pg_Category_Element_ScheduleRange, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var oCES models.Pg_Category_Element_ScheduleRange

	db := models.Conectar_Pg_DB(rand.Intn(5))

	q := "SELECT (SELECT count(*) FROM category WHERE idbusiness=$1 AND available=true), (SELECT count(*) FROM element e JOIN category c on e.idcategory=c.idcategory WHERE c.idbusiness=$2 AND e.available=true),(SELECT count(*) FROM ScheduleRange WHERE idbusiness=$3 AND available=true)"
	error_showname := db.QueryRow(ctx, q, idbsuiness, idbsuiness, idbsuiness).Scan(&oCES.Category, &oCES.Element, &oCES.Schedule)

	if error_showname != nil {

		return oCES, error_showname
	}

	//Si todo esta bien
	return oCES, nil
}
