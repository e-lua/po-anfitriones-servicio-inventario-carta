package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(idbusiness int, automaticdiscount models.Pg_AutomaticDiscount) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `INSERT INTO automaticdiscount(idbusiness,description,discount,classdiscount,typediscount,groupid) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := db.Query(ctx, query, idbusiness, automaticdiscount.Description, automaticdiscount.Discount, automaticdiscount.ClassDiscount, automaticdiscount.TypeDiscount, automaticdiscount.Group)

	if err != nil {
		return err
	}

	return nil
}
