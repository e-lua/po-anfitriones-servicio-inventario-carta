package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_AvailableToFalse(idcategory int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "UPDATE Category set available=false,isexported=false,updateddate=$1 WHERE idcategory=$2 AND (SELECT COUNT(e.idelement) FROM Category c LEFT OUTER JOIN Element e on c.idcategory=e.idcategory WHERE c.idbusiness=$3 AND c.idcategory=$4 AND e.isdeleted=false GROUP BY c.idcategory)=0"
	if _, err_update := db.Exec(ctx, q, time.Now(), idcategory, idbusiness, idcategory); err_update != nil {
		return err_update
	}

	return nil
}
