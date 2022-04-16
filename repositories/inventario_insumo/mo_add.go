package insumo

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Mo_Add(input_insumo models.Mo_Insumo) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	_, err := col.InsertOne(ctx, input_insumo)
	if err != nil {
		return err
	}

	return nil
}
