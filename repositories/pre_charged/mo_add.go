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

func Mo_Add_Multiple(input_element_precargado []models.Mo_Precharged_Element) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_precargada")
	col := db.Collection("elements")

	var elements_precargados []interface{}

	for _, element_precargado := range input_element_precargado {
		var precharged models.Mo_Precharged_Element
		precharged.Name = element_precargado.Name
		precharged.URL = element_precargado.URL
		elements_precargados = append(elements_precargados, precharged)
	}

	_, err := col.InsertMany(ctx, elements_precargados)
	if err != nil {
		return err
	}

	return nil
}
