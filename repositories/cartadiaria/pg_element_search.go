package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Elements_SearchByText(date string, idbusiness int, text string, limit int, offset int) ([]models.Pg_Element_With_Stock_External, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB_External()

	q := "SELECT e.idelement,e.idbusiness,e.idcategory,e.namecategory,e.urlphotcategory,e.name,e.price,e.description,e.urlphoto,e.typemoney,e.stock,e.typefood,e.insumos,e.availableorders FROM element e LEFT JOIN carta c ON e.idcarta=c.idcarta WHERE c.date=$1 AND e.idbusiness=$2 AND (lower(e.name) ~ $3 OR lower(e.description) ~ $3) ORDER BY e.name ASC LIMIT $4 OFFSET $5"
	rows, error_shown := db.Query(ctx, q, date, idbusiness, text, limit, offset)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElementsWithStock []models.Pg_Element_With_Stock_External

	if error_shown != nil {

		return oListElementsWithStock, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oElement models.Pg_Element_With_Stock_External
		rows.Scan(&oElement.IDElement, &oElement.IDBusiness, &oElement.IDCategory, &oElement.NameCategory, &oElement.UrlPhotoCategory, &oElement.Name, &oElement.Price, &oElement.Description, &oElement.UrlPhoto, &oElement.TypeMoney, &oElement.Stock, &oElement.Typefood, &oElement.Insumos, &oElement.AvailableOrders)
		oListElementsWithStock = append(oListElementsWithStock, oElement)
	}

	//Si todo esta bien
	return oListElementsWithStock, nil
}
