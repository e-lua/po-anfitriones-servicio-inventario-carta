package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(element models.Pg_Element) (int, error) {

	db := models.Conectar_Pg_DB()

	var idelement int

	query := `INSERT INTO Element(idcategory,name,price,description,typemoney,updateddate,available) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	err := db.QueryRow(context.Background(), query, element.IDCategory, element.Name, element.Price, element.Description, element.TypeMoney, time.Now(), true).Scan(&idelement)

	if err != nil {
		return idelement, err
	}

	return idelement, nil
}
