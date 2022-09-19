package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

func Re_Set_Notified(idbusiness int, type_n string) error {

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(idbusiness)+type_n, strconv.Itoa(idbusiness), "EX", 3600)
	if err_do != nil {
		return err_do
	}

	return nil
}
