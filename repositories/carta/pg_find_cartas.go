package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Cartas(idbusiness int) ([]models.Pg_Carta_Found, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB_External()

	q := "SELECT c.date::date,COUNT(e.idelement) FROM carta AS c JOIN element AS e ON c.idcarta=e.idcarta WHERE c.idbusiness=$1 AND c.date>=now()::date-INTERVAL '1 DAY' GROUP BY c.date::date"
	rows, error_shown := db.Query(ctx, q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListCarta []models.Pg_Carta_Found

	if error_shown != nil {

		return oListCarta, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oCarta models.Pg_Carta_Found
		rows.Scan(&oCarta.Date, &oCarta.Elements)
		oListCarta = append(oListCarta, oCarta)
	}

	//Si todo esta bien
	return oListCarta, nil
}
