package insumo

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Find_Notify_ToEnded() ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	var resultado []interface{}

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"issendtodelete": false}})
	condiciones = append(condiciones, bson.M{"$addFields": bson.M{"sumstock": bson.M{"$sum": bson.A{"$stock.quantity"}}}})
	condiciones = append(condiciones, bson.M{"$addFields": bson.M{"sumstocktotal": bson.M{"$add": bson.A{"$outputstock", "$sumstock"}}}})
	condiciones = append(condiciones, bson.M{"$match": bson.M{"sumstocktotal": bson.M{"$gte": 1, "$lte": 5}}})
	condiciones = append(condiciones, bson.M{"$group": bson.M{"_id": "$idbusiness", "count": bson.M{"$sum": 1}}})

	cursor, err := col.Aggregate(ctx, condiciones)
	if err != nil {
		return resultado, err
	}

	for cursor.Next(context.TODO()) {
		var registro interface{}
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		resultado = append(resultado, &registro)
	}

	return resultado, nil
}
