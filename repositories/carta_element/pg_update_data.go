package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Data(element models.Pg_Element, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()
	q := "UPDATE Element SET isexported=false,price=$1,description=$2,typemoney=$3, updateddate=$4,idcategory=$5,insumos=$6 FROM Category WHERE idelement=$7 AND Category.idbusiness=$8"
	if _, err_update := db.Exec(ctx, q, element.Price, element.Description, element.TypeMoney, time.Now(), element.IDCategory, element.Insumos, element.IDElement, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
