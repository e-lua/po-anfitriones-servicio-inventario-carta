package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Data(element models.Pg_Element, idbusiness int) error {

	//Validar si hay datos de insumos
	counter_check := 0

	quantity_insumos := 0
	costo_medio_insumos := 0.0

	if len(element.Insumos) > 0 && element.IsAutomaticCost {

		for _, insumo := range element.Insumos {

			quantity_stock := 0.0
			costo := 0.0
			validate_if_have_stock := 0

			for _, stock := range insumo.Stock {
				quantity_stock += 1
				costo += stock.Price
				validate_if_have_stock += 1
			}

			quantity_insumos += 1
			costo_medio_insumos += ((costo / quantity_stock) * float64(insumo.Quantity))

			//Validar si hay insumos
			counter_check += 1
		}

		if counter_check != 0 {
			element.Costo = costo_medio_insumos
		} else {
			element.Costo = 0
		}
	}

	if len(element.Insumos) == 0 && element.IsAutomaticCost {
		element.Costo = 0
	}

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	if !element.IsURLPrecharged {
		q := "UPDATE Element SET isexported=false,price=$1,description=$2,typemoney=$3, updateddate=$4,idcategory=$5,insumos=$6,costo=$7,isautomaticcost=$8,additionals=$9 FROM Category WHERE idelement=$10 AND Category.idbusiness=$11"
		if _, err_update := db.Exec(ctx, q, element.Price, element.Description, element.TypeMoney, time.Now(), element.IDCategory, element.Insumos, element.Costo, element.IsAutomaticCost, element.Additionals, element.IDElement, idbusiness); err_update != nil {
			return err_update
		}
	} else {
		if element.UrlPhoto != "" {
			q := "UPDATE Element SET isexported=false,price=$1,description=$2,typemoney=$3, updateddate=$4,idcategory=$5,insumos=$6,costo=$7,isautomaticcost=$8,urlphoto=$9,additionals=$10 FROM Category WHERE idelement=$11 AND Category.idbusiness=$12"
			if _, err_update := db.Exec(ctx, q, element.Price, element.Description, element.TypeMoney, time.Now(), element.IDCategory, element.Insumos, element.Costo, element.IsAutomaticCost, element.UrlPhoto, element.Additionals, element.IDElement, idbusiness); err_update != nil {
				return err_update
			}
		}
	}

	return nil
}
