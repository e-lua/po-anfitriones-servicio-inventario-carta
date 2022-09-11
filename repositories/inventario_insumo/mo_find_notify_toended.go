package insumo

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Find_Notify_ToEnded() ([]models.Mo_NotifyData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	var results []bson.M

	condiciones := make([]bson.D, 0)
	condiciones = append(condiciones, bson.D{{Key: "$match", Value: bson.D{{Key: "issendtodelete", Value: false}}}})
	condiciones = append(condiciones, bson.D{{Key: "$addFields", Value: bson.D{{Key: "sumstock", Value: bson.M{"$sum": bson.A{"$stock.quantity"}}}}}})
	condiciones = append(condiciones, bson.D{{Key: "$addFields", Value: bson.D{{Key: "sumstocktotal", Value: bson.M{"$add": bson.A{"$outputstock", "$sumstock"}}}}}})
	condiciones = append(condiciones, bson.D{{Key: "$match", Value: bson.D{{Key: "sumstocktotal", Value: bson.M{"$gte": 1, "$lte": 5}}}}})
	condiciones = append(condiciones, bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$idbusiness"}, {Key: "count", Value: bson.M{"$sum": 1}}}}})

	var array_notifydata []models.Mo_NotifyData

	cursor, err := col.Aggregate(ctx, condiciones)
	if err != nil {
		return array_notifydata, err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println("ERROR CURSOR", err.Error())
		return array_notifydata, err
	}

	for _, result := range results {

		var notifydata models.Mo_NotifyData

		notification := map[string]interface{}{
			"idbusiness": result["_id"],
			"quantity":   result["count"],
		}
		json_data, _ := json.Marshal(notification)

		error_decode_respuesta := json.NewDecoder(bytes.NewReader(json_data)).Decode(&notifydata)
		if error_decode_respuesta != nil {
			log.Println("ERROR DECODING: >>>", error_decode_respuesta.Error())
			return array_notifydata, error_decode_respuesta
		}

		array_notifydata = append(array_notifydata, notifydata)
	}

	return array_notifydata, nil
}
