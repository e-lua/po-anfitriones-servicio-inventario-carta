package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Address(idbusiness int) (float64, float64, error) {

	var latitude float64
	var longitude float64

	db := models.Conectar_Pg_DB()
	q := "SELECT latitude,longitude FROM idbusiness=$1"
	error_shown := db.QueryRow(context.Background(), q, idbusiness).Scan(&latitude, &longitude)

	if error_shown != nil {
		return latitude, longitude, error_shown
	}

	//Si todo esta bien
	return latitude, longitude, nil
}
