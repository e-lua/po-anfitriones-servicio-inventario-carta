package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(idbusiness int, input_date string) error {

	db := models.Conectar_Pg_DB()

	query := `INSERT INTO Menu(idbusiness,date) VALUES ($1,$2)`
	if _, err := db.Exec(context.Background(), query, idbusiness, input_date); err != nil {
		return err
	}

	return nil
}
