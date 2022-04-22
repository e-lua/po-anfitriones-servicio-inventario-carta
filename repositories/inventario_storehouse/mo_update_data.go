package storehouse

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Update_MainData(idbusiness int, input_storehouse models.Mo_StoreHouse_Response) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("storehouse")

	updtString := bson.M{
		"$set": bson.M{
			"description": input_storehouse.Description,
			"isexported":  false,
		},
	}

	filtro := bson.M{
		"idbusiness": idbusiness,
		"_id":        input_storehouse.ID,
	}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
