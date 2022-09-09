package insumo

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mo_Find_Notify_Ended() ([]*models.Mo_Insumo_NotifyData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	/*Aca pude haber hecho un make, es decir, resultado:=make([]...)*/
	var resultado []*models.Mo_Insumo_NotifyData

	/*condicion := bson.M{
		"issendtodelete": false,
		"idcombyanf": bson.M{
			"$gt": 0,
		},
	}*/

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"issendtodelete": false}})
	condiciones = append(condiciones, bson.M{"$addFields": bson.M{"sumstock": bson.M{"$sum": bson.A{"$stock.quantity"}}}})
	condiciones = append(condiciones, bson.M{"$addFields": bson.M{"sumstocktotal": bson.M{"$add": bson.A{"$outputstock", "$sumstock"}}}})
	condiciones = append(condiciones, bson.M{"$match": bson.M{"sumstocktotal": bson.M{"$lte": 0}}})
	condiciones = append(condiciones, bson.M{"$group": bson.D{{Key: "_id", Value: "$idbusiness"}, {Key: "count", Value: bson.M{"$sum": 1}}}})
	//condiciones = append(condiciones, bson.M{"$sort": bson.M{"sum": bson.M{"$lte": 0}}})

	opciones := options.Find()
	/*Indicar como ira ordenado*/
	opciones.SetSort(bson.D{{Key: "outputstock", Value: 1}})

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Aggregate(ctx, condiciones)
	if err != nil {
		return resultado, err
	}

	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro models.Mo_Insumo_NotifyData
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)
	}

	return resultado, nil
}
