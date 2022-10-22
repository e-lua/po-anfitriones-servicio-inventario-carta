package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mo_Update_Many(input_insumos []models.Mqtt_Import_InsumoStock) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("insumo")

	models := []mongo.WriteModel{}

	for _, insumo := range input_insumos {

		updtString := bson.M{
			"$inc": bson.M{
				"outputstock": -(insumo.Quantity),
			},
		}

		objID, _ := primitive.ObjectIDFromHex(insumo.Insumos)

		models = append(models,
			mongo.NewUpdateOneModel().SetFilter(
				bson.M{
					"_id": objID,
				},
			).
				SetUpdate(updtString).SetUpsert(true),
		)

	}

	opts := options.BulkWrite().SetOrdered(true)

	_, error_update := col.BulkWrite(ctx, models, opts)

	if error_update != nil {
		return error_update
	}

	return nil
}
