package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_RecoverSendDelete(idelement int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB(2)

	//Actualizamos la foto de la categoría
	q := "UPDATE Element SET isexported=false,issendtodelete=false WHERE idelement=$1"
	if _, err_update := db.Exec(ctx, q, idelement); err_update != nil {
		return err_update
	}

	return nil
}
