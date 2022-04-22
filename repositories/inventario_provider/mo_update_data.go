package provider

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Update_MainData(idbusiness int, input_provider models.Mo_Providers_Response) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("provider")

	updtString := bson.M{
		"$set": bson.M{
			"type":             input_provider.Type,
			"number":           input_provider.Number,
			"email":            input_provider.Email,
			"phone":            input_provider.Phone,
			"namecontact":      input_provider.NameContact,
			"address":          input_provider.Address,
			"referenceaddress": input_provider.ReferenceAddress,
			"description":      input_provider.Description,
			"isexported":       false,
		},
	}

	filtro := bson.M{
		"idbusiness": idbusiness,
		"_id":        input_provider.ID,
	}

	_, error_update := col.UpdateOne(ctx, filtro, updtString)

	if error_update != nil {
		return error_update
	}

	return nil
}
