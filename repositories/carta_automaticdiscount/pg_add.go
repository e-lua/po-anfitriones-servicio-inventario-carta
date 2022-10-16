package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(idbusiness int, automaticdiscount models.Pg_AutomaticDiscount) (int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	var idcategory int

	query := `INSERT INTO automaticdiscount(idbusiness,description,discount,typediscount,groupid) VALUES ($1,$2,$3,$4,$5) RETURNING`
	err := db.QueryRow(ctx, query, idbusiness, automaticdiscount.Description, automaticdiscount.Discount, automaticdiscount.TypeDiscount, automaticdiscount.Group).Scan(&idcategory)

	if err != nil {
		return idcategory, err
	}

	return idcategory, nil
}
