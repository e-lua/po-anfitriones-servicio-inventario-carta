package insumo

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_Find_Stock(idinsumo string, idbusiness int) ([]*models.Mo_Stock, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	objID, _ := primitive.ObjectIDFromHex(idinsumo)

	/*Aca pude haber hecho un make, es decir, resultado:=make([]...)*/
	var resultado models.Mo_Insumo

	condicion := bson.M{
		"_id":        objID,
		"idbusiness": idbusiness,
		"isdeleted":  false,
	}

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return resultado.Stock, err
	}

	return resultado.Stock, nil
}
