package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_ByCategory(idbusiness int, idcategory int) ([]models.Pg_ElementsByCategory, int, error) {

	quantity := 0

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "SELECT e.name,e.available FROM Element as e JOIN Category as c ON e.idcategory=e.idcategory WHERE c.idbusiness=$1 AND c.idcategory=$2 AND e.isdeleted=false  AND e.issendtodelete=false"
	rows, error_shown := db.Query(ctx, q, idbusiness, idcategory)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElementByCategory []models.Pg_ElementsByCategory

	if error_shown != nil {
		return oListElementByCategory, 0, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		quantity = quantity + 1
		oElement_by_Category := models.Pg_ElementsByCategory{}
		rows.Scan(&oElement_by_Category.Element, &oElement_by_Category.Available)
		oListElementByCategory = append(oListElementByCategory, oElement_by_Category)
	}

	//Si todo esta bien
	return oListElementByCategory, quantity, nil

}
