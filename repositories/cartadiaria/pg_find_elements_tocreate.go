package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Elements_ToCreate(date string, idbusiness int, idcategory int) ([]models.Pg_Element_ToCreate, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB_External()

	q := "SELECT e.idelement,e.idbusiness,e.idcategory,e.namecategory,e.urlphotcategory,e.name,e.price,e.description,e.urlphoto,e.typemoney,e.stock,e.typefood,e.insumos FROM element e LEFT JOIN carta c ON e.idcarta=c.idcarta WHERE c.date=$1 AND e.idbusiness=$2 AND e.idcategory=$3 ORDER BY stock ASC"
	rows, error_shown := db.Query(ctx, q, date, idbusiness, idcategory)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElementsWithStock []models.Pg_Element_ToCreate

	if error_shown != nil {

		return oListElementsWithStock, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oElement models.Pg_Element_ToCreate
		rows.Scan(&oElement.IDElement, &oElement.IDBusiness, &oElement.IDCategory, &oElement.NameCategory, &oElement.UrlPhotoCategory, &oElement.Name, &oElement.Price, &oElement.Description, &oElement.UrlPhoto, &oElement.TypeMoney, &oElement.Stock, &oElement.TypeFood, &oElement.Insumos)
		oListElementsWithStock = append(oListElementsWithStock, oElement)
	}

	//Si todo esta bien
	return oListElementsWithStock, nil
}
