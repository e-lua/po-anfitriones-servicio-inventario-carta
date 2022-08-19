package insumo

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Find_All(idbusiness int, limit int64, offset int64) ([]*models.Mo_Insumo_Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	/*Aca pude haber hecho un make, es decir, resultado:=make([]...)*/
	var resultado []*models.Mo_Insumo_Response

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"idbusiness": idbusiness}})
	condiciones = append(condiciones, bson.M{"$match": bson.M{"issendtodelete": false}})
	condiciones = append(condiciones, bson.M{"$addFields": bson.M{"sumstock": bson.M{"$sum": bson.A{"$stock.quantity"}}}})
	condiciones = append(condiciones, bson.M{"$addFields": bson.M{"sum": bson.M{"$add": bson.A{"$outputstock", "$sumstock"}}}})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"sum": -1}})
	condiciones = append(condiciones, bson.M{"$skip": offset - 1})
	condiciones = append(condiciones, bson.M{"$limit": limit})

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Aggregate(ctx, condiciones)
	if err != nil {
		return resultado, err
	}

	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro models.Mo_Insumo_Response
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)
	}

	return resultado, nil
}
