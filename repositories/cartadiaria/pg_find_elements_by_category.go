package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Elements_ByCategory(idcarta int, idbusiness int, idcategory int) ([]models.Pg_Element_With_Stock_External, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB_External()

	q := "SELECT idelement,idcarta,idbusiness,idcategory,namecategory,urlphotcategory,name,price,description,urlphoto,typemoney,stock,typefood,insumos,availableorders,costo FROM element WHERE idcarta=$1 AND idbusiness=$2 AND idcategory=$3 ORDER BY stock ASC"
	rows, error_shown := db.Query(ctx, q, idcarta, idbusiness, idcategory)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElementsWithStock []models.Pg_Element_With_Stock_External

	if error_shown != nil {

		return oListElementsWithStock, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oElement models.Pg_Element_With_Stock_External
		rows.Scan(&oElement.IDElement, &oElement.IDCarta, &oElement.IDBusiness, &oElement.IDCategory, &oElement.NameCategory, &oElement.UrlPhotoCategory, &oElement.Name, &oElement.Price, &oElement.Description, &oElement.UrlPhoto, &oElement.TypeMoney, &oElement.Stock, &oElement.Typefood, &oElement.Insumos, &oElement.AvailableOrders, &oElement.Costo)
		oListElementsWithStock = append(oListElementsWithStock, oElement)
	}

	//Si todo esta bien
	return oListElementsWithStock, nil
}
