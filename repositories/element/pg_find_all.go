package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_All(idbusiness int, limit int, offset int) ([]models.Pg_Element, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT c.typefood,c.idcategory,COALESCE(c.urlphoto,'https://restoner-public-space.sfo3.cdn.digitaloceanspaces.com/restoner-general/default-image/default-img.png'),c.name,e.idelement,e.name,e.description,e.typemoney,e.price,COALESCE(e.urlphoto,'noimage'),e.available FROM element e JOIN category c on e.idcategory=c.idcategory WHERE c.idbusiness=$1 ORDER BY e.name ASC LIMIT $2 OFFSET $3"
	rows, error_shown := db.Query(context.Background(), q, idbusiness, limit, offset)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElement []models.Pg_Element

	if error_shown != nil {

		return oListElement, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oElement := models.Pg_Element{}
		rows.Scan(&oElement.Typefood, &oElement.IDCategory, &oElement.URLPhotoCategory, &oElement.NameCategory, &oElement.IDElement, &oElement.Name, &oElement.Description, &oElement.TypeMoney, &oElement.Price, &oElement.UrlPhoto, &oElement.Available)
		oListElement = append(oListElement, oElement)
	}

	//Si todo esta bien
	return oListElement, nil

}
