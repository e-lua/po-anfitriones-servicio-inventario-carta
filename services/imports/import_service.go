package imports

import (
	//REPOSITORIES

	"github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	imports_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/imports"
)

func UpdateInsumoStock_Service(input_insumos []models.Mqtt_Import_InsumoStock) (int, bool, string, string) {

	error_update := imports_repository.Mo_Update_Many(input_insumos)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el stock" + error_update.Error(), ""
	}

	return 200, false, "", "Actualización correcta"
}
