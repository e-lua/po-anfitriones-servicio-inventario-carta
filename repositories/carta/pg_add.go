package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(idbusiness int, date string) (int, error) {

	var idcarta int

	db_external := models.Conectar_Pg_DB_External()

	query := `INSERT INTO Carta(idbusiness,date,updateddate) VALUES ($1,$2,$3) RETURNING idcarta`
	err := db_external.QueryRow(context.TODO(), query, idbusiness, date, time.Now()).Scan(&idcarta)

	if err != nil {
		return idcarta, err
	}

	return idcarta, nil
}
