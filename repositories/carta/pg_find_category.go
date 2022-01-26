package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Category(idcarta int, idbusiness int) ([]models.Pg_Category_External, error) {

	db := models.Conectar_Pg_DB_External()

	q := "SELECT idcategory,namecategory,urlphotcategory,COUNT(idelement) FROM Element WHERE idcarta=$1 AND idbusiness=$2 GROUP BY idcategory,namecategory,urlphotcategory ORDER BY namecategory ASC"
	rows, error_shown := db.Query(context.TODO(), q, idcarta, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oLisstCategory []models.Pg_Category_External

	if error_shown != nil {

		return oLisstCategory, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oCategory models.Pg_Category_External
		rows.Scan(&oCategory.IDCategory, &oCategory.Name, &oCategory.UrlPhoto, &oCategory.AmountOfElements)
		oLisstCategory = append(oLisstCategory, oCategory)
	}

	//Si todo esta bien
	return oLisstCategory, nil
}
