package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(element models.Pg_Element) (int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	var idelement int

	query := `INSERT INTO Element(idcategory,name,price,description,typemoney,updateddate,available,insumos) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING idelement`
	err := db.QueryRow(ctx, query, element.IDCategory, element.Name, element.Price, element.Description, element.TypeMoney, time.Now(), true, element.Insumos).Scan(&idelement)

	if err != nil {
		return idelement, err
	}

	return idelement, nil
}
