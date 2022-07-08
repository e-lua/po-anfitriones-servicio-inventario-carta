package repositories

import (
	"encoding/json"
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Re_Set_DataCard_Business(idbusiness int, cartaData models.Pg_Category_Element_ScheduleRange) error {

	var main_data_card models.Re_SetGetCode
	main_data_card.CartaData = cartaData
	main_data_card.IdBusiness = idbusiness

	uJson, err_marshal := json.Marshal(main_data_card.CartaData)
	if err_marshal != nil {
		return err_marshal
	}

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(main_data_card.IdBusiness), uJson, "EX", 300)
	if err_do != nil {
		return err_do
	}

	return nil
}
