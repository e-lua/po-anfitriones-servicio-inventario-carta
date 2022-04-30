package imports

import (
	//REPOSITORIES

	"github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	imports_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/imports"
)

func UpdateInsumoStock_Service(input_insumos []models.Mqtt_Import_InsumoStock) error {

	error_update := imports_repository.Mo_Update_Many(input_insumos)
	if error_update != nil {
		return error_update
	}

	return nil
}
