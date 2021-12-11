package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Data(element models.Pg_Element, idbusiness int) error {

	db := models.Conectar_Pg_DB()
	q := "UPDATE Element SET name=$1,price=$2,description=$3,typemoney=$4,updateddate=$5 FROM Category WHERE idelement=$6 AND Category.idbusiness=$7"
	if _, err_update := db.Exec(context.Background(), q, element.Name, element.Price, element.Description, element.TypeMoney, time.Now(), element.IDElement, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
