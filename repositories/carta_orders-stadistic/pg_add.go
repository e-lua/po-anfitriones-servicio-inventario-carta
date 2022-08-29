package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Insert_OrderStadistic(orderstadistic []models.Pg_Import_StadisticOrders) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Instanciando los valores
	id_pg, idelement_pg, quantity_pg, datetime_pg, totalamount_pg, totalcost_pg := []int64{}, []int{}, []int{}, []string{}, []float32{}, []float32{}

	//counter
	counter := 0
	for _, os := range orderstadistic {
		id_pg = append(id_pg, time.Now().UnixMilli()+int64(counter))
		idelement_pg = append(idelement_pg, os.IdElement)
		totalamount_pg = append(totalamount_pg, os.TotalAmount)
		totalcost_pg = append(totalcost_pg, os.TotalCost)
		quantity_pg = append(quantity_pg, os.Quantity)
		datetime_pg = append(datetime_pg, os.Datetime)
		counter = counter + 1
	}

	//Enviado los datos a la base de datos
	db := models.Conectar_Pg_DB()
	query := `INSERT INTO orders(id,idelement,quantity,datetime,totalamount,totalcost) (select * from unnest($1::bigint[], $2::int[],$3::int[],$4::varchar(30)[],$5::decimal(10,2)[],$6::decimal(10,2)[]))`
	if _, err := db.Exec(ctx, query, id_pg, idelement_pg, quantity_pg, datetime_pg, totalamount_pg, totalcost_pg); err != nil {
		return err
	}

	return nil
}
