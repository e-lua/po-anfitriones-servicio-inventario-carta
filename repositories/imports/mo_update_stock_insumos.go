package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mo_Update_Many(input_elements []models.Mqtt_Import_InsumoStock) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	models := []mongo.WriteModel{}

	for _, element_extracted := range input_elements {

		for _, insumo := range element_extracted.Insumos {

			updtString := bson.M{
				"$inc": bson.M{
					"outputstock": -(element_extracted.Quantity * insumo.Quantity),
				},
			}

			models = append(models,
				mongo.NewUpdateOneModel().SetFilter(
					bson.M{
						"_id": insumo.ID,
					},
				).
					SetUpdate(updtString).SetUpsert(true),
			)

		}

	}

	opts := options.BulkWrite().SetOrdered(true)

	_, error_update := col.BulkWrite(ctx, models, opts)

	if error_update != nil {
		return error_update
	}

	return nil
}
