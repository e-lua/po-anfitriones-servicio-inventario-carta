package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Add(element models.Pg_Element) (int, error) {

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

	var idelement int

	query := `INSERT INTO Element(idcategory,name,price,description,typemoney,updateddate,available,insumos,costo,urlphoto,additionals) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING idelement`
	err := db.QueryRow(ctx, query, element.IDCategory, element.Name, element.Price, element.Description, element.TypeMoney, time.Now(), true, element.Insumos, element.Costo, element.UrlPhoto, element.Additionals).Scan(&idelement)

	if err != nil {
		return idelement, err
	}

	return idelement, nil
}
