package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_ByDayWeek(idbusiness int, dayofweek int, limit int, offset int) ([]models.Pg_Element_WithRating, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()
	q := "SELECT c.typefood,c.idcategory,COALESCE(c.urlphoto,'https://restoner-public-space.sfo3.cdn.digitaloceanspaces.com/restoner-general/default-image/default-img.png'),c.name,e.idelement,e.name,e.description,e.typemoney,e.price,COALESCE(e.urlphoto,'noimage'),e.available,SUM(ord.quantity) FROM element e JOIN category c on e.idcategory=c.idcategory JOIN orders AS ord ON e.idelement=ord.idelement WHERE c.idbusiness=$1 AND extract(isodow from ord.datetime)=$2 GROUP BY c.typefood,c.idcategory,COALESCE(c.urlphoto,'https://restoner-public-space.sfo3.cdn.digitaloceanspaces.com/restoner-general/default-image/default-img.png'),c.name,e.idelement,e.name,e.description,e.typemoney,e.price,COALESCE(e.urlphoto,'noimage'),e.available ORDER BY SUM(ord.quantity) DESC LIMIT $3 OFFSET $4"
	rows, error_shown := db.Query(ctx, q, idbusiness, dayofweek, limit, offset)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElement []models.Pg_Element_WithRating

	if error_shown != nil {

		return oListElement, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oElement := models.Pg_Element_WithRating{}
		rows.Scan(&oElement.Typefood, &oElement.IDCategory, &oElement.URLPhotoCategory, &oElement.NameCategory, &oElement.IDElement, &oElement.Name, &oElement.Description, &oElement.TypeMoney, &oElement.Price, &oElement.UrlPhoto, &oElement.Available, &oElement.Orders)
		oListElement = append(oListElement, oElement)
	}

	//Si todo esta bien
	return oListElement, nil
}
