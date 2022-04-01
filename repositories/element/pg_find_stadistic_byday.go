package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Stadistic_BYDay(idelement int) ([]interface{}, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()
	q := "SELECT json_build_object(extract(isodow from o.datetime),SUM(o.quantity)) FROM element AS e JOIN orders AS o ON e.idelement=o.idelement WHERE e.idelement=$1 GROUP BY extract(isodow from o.datetime)"
	rows, error_shown := db.Query(ctx, q, idelement)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListStadistic []interface{}

	if error_shown != nil {

		return oListStadistic, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oStadistic interface{}
		rows.Scan(&oStadistic)
		oListStadistic = append(oListStadistic, oStadistic)
	}

	//Si todo esta bien
	return oListStadistic, nil
}
