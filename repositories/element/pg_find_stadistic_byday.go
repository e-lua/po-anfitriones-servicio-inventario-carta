package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Stadistic_BYDay(idelement int) (models.Pg_StadisticByElement, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var stadistic_byelement models.Pg_StadisticByElement

	db := models.Conectar_Pg_DB()
	q := "SELECT json_build_object(extract(isodow from o.datetime),SUM(o.quantity)) FROM element AS e JOIN orders AS o ON e.idelement=o.idelement WHERE e.idelement=$1 GROUP BY extract(isodow from o.datetime)"
	error_shown := db.QueryRow(ctx, q, idelement).Scan(&stadistic_byelement)

	if error_shown != nil {

		return stadistic_byelement, error_shown
	}

	//Si todo esta bien
	return stadistic_byelement, nil
}
