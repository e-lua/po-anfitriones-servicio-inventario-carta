package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Data(discount models.Pg_AutomaticDiscount, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "UPDATE Element SET description=$1,discount=$2,classdiscount=$3,typediscount=$4,groupid=$5 WHERE iddiscount=$6 AND idbusiness=$7"
	if _, err_update := db.Exec(ctx, q, discount.Description, discount.Discount, discount.ClassDiscount, discount.TypeDiscount, discount.Group, discount.IDAutomaticDiscount, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
