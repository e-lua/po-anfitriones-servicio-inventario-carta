package repositories

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_Find_All(idbusiness int, limit int, offset int) ([]models.Pg_Element_Tofind, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()
	var db *pgxpool.Pool

	random := rand.Intn(4)
	if random%2 == 0 {
		db = models.Conectar_Pg_DB()
	} else {
		db = models.Conectar_Pg_DB_Slave()
	}

	q := "SELECT c.typefood,c.idcategory,COALESCE(c.urlphoto,'https://restoner-public-space.sfo3.cdn.digitaloceanspaces.com/restoner-general/default-image/default-img.png'),c.name,e.idelement,e.name,e.description,e.typemoney,e.price,COALESCE(e.urlphoto,'noimage'),e.available,e.insumos,e.costo,e.isautomaticcost FROM element e JOIN category c on e.idcategory=c.idcategory WHERE c.idbusiness=$1 AND e.issendtodelete=false ORDER BY e.name ASC LIMIT $2 OFFSET $3"
	rows, error_shown := db.Query(ctx, q, idbusiness, limit, offset)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElement []models.Pg_Element_Tofind

	if error_shown != nil {

		return oListElement, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oElement := models.Pg_Element_Tofind{}
		rows.Scan(&oElement.Typefood, &oElement.IDCategory, &oElement.URLPhotoCategory, &oElement.NameCategory, &oElement.IDElement, &oElement.Name, &oElement.Description, &oElement.TypeMoney, &oElement.Price, &oElement.UrlPhoto, &oElement.Available, &oElement.Insumos, &oElement.Costo, &oElement.IsAutomaticCost)
		oListElement = append(oListElement, oElement)
	}

	//Si todo esta bien
	return oListElement, nil

}
