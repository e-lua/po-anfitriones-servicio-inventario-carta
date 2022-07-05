package export_to_file

import (
	"bytes"
	"context"
	"encoding/json"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Pg_Elements_ToFile(element_data models.Mqtt_Request_Element) error {

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

	q := "SELECT c.typefood,c.idcategory,COALESCE(c.urlphoto,'https://restoner-public-space.sfo3.cdn.digitaloceanspaces.com/restoner-general/default-image/default-img.png'),c.name,e.idelement,e.name,e.description,e.typemoney,e.price,COALESCE(e.urlphoto,'noimage'),e.available,e.insumos,e.costo FROM element e JOIN category c on e.idcategory=c.idcategory WHERE c.idbusiness=$1 AND c.isdeleted=false AND c.issendtodelete=false ORDER BY e.name ASC"
	rows, error_shown := db.Query(ctx, q, element_data.IDBusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElement []models.Pg_Element_Tofind

	if error_shown != nil {

		return error_shown
	}

	quantity := 0

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oElement := models.Pg_Element_Tofind{}
		rows.Scan(&oElement.Typefood, &oElement.IDCategory, &oElement.URLPhotoCategory, &oElement.NameCategory, &oElement.IDElement, &oElement.Name, &oElement.Description, &oElement.TypeMoney, &oElement.Price, &oElement.UrlPhoto, &oElement.Available, &oElement.Insumos, &oElement.Costo)
		oListElement = append(oListElement, oElement)

		if oElement.IDCategory > 0 {
			quantity += 1
		}
	}

	element_data.Elements = oListElement

	if quantity > 0 {

		/*----------------------------MQTT----------------------------*/

		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes_element, error_serializar_ele := serialize_elements(element_data)
		if error_serializar_ele != nil {
			log.Error(error_serializar_ele)
		}

		error_publish_2 := ch.Publish("", "anfitrion/element_to_file", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes_element,
			})
		if error_publish_2 != nil {
			log.Error(error_publish_2)
		}

	}

	//Si todo esta bien
	return nil

}

//SERIALIZADORA ELEMENT
func serialize_elements(element_data models.Mqtt_Request_Element) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(element_data)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
