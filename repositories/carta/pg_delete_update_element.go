package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Delete_Update_Element(pg_element_withaction_external []models.Pg_Element_With_Stock_External, idcarta int, idbusiness int, latitude float64, longitude float64) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db_external := models.Conectar_Pg_DB_External()

	//Variables a Insertar (Action=1)
	idelement_pg_insert, idcarta_pg_insert, idcategory_pg_insert, namecategory_pg_insert, urlphotocategory_pg_insert, name_pg_insert, price_pg_insert, description_pg_insert, urlphot_pg_insert, typem_pg_insert, stock_pg_insert, idbusiness_pg_insert, typefood_pg_insert, latitude_pg_insert, longitude_pg_insert := []int{}, []int{}, []int{}, []string{}, []string{}, []string{}, []float32{}, []string{}, []string{}, []int{}, []int{}, []int{}, []string{}, []float64{}, []float64{}

	//Repartiendo los datos
	for _, e := range pg_element_withaction_external {
		idelement_pg_insert = append(idelement_pg_insert, e.IDElement)
		idcarta_pg_insert = append(idcarta_pg_insert, idcarta)
		idcategory_pg_insert = append(idcategory_pg_insert, e.IDCategory)
		namecategory_pg_insert = append(namecategory_pg_insert, e.NameCategory)
		urlphotocategory_pg_insert = append(urlphotocategory_pg_insert, e.UrlPhotoCategory)
		name_pg_insert = append(name_pg_insert, e.Name)
		price_pg_insert = append(price_pg_insert, e.Price)
		description_pg_insert = append(description_pg_insert, e.Description)
		urlphot_pg_insert = append(urlphot_pg_insert, e.UrlPhoto)
		typem_pg_insert = append(typem_pg_insert, e.TypeMoney)
		stock_pg_insert = append(stock_pg_insert, e.Stock)
		idbusiness_pg_insert = append(idbusiness_pg_insert, idbusiness)
		typefood_pg_insert = append(typefood_pg_insert, e.Typefood)
		latitude_pg_insert = append(latitude_pg_insert, latitude)
		longitude_pg_insert = append(longitude_pg_insert, longitude)
	}

	//BEGIN
	tx, error_tx := db_external.Begin(ctx)
	if error_tx != nil {
		return error_tx
	}

	//ELIMINAR LOS ELEMENTOS
	q_delete_list := `DELETE FROM Element WHERE idbusiness=$1 AND idcarta=$2`
	if _, err_update := tx.Exec(ctx, q_delete_list, idbusiness, idcarta); err_update != nil {
		tx.Rollback(ctx)
		return err_update
	}

	//INSERTAMOS LOS ELEMENTOS
	query_insert := `INSERT INTO element(idelement,idcarta,idcategory,namecategory,urlphotcategory,name,price,description,urlphoto,typemoney,stock,idbusiness,typefood,latitude,longitude) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(100)[],$5::varchar(230)[],$6::varchar(100)[],$7::decimal(8,2)[],$8::varchar(250)[],$9::varchar(230)[],$10::int[],$11::int[],$12::int[],$13::varchar(100)[],$14::real[],$15::real[]))`
	if _, err_i := tx.Exec(ctx, query_insert, idelement_pg_insert, idcarta_pg_insert, idcategory_pg_insert, namecategory_pg_insert, urlphotocategory_pg_insert, name_pg_insert, price_pg_insert, description_pg_insert, urlphot_pg_insert, typem_pg_insert, stock_pg_insert, idbusiness_pg_insert, typefood_pg_insert, latitude_pg_insert, longitude_pg_insert); err_i != nil {
		tx.Rollback(ctx)
		return err_i
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	return nil
}
