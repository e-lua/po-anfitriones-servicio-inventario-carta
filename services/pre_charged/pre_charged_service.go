package pre_charged

import (
	//REPOSITORIES

	"github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	pre_charged_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/pre_charged"
)

func AddPreCharged_Service(input_precharged models.Mo_Precharged_Element) (int, bool, string, string) {

	error_add_mo := pre_charged_repository.Mo_Add(input_precharged)
	if error_add_mo != nil {
		return 500, true, "Error en el servidor interno al intentar agregar el elemento a precargar, detalles: " + error_add_mo.Error(), ""
	}

	return 201, false, "", "Elemento registrado para su precarga"
}

func FindPreCharged_Service(name string) (int, bool, string, []*models.Mo_Precharged_Element) {

	elements, error_add_mo := pre_charged_repository.Mo_Search_Name(name)
	if error_add_mo != nil {
		return 500, true, "Error en el servidor interno al intentar buscar las imagenes precargadas, detalles: " + error_add_mo.Error(), elements
	}

	return 201, false, "", elements
}
