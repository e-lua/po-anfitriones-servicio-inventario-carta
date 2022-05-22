package pre_charged

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Mo_Add(input_element_precargado models.Mo_Precharged_Element) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_precargada")
	col := db.Collection("elements")

	_, err := col.InsertOne(ctx, input_element_precargado)
	if err != nil {
		return err
	}

	return nil
}
