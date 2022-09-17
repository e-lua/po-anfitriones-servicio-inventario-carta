package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Find_Papelera(idbusiness int) ([]models.Pg_Element_Tofind, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "SELECT c.typefood,c.idcategory,COALESCE(c.urlphoto,'https://restoner-public-space.sfo3.cdn.digitaloceanspaces.com/restoner-general/default-image/default-img.png'),c.name,e.idelement,e.name,e.description,e.typemoney,e.price,COALESCE(e.urlphoto,'noimage'),e.available,e.sendtodelete,e.insumos,e.costo,e.deleteddate FROM element e JOIN category c on e.idcategory=c.idcategory WHERE c.idbusiness=$1 AND e.isdeleted=false AND e.issendtodelete=true ORDER BY e.sendtodelete DESC"
	rows, error_shown := db.Query(ctx, q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElement []models.Pg_Element_Tofind

	if error_shown != nil {

		return oListElement, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oElement := models.Pg_Element_Tofind{}
		rows.Scan(&oElement.Typefood, &oElement.IDCategory, &oElement.URLPhotoCategory, &oElement.NameCategory, &oElement.IDElement, &oElement.Name, &oElement.Description, &oElement.TypeMoney, &oElement.Price, &oElement.UrlPhoto, &oElement.Available, &oElement.SendToDelete, &oElement.Insumos, &oElement.Costo, &oElement.DeletedDate)
		oListElement = append(oListElement, oElement)
	}

	//Si todo esta bien
	return oListElement, nil
}
