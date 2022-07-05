package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_SearchToNotify() ([]int, int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB(1)

	q := "SELECT idbusiness FROM schedulerange GROUP BY idbusiness"
	rows, error_shown := db.Query(ctx, q)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListBusiness []int
	quantity := 0

	if error_shown != nil {

		return oListBusiness, quantity, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oBusiness int
		rows.Scan(&oBusiness)
		oListBusiness = append(oListBusiness, oBusiness)
		quantity = quantity + 1
	}

	//Si todo esta bien
	return oListBusiness, quantity, nil

}
