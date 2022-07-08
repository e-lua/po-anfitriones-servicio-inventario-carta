package repositories

import (
	"encoding/json"
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_DataCard_Business(idbusiness int) (models.Re_SetGetCode, error) {

	var main_data_card models.Re_SetGetCode

	reply, err := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idbusiness)))

	if err != nil {
		return main_data_card, err
	}

	err = json.Unmarshal([]byte(reply), &main_data_card)

	if err != nil {
		return main_data_card, err
	}

	return main_data_card, nil
}
