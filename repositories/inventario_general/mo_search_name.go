package insumo

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Search_Name(idbusiness int, name string) ([]*models.Mo_Insumo_Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	var resultado []*models.Mo_Insumo_Response

	pipeline := []bson.M{
		{"$match": bson.M{"idbusiness": idbusiness}},
		{"$match": bson.M{"isdeleted": false}},
		{"$match": bson.M{
			"name":           "/" + name + "/",
			"$caseSensitive": true,
		},
		},
		{"$sort": bson.M{"tweet.fecha": -1}},
	}

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Aggregate(ctx, pipeline)
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
