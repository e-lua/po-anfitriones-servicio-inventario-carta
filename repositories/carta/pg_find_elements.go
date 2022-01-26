package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Elements(idcarta int, idbusiness int) ([]models.Pg_Element_With_Stock_External, error) {

	db := models.Conectar_Pg_DB_External()

	q := "SELECT idelement,idcarta,idbusiness,idcategory,namecategory,urlphotcategory,name,price,description,urlphoto,typemoney,stock,typefood FROM element WHERE idcarta=$1 AND idbusiness=$2 ORDER BY stock ASC"
	rows, error_shown := db.Query(context.TODO(), q, idcarta, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElementsWithStock []models.Pg_Element_With_Stock_External

	if error_shown != nil {

		return oListElementsWithStock, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oElement models.Pg_Element_With_Stock_External
		rows.Scan(&oElement.IDElement, &oElement.IDCarta, &oElement.IDBusiness, &oElement.IDCategory, &oElement.NameCategory, &oElement.UrlPhotoCategory, &oElement.Name, &oElement.Price, &oElement.Description, &oElement.UrlPhoto, &oElement.TypeMoney, &oElement.Stock, &oElement.Typefood)
		oListElementsWithStock = append(oListElementsWithStock, oElement)
	}

	//Si todo esta bien
	return oListElementsWithStock, nil
}
