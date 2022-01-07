package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(idbusiness int, name_category string, typefood string) (int, error) {

	db := models.Conectar_Pg_DB()

	var idcategory int

	query := `INSERT INTO Category(idbusiness,name,updateddate,available,typefood) VALUES ($1,$2,$3,$4,$5) RETURNING idcategory`
	err := db.QueryRow(context.Background(), query, idbusiness, name_category, time.Now(), true, typefood).Scan(&idcategory)

	if err != nil {
		return idcategory, err
	}

	return idcategory, nil
}
