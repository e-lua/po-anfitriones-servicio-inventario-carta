package insumo

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mo_Update_Stock(idinsumo string, idbusiness int, input_insumo models.Mo_Insumo_Stock_Adjust_Requst) error {

	var stock_list []models.Mo_Stock

	if input_insumo.IsAdjust {

		monto := input_insumo.AmountToAdjust * -1
		counter := 0

		longitud := len(input_insumo.Stock)

		for _, stock := range input_insumo.Stock {

			var stock_one models.Mo_Stock
			//var resultado_monto int

			stock_one.Price = stock.Price
			stock_one.IdProvider = stock.IdProvider
			stock_one.TimeZone = stock.TimeZone
			stock_one.CreatedDate = stock.CreatedDate

			if monto > 0 {

				if counter == longitud {
					stock_one.Quantity = stock.Quantity - monto
					break
				} else {
					monto = monto - stock.Quantity
				}

				if monto >= 0 {
					stock_one.Quantity = 0
				} else {
					stock_one.Quantity = monto
				}

			} else {
				stock_one.Quantity = stock.Quantity
			}

			stock_one.ProviderName = stock.ProviderName

			stock_list = append(stock_list, stock_one)

			counter += 1
		}
	} else {
		stock_list = input_insumo.Stock
	}

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	updtString := bson.M{
		"$set": bson.M{
			"stock":      stock_list,
			"isexported": false,
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
