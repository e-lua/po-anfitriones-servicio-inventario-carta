package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(idbusiness int, name_category string, typefood string, urlphoto string) (int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	var idcategory int

	query := `INSERT INTO Category(idbusiness,name,updateddate,available,typefood,urlphoto) VALUES ($1,$2,$3,$4,$5,$6) RETURNING idcategory`
	err := db.QueryRow(ctx, query, idbusiness, name_category, time.Now(), true, typefood, urlphoto).Scan(&idcategory)

	if err != nil {
		return idcategory, err
	}

	return idcategory, nil
}
