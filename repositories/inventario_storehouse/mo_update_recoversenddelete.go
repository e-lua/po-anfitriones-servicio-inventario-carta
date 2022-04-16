package storehouse

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_Update_RecoverSendDelete(idstorehouse string, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("storehouse")

	updtString := bson.M{
		"$set": bson.M{
			"issendtodelete": false,
			"isexported":     false,
		},
	}

	objID, _ := primitive.ObjectIDFromHex(idstorehouse)

	filtro := bson.M{
		"idbusiness": idbusiness,
		"_id":        objID,
	}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
