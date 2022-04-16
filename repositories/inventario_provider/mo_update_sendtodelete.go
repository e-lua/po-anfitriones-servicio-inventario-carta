package provider

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_Update_SendToDelete(idprovider string, timezone int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("provider")

	updtString := bson.M{
		"$set": bson.M{
			"issendtodelete": true,
			"sendtodelete":   time.Now().Add(time.Hour * time.Duration(timezone)),
			"deleteddate":    time.Now().AddDate(0, 0, 7),
			"isexported":     false,
		},
	}

	objID, _ := primitive.ObjectIDFromHex(idprovider)

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
