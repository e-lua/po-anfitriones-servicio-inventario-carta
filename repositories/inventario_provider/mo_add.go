package provider

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Mo_Add(input_provider models.Mo_Providers) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("provider")

	_, err := col.InsertOne(ctx, input_provider)
	if err != nil {
		return err
	}

	return nil
}
