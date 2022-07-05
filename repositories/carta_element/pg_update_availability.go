package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Available(status bool, idelement int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "UPDATE Element SET isexported=false,available=$1,updateddate=$2 FROM Category WHERE idelement=$3 AND Category.idbusiness=$4"
	if _, err_update := db.Exec(ctx, q, status, time.Now(), idelement, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
