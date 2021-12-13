package carta

import (
	//REPOSITORIES
	carta_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/carta"
)

func AddCarta_Service(input_carta Carta, idbusiness int) (int, bool, string, int) {

	//Insertamos los datos en Mo
	idcarta, error_add_carta := carta_repository.Pg_Add(idbusiness, input_carta.Date)
	if error_add_carta != nil {
		return 404, true, "Error en el servidor interno al intentar crear la carta, detalles: " + error_add_carta.Error(), 0
	}

	return 201, false, "", idcarta
}

func UpdateCartaStatus_Service(carta_status CartaStatus, idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_update := carta_repository.Pg_Update_AvailableToFalse(carta_status.Available, carta_status.Visible, carta_status.IDCarta, idbusiness)
	if error_update != nil {
		return 404, true, "Error en el servidor interno al intentar actualizar la visibilidad y disponibilidad de la carta , detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "La disponibilidad y visibilidad se actualizaron correctamnte"
}
