package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_UrlPhoto(idcategory int, urlphoto string, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	//Actualizamos la foto de la categoría
	q := "UPDATE Category SET urlphoto=$1,updateddate=$2 WHERE idcategory=$3 AND idbusiness=$4"
	if _, err_update := db.Exec(ctx, q, urlphoto, time.Now(), idcategory, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
