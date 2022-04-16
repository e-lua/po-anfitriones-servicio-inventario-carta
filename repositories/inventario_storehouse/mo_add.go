package storehouse

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Mo_Add(input_storehouse models.Mo_StoreHouse) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("storehouse")

	_, err := col.InsertOne(ctx, input_storehouse)
	if err != nil {
		return err
	}

	return nil
}
