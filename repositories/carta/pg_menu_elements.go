package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_ElementsOfMenu_WithAction(pg_element_withaction_external []models.Pg_Element_With_Stock_External_Action, idcarta int, idbusiness int, latitude float64, longitude float64) string {

	//Variables a Insertar (Action=1)
	idelement_pg_insert, idcarta_pg_insert, idcategory_pg_insert, namecategory_pg_insert, urlphotocategory_pg_insert, name_pg_insert, price_pg_insert, description_pg_insert, urlphot_pg_insert, typem_pg_insert, stock_pg_insert, idbusiness_pg_insert, typefood_pg_insert, latitude_pg_insert, longitude_pg_insert := []int{}, []int{}, []int{}, []string{}, []string{}, []string{}, []float32{}, []string{}, []string{}, []int{}, []int{}, []int{}, []string{}, []float64{}, []float64{}

	//Variables a Eliminar (Action=3)
	idelement_pg_delete, idcarta_pg_delete := []int{}, []int{}

	//Variables a Actualizar (Action=2)
	idelement_pg_update, idcarta_pg_update, idstock_pg_update := []int{}, []int{}, []int{}

	//Repartiendo los datos
	for _, e := range pg_element_withaction_external {
		if e.Action == 1 {
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
		if e.Action == 3 {
			idelement_pg_delete = append(idelement_pg_delete, e.IDElement)
			idcarta_pg_delete = append(idcarta_pg_delete, idcarta)
		}
		if e.Action == 2 {
			idelement_pg_update = append(idelement_pg_update, e.IDElement)
			idcarta_pg_update = append(idcarta_pg_update, idcarta)
			idstock_pg_update = append(idstock_pg_update, e.Stock)
		}
	}

	/*=====================INCIAMOS TRANSACCIÓN=====================*/

	db_external := models.Conectar_Pg_DB_External()

	//INICIAMOS TRANSACCIÓN
	tx, error_tx := db_external.Begin(context.Background())
	if error_tx != nil {
		return error_tx.Error()
	}

	//ELIMINAMOS
	query_delete := `DELETE FROM element WHERE (idelement,idcarta) IN (select * from  unnest($1::int[],$2::int[]))`
	if _, err_d := tx.Exec(context.Background(), query_delete, idelement_pg_delete, idcarta_pg_delete); err_d != nil {
		tx.Rollback(context.Background())
		return "- delete " + err_d.Error()
	}

	//INSERTAMOS
	query_insert := `INSERT INTO element(idelement,idcarta,idcategory,namecategory,urlphotcategory,name,price,description,urlphoto,typemoney,stock,idbusiness,typefood) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(100)[],$5::varchar(230)[],$6::varchar(100)[],$7::decimal(8,2)[],$8::varchar(250)[],$9::varchar(230)[],$10::int[],$11::int[],$12::int[],$13::varchar(100)[],$14::real[],$15::real[]))`
	if _, err_i := tx.Exec(context.Background(), query_insert, idelement_pg_insert, idcarta_pg_insert, idcategory_pg_insert, namecategory_pg_insert, urlphotocategory_pg_insert, name_pg_insert, price_pg_insert, description_pg_insert, urlphot_pg_insert, typem_pg_insert, stock_pg_insert, idbusiness_pg_insert, typefood_pg_insert, latitude_pg_insert, longitude_pg_insert); err_i != nil {
		tx.Rollback(context.Background())
		return "- insert " + err_i.Error()
	}

	//ELIMINAMOS
	query_delete_2 := `DELETE FROM element WHERE (idelement,idcarta) IN (select * from  unnest($1::int[],$2::int[]))`
	if _, err_d_2 := tx.Exec(context.Background(), query_delete_2, idelement_pg_delete, idcarta_pg_delete); err_d_2 != nil {
		tx.Rollback(context.Background())
		return "- delete " + err_d_2.Error()
	}

	//ACTUALIZAMOS
	query_update := `UPDATE element SET stock=ex.stck FROM (select * from  unnest($1::int[],$2::int[],$3::int[])) as ex(idelem,idcart,stck) WHERE idelement=ex.idelem AND idcarta=ex.idcart`
	if _, err_u := tx.Exec(context.Background(), query_update, idelement_pg_update, idcarta_pg_update, idstock_pg_update); err_u != nil {
		tx.Rollback(context.Background())
		return "- update " + err_u.Error()
	}

	err_commit := tx.Commit(context.Background())
	if err_commit != nil {
		tx.Rollback(context.Background())
		return err_commit.Error()
	}

	/*=====================TERMINAMOS LA TRANSACCIÓN=====================*/

	return ""
}
