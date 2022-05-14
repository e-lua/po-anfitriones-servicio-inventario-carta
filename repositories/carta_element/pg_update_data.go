package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Data(element models.Pg_Element, idbusiness int) error {

	var quantity_insumos int
	var costo_medio_insumos float64

	if len(element.Insumos) > 0 {

		for _, insumo := range element.Insumos {
			var quantity_stock int
			var costo float64
			for _, stock := range insumo.Stock {
				quantity_stock += 1
				costo += stock.Price
			}
			quantity_insumos += 1
			costo_medio_insumos += ((costo / float64(quantity_stock)) * float64(insumo.Quantity))
		}
		element.Costo = (costo_medio_insumos / float64(quantity_insumos))

	}
	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()
	q := "UPDATE Element SET isexported=false,price=$1,description=$2,typemoney=$3, updateddate=$4,idcategory=$5,insumos=$6,costo=$7 FROM Category WHERE idelement=$8 AND Category.idbusiness=$9"
	if _, err_update := db.Exec(ctx, q, element.Price, element.Description, element.TypeMoney, time.Now(), element.IDCategory, element.Insumos, element.Costo, element.IDElement, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
