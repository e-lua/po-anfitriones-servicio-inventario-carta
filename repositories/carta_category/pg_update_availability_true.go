package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_AvailableToTrue(idcategory int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB(2)

	q := "UPDATE Category SET available=true,isexported=false,updateddate=$1 WHERE idcategory=$2 AND idbusiness=$3"
	if _, err_update := db.Exec(ctx, q, time.Now(), idcategory, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
