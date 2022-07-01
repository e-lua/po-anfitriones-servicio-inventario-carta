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

	//Variables para el MQTT
	var elements_mqtt []interface{}

	//Variables a Insertar
	idelement_pg_insert, idcarta_pg_insert, idcategory_pg_insert, namecategory_pg_insert, urlphotocategory_pg_insert, name_pg_insert, price_pg_insert, description_pg_insert, urlphot_pg_insert, typem_pg_insert, stock_pg_insert, idbusiness_pg_insert, typefood_pg_insert, latitude_pg_insert, longitude_pg_insert, costo_pg_insert := []int{}, []int{}, []int{}, []string{}, []string{}, []string{}, []float32{}, []string{}, []string{}, []int{}, []int{}, []int{}, []string{}, []float64{}, []float64{}, []float64{}
	var insumos_pg_insert []interface{}

	//Repartiendo los datos
	for _, e := range pg_element_withaction_external {

		//Variables MQTT
		var one_element_mqtt models.Mqtt_Element_With_Stock
		one_element_mqtt.AvailableOrders = e.AvailableOrders
		one_element_mqtt.Date = e.Date
		one_element_mqtt.DeletedDate = time.Now().AddDate(0, 0, 5)
		one_element_mqtt.Description = e.Description
		one_element_mqtt.IDBusiness = idbusiness
		one_element_mqtt.IDCarta = idcarta
		one_element_mqtt.IDCategory = e.IDCategory
		one_element_mqtt.IDElement = e.IDElement
		one_element_mqtt.IsExported = false
		one_element_mqtt.Name = e.Name
		one_element_mqtt.NameCategory = e.NameCategory
		one_element_mqtt.Price = e.Price
		one_element_mqtt.Stock = e.Stock
		one_element_mqtt.TypeMoney = e.TypeMoney
		one_element_mqtt.Typefood = e.Typefood
		one_element_mqtt.UrlPhoto = e.UrlPhoto
		one_element_mqtt.UrlPhotoCategory = e.UrlPhotoCategory
		one_element_mqtt.Insumos = e.Insumos
		one_element_mqtt.Costo = e.Costo
		elements_mqtt = append(elements_mqtt, one_element_mqtt)

		//Variables a insertar
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
		insumos_pg_insert = append(insumos_pg_insert, e.Insumos)
		costo_pg_insert = append(costo_pg_insert, e.Costo)
	}

	db_external := models.Conectar_Pg_DB_External()

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
	query_insert := `INSERT INTO element(idelement,idcarta,idcategory,namecategory,urlphotcategory,name,price,description,urlphoto,typemoney,stock,idbusiness,typefood,latitude,longitude,insumos,costo) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(100)[],$5::varchar(230)[],$6::varchar(100)[],$7::decimal(8,2)[],$8::varchar(250)[],$9::varchar(230)[],$10::int[],$11::int[],$12::int[],$13::varchar(100)[],$14::real[],$15::real[],$16::jsonb[],$17::real[]))`
	if _, err_i := tx.Exec(ctx, query_insert, idelement_pg_insert, idcarta_pg_insert, idcategory_pg_insert, namecategory_pg_insert, urlphotocategory_pg_insert, name_pg_insert, price_pg_insert, description_pg_insert, urlphot_pg_insert, typem_pg_insert, stock_pg_insert, idbusiness_pg_insert, typefood_pg_insert, latitude_pg_insert, longitude_pg_insert, insumos_pg_insert, costo_pg_insert); err_i != nil {
		tx.Rollback(ctx)
		return err_i
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	//Serializamos el MQTT
	var serilize_elements_withstock models.Mqtt_Element_With_Stock_export
	serilize_elements_withstock.IdCarta = idcarta
	serilize_elements_withstock.Elements_with_stock = elements_mqtt

	//Comienza el proceso de MQTT
	/*ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Error(error_conection)
	}

	bytes, error_serializar := serialize(serilize_elements_withstock)
	if error_serializar != nil {
		log.Error(error_serializar)
	}

	error_publish := ch.Publish("", "anfitrion/cartadiaria_elements", false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         bytes,
		})
	if error_publish != nil {
		log.Error(error_publish)
	}*/

	return nil
}

//SERIALIZADORA
/*func serialize(serialize_elements_mqtt models.Mqtt_Element_With_Stock_export) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_elements_mqtt)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}*/
