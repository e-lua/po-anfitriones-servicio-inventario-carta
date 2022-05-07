package exportfile

import (
	//REPOSITORIES

	"github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	export_file "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/export_to_file"
)

func ExportFile_Insumo_Service(insumo_data models.Mqtt_Request_Insumo) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_export := export_file.Mo_Insumos_ToFile(insumo_data)
	if error_export != nil {
		return 500, true, "Error en el servidor interno al intentar exportar los insumos, detalles: " + error_export.Error(), ""
	}

	return 201, false, "", "Enviado correctamente"
}

func ExportFile_Element_Service(element_data models.Mqtt_Request_Element) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_export := export_file.Pg_Elements_ToFile(element_data)
	if error_export != nil {
		return 500, true, "Error en el servidor interno al intentar exportar los elementos, detalles: " + error_export.Error(), ""
	}

	return 201, false, "", "Enviado correctamente"
}
