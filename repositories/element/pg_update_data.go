package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Data(element models.Pg_Element, idbusiness int) error {

	db := models.Conectar_Pg_DB()
	q := "UPDATE Element SET price=$1,description=$2,typemoney=$3, updateddate=$4 FROM Category WHERE idelement=$5 AND Category.idbusiness=$6"
	if _, err_update := db.Exec(context.Background(), q, element.Price, element.Description, element.TypeMoney, time.Now(), element.IDElement, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
