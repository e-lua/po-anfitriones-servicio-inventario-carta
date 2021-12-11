package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_All(idbusiness int) ([]models.Pg_Category, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT idcategory,name,urlphoto,available FROM Category WHERE idbusiness=$1"
	rows, error_shown := db.Query(context.Background(), q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListCategory []models.Pg_Category

	if error_shown != nil {
		defer db.Close()
		return oListCategory, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oCategory := models.Pg_Category{}
		rows.Scan(&oCategory.IDCategory, &oCategory.Name, &oCategory.UrlPhoto, &oCategory.Available)
		oListCategory = append(oListCategory, oCategory)
	}

	defer db.Close()

	//Si todo esta bien
	return oListCategory, nil

}
