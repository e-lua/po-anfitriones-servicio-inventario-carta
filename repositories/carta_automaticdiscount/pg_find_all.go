package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_All(idbusiness int) ([]models.Pg_AutomaticDiscount, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "SELECT iddiscount,idbusiness,description,discount,classdiscount,typediscount,groupid FROM Category c LEFT OUTER JOIN Element e on c.idcategory=e.idcategory WHERE c.idbusiness=$1 AND c.isdeleted=false  AND c.issendtodelete=false GROUP BY c.idcategory"
	rows, error_shown := db.Query(ctx, q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListAutomatcDiscounts []models.Pg_AutomaticDiscount

	if error_shown != nil {

		return oListAutomatcDiscounts, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oDiscounts := models.Pg_AutomaticDiscount{}
		rows.Scan(&oDiscounts.IDAutomaticDiscount, &oDiscounts.IDBusiness, &oDiscounts.Description, &oDiscounts.Discount, &oDiscounts.ClassDiscount, &oDiscounts.TypeDiscount, &oDiscounts.Group)
		oListAutomatcDiscounts = append(oListAutomatcDiscounts, oDiscounts)
	}

	//Si todo esta bien
	return oListAutomatcDiscounts, nil

}
