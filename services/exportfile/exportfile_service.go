package exportfile

import (
	//REPOSITORIES

	export_file "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/export_to_file"
)

func ExportFile_Insumo_Service(idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_export := export_file.Mo_Insumos_ToFile(idbusiness)
	if error_export != nil {
		return 500, true, "Error en el servidor interno al intentar exportar los insumos, detalles: " + error_export.Error(), ""
	}

	return 201, false, "", "Enviado correctamente"
}

func ExportFile_Element_Service(idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_export := export_file.Pg_Elements_ToFile(idbusiness)
	if error_export != nil {
		return 500, true, "Error en el servidor interno al intentar exportar los elementos, detalles: " + error_export.Error(), ""
	}

	return 201, false, "", "Enviado correctamente"
}
