package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(idbusiness int, date string) (int, error) {
	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var idcarta int

	db_external := models.Conectar_Pg_DB_External()

	query := `INSERT INTO Carta(idbusiness,date,updateddate,deleteddate) VALUES ($1,$2,$3,$4) RETURNING idcarta`
	err := db_external.QueryRow(ctx, query, idbusiness, date, time.Now(), time.Now().AddDate(0, 0, 6)).Scan(&idcarta)

	if err != nil {
		return idcarta, err
	}

	return idcarta, nil
}
