package repositories

import (
	"context"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Pg_Update_Elements(pg_element_external []models.Pg_Element_With_Stock_External, idcarta int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB_External()

	idelement_pg, idcarta_pg, idcategory_pg, namecategory_pg, urlphotocategory_pg, name_pg, price_pg, description_pg, urlphot_pg, typem_pg, stock_pg, idbusiness_pg := []int{}, []int{}, []int{}, []string{}, []interface{}{}, []string{}, []float32{}, []string{}, []interface{}{}, []int{}, []int{}, []int{}

	for _, e := range pg_element_external {
		idelement_pg = append(idelement_pg, e.IDElement)
		idcarta_pg = append(idcarta_pg, e.IDCarta)
		idcategory_pg = append(idcategory_pg, e.IDCategory)
		namecategory_pg = append(namecategory_pg, e.NameCategory)
		urlphotocategory_pg = append(urlphotocategory_pg, e.UrlPhotoCategory)
		name_pg = append(name_pg, e.Name)
		price_pg = append(price_pg, e.Price)
		description_pg = append(description_pg, e.Description)
		urlphot_pg = append(urlphot_pg, e.UrlPhoto)
		typem_pg = append(typem_pg, e.TypeMoney)
		stock_pg = append(stock_pg, e.Stock)
		idbusiness_pg = append(idbusiness_pg, idbusiness)
	}

	q := `INSERT INTO element(idelement,idcarta,idcategory,namecategory,urlphotcategory,name,price,description,urlphoto,typemoney,stock,idbusiness) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(100)[],$5::interface[],$6::varchar(100)[],$7::decimal(8,2)[],$8::varchar(250)[],$9::varchar(230)[],$10::int[],$11::int[],$12::int[]));`
	if _, err_update := db_external.Exec(context.Background(), q, idelement_pg, idcarta_pg, idcategory_pg, namecategory_pg, urlphotocategory_pg, name_pg, price_pg, description_pg, urlphot_pg, typem_pg, stock_pg, idbusiness_pg); err_update != nil {
		return err_update
	}

	return nil
}
