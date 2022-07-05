package repositories

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Papelera(idbusiness int) ([]models.Pg_Category_Response, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB(rand.Intn(4))
	q := "SELECT c.idcategory, COUNT(e.idelement) as elements ,c.name,c.urlphoto,c.available,c.typefood,c.sendtodelete FROM Category c LEFT OUTER JOIN Element e on c.idcategory=e.idcategory WHERE c.idbusiness=$1 AND c.isdeleted=false AND c.issendtodelete=true GROUP BY c.idcategory ORDER BY c.sendtodelete DESC"
	rows, error_shown := db.Query(ctx, q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListCategory []models.Pg_Category_Response

	if error_shown != nil {

		return oListCategory, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oCategory := models.Pg_Category_Response{}
		rows.Scan(&oCategory.IDCategory, &oCategory.Elements, &oCategory.Name, &oCategory.UrlPhoto, &oCategory.Available, &oCategory.TypeFood, &oCategory.SendToDelete)
		oListCategory = append(oListCategory, oCategory)
	}

	//Si todo esta bien
	return oListCategory, nil

}
