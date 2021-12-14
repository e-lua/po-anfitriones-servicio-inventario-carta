package carta

import (
	//REPOSITORIES
	"github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
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
	error_update := carta_repository.Pg_Update_Available_Visible(carta_status.Available, carta_status.Visible, carta_status.IDCarta, idbusiness)
	if error_update != nil {
		return 404, true, "Error en el servidor interno al intentar actualizar la visibilidad y disponibilidad de la carta , detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "La disponibilidad y visibilidad se actualizaron correctamente"
}

func UpdateCartaOneElement_Service(stock int, idelement int, idcarta int, idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_update := carta_repository.Pg_Update_One_Element(stock, idelement, idcarta, idbusiness)
	if error_update != nil {
		return 404, true, "Error en el servidor interno al intentar actualizar el elemento, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Elemento actualizado correctamente"
}

func UpdateCartaElements_Service(carta_elements CartaElements, idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_update := carta_repository.Pg_Update_Elements(carta_elements.Elements, carta_elements.IDCarta, idbusiness)
	if error_update != nil {
		return 404, true, "Error en el servidor interno al intentar actualizar los elementos de la carta, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Los elementos se actualizaron correctamnte"
}

func UpdateCartaScheduleRanges_Service(carta_schedule CartaSchedule, idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_update := carta_repository.Pg_Update_ScheduleRange_List(carta_schedule.ScheduleRanges, carta_schedule.IDCarta, idbusiness)
	if error_update != nil {
		return 404, true, "Error en el servidor interno al intentar actualizar los rangos horarios de la carta, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Los rangos horario se actualizaron correctamente"
}

/*----------------------GET DATA OF MENU----------------------*/

func GetCartaBasicData_Service(date string, idbusiness int) (int, bool, string, models.Pg_Carta_External) {

	//Insertamos los datos en Mo
	carta_ini_values, error_update := carta_repository.Pg_Find_IniData(date, idbusiness)
	if error_update != nil {
		return 404, true, "Error en el servidor interno al intentar encontrar la informacion basica de la carta, detalles: " + error_update.Error(), carta_ini_values
	}

	return 201, false, "", carta_ini_values
}

func GetCartaElements_Service(idcarta_int int, idbusiness int) (int, bool, string, []models.Pg_Element_With_Stock_External) {

	//Insertamos los datos en Mo
	carta_elements, error_update := carta_repository.Pg_Find_Elements(idcarta_int, idbusiness)
	if error_update != nil {
		return 404, true, "Error en el servidor interno al intentar encontrar los elementos de la carta, detalles: " + error_update.Error(), carta_elements
	}

	return 201, false, "", carta_elements
}

func GetCartaScheduleRanges_Service(idcarta_int int, idbusiness int) (int, bool, string, []models.Pg_ScheduleRange_External) {

	//Insertamos los datos en Mo
	carta_scheduleranges, error_update := carta_repository.Pg_Find_ScheduleRanges(idcarta_int, idbusiness)
	if error_update != nil {
		return 404, true, "Error en el servidor interno al intentar encontrar los rangos horarios de la carta, detalles: " + error_update.Error(), carta_scheduleranges
	}

	return 201, false, "", carta_scheduleranges
}
