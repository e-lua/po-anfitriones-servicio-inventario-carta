package insumo

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_Update_MainData(idbusiness int, idinsumo string, input_insumo models.Mo_Insumo) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	updtString := bson.M{
		"$set": bson.M{
			"measure":        input_insumo.Measure,
			"idstorehouse":   input_insumo.IDStoreHouse,
			"namestorehouse": input_insumo.NameStoreHouse,
			"description":    input_insumo.Description,
			"isexported":     false,
		},
	}

	objID, _ := primitive.ObjectIDFromHex(idinsumo)

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
