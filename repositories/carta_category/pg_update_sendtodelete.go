package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_SendToDelete(idcategory int, timezone int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	//Actualizamos la foto de la categoría
	q := "UPDATE Category SET isexported=false,issendtodelete=true,sendtodelete=$1,deleteddate=$2 WHERE idcategory=$3 AND idbusiness=$4"
	if _, err_update := db.Exec(ctx, q, time.Now().Add(time.Hour*time.Duration(timezone)), time.Now().AddDate(0, 0, 7), idcategory, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
