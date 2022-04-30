package imports

import (
	"log"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
)

var ImportsRouter_pg *importsRouter_pg

type importsRouter_pg struct {
}

/*----------------------UDPATE DATA CONSUME----------------------*/

func (ir *importsRouter_pg) UpdateElementStock(import_elements []models.Mqtt_Import_InsumoStock) {

	//Enviamos los datos al servicio
	error_update := UpdateInsumoStock_Service(import_elements)
	if error_update != nil {
		log.Println("Error al actualizar el stock de los insumos: " + error_update.Error())
	}

}
