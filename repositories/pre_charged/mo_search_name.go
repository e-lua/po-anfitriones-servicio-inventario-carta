package pre_charged

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mo_Search_Name(name string) ([]*models.Mo_Precharged_Element, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_precargada")
	col := db.Collection("elements")

	var resultado []*models.Mo_Precharged_Element

	condicion := bson.M{
		"name": primitive.Regex{
			Pattern: name,
			Options: "i",
		},
	}

	opciones := options.Find()
	/*Indicar como ira ordenado*/
	opciones.SetSort(bson.D{{Key: "name", Value: 1}})

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultado, err
	}

	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro models.Mo_Precharged_Element
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)
	}

	return resultado, nil
}
