package export_to_file

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mo_Insumos_ToFile(idbusiness int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	/*Aca pude haber hecho un make, es decir, resultado:=make([]...)*/
	var resultado []*models.Mo_Insumo_Response

	condicion := bson.M{
		"idbusiness": idbusiness,
		"isdeleted":  false,
	}

	opciones := options.Find()
	/*Indicar como ira ordenado*/
	opciones.SetSort(bson.D{{Key: "name", Value: 1}})

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return err
	}

	quantity := 0

	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro models.Mo_Insumo_Response
		err := cursor.Decode(&registro)
		if err != nil {
			return err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)

		if len(registro.Measure) > 0 {
			quantity += 1
		}
	}

	if quantity > 0 {

		/*----------------------------MQTT----------------------------*/

		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes_element, error_serializar_ele := serialize_insumos(resultado)
		if error_serializar_ele != nil {
			log.Error(error_serializar_ele)
		}

		error_publish_2 := ch.Publish("", "anfitrion/insumo_to_file", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes_element,
			})
		if error_publish_2 != nil {
			log.Error(error_publish_2)
		}

	}

	return nil
}

//SERIALIZADORA INSUMO
func serialize_insumos(serialize_insumo []*models.Mo_Insumo_Response) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_insumo)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
