package provider

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Update_Delete() error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("provider")

	updtString := bson.M{
		"$set": bson.M{
			"isdeleted":  true,
			"isexported": false,
		},
	}

	filtro := bson.M{
		"issendtodelete": true,
		"deleteddate": bson.M{
			"$lt": time.Now(),
		},
	}

	_, error_update := col.UpdateMany(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
