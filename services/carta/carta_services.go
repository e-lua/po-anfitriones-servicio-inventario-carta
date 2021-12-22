package carta

import (
	//REPOSITORIES
	"log"
	"time"

	"github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	carta_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/carta"
)

func AddCarta_Service(input_carta Carta, idbusiness int) (int, bool, string, int) {

	//Insertamos los datos en Mo
	idcarta, error_add_carta := carta_repository.Pg_Add(idbusiness, input_carta.Date)
	if error_add_carta != nil {
		return 500, true, "Error en el servidor interno al intentar crear la carta, detalles: " + error_add_carta.Error(), 0
	}

	return 201, false, "", idcarta
}

func UpdateCartaStatus_Service(carta_status CartaStatus, idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_update := carta_repository.Pg_Update_Available_Visible(carta_status.Available, carta_status.Visible, carta_status.IDCarta, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la visibilidad y disponibilidad de la carta , detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "La disponibilidad y visibilidad se actualizaron correctamente"
}

func UpdateCartaOneElement_Service(stock int, idelement int, idcarta int, idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_update := carta_repository.Pg_Update_One_Element(stock, idelement, idcarta, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el elemento, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Elemento actualizado correctamente"
}

func UpdateCartaElements_Service(carta_elements CartaElements, idbusiness int) (int, bool, string, string) {

	//Eliminamos los datos anteriores
	carta_repository.Pg_Delete_Elements(carta_elements.IDCarta, idbusiness)

	//Insertamos los datos en Mo
	go func() {
		error_update := carta_repository.Pg_Update_Elements(carta_elements.Elements, carta_elements.IDCarta, idbusiness)
		if error_update != nil {
			log.Fatal("Error en el servidor interno al intentar actualizar los elementos, detalles: " + error_update.Error())
		}
	}()
	time.Sleep(1 * time.Second)

	return 201, false, "", "Los elementos se actualizaron correctamnte"
}

func UpdateCartaScheduleRanges_Service(carta_schedule CartaSchedule, idbusiness int) (int, bool, string, string) {

	//Eliminamos los datos anteriores
	error_delete_list := carta_repository.Pg_Delete_ScheduleRange_List(carta_schedule.IDCarta, idbusiness)
	if error_delete_list != nil {
		log.Fatal("Error en el servidor interno al intentar eliminar la lista de rangos horarios, detalles: " + error_delete_list.Error())
	}

	error_delete := carta_repository.Pg_Delete_ScheduleRange(carta_schedule.IDCarta, idbusiness)
	if error_delete != nil {
		log.Fatal("Error en el servidor interno al intentar eliminar los rangos horarios, detalles: " + error_delete.Error())
	}

	go func() {
		error_update := carta_repository.Pg_Update_ScheduleRange(carta_schedule.ScheduleRanges, carta_schedule.IDCarta, idbusiness)
		if error_update != nil {
			log.Fatal("Error en el servidor interno al intentar actualizar los rangos horarios, detalles: " + error_update.Error())
		}
	}()

	//Insertamos los datos en la lista de horario
	go func() {
		error_update := carta_repository.Pg_Update_ScheduleRange_List(carta_schedule.ScheduleRanges, carta_schedule.IDCarta, idbusiness)
		if error_update != nil {
			log.Fatal("Error en el servidor interno al intentar actualizar la lista de rangos horarios, detalles: " + error_update.Error())
		}
	}()

	time.Sleep(1 * time.Second)

	return 201, false, "", "Los rangos horario se actualizaron correctamente"
}

/*----------------------GET DATA OF MENU----------------------*/

func GetCartaBasicData_Service(date string, idbusiness int) (int, bool, string, models.Pg_Carta_External) {

	//Insertamos los datos en Mo
	carta_ini_values, error_show := carta_repository.Pg_Find_IniData(date, idbusiness)
	if error_show != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar la informacion basica de la carta, detalles: " + error_show.Error(), carta_ini_values
	}

	return 201, false, "", carta_ini_values
}

func GetCartaCategory_Service(idcarta_int int, idbusiness int) (int, bool, string, []models.Pg_Category_External) {

	//Obtenemos las categorias
	carta_category, error_update := carta_repository.Pg_Find_Category(idcarta_int, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_category
	}

	return 201, false, "", carta_category
}

func GetCartaElementsByCarta_Service(idcarta_int int, idbusiness int, idcategory int) (int, bool, string, []models.Pg_Element_With_Stock_External) {

	//Obtenemos las categorias
	carta_category, error_update := carta_repository.Pg_Find_Elements_ByCategory(idcarta_int, idbusiness, idcategory)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los elementos de la categoria seleccionada, detalles: " + error_update.Error(), carta_category
	}

	return 201, false, "", carta_category
}

func GetCartaElements_Service(idcarta_int int, idbusiness int) (int, bool, string, []models.Pg_Element_With_Stock_External) {

	//Insertamos los datos en Mo
	carta_elements, error_update := carta_repository.Pg_Find_Elements(idcarta_int, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los elementos de la carta, detalles: " + error_update.Error(), carta_elements
	}

	return 201, false, "", carta_elements
}

func GetCartaScheduleRanges_Service(idcarta_int int, idbusiness int) (int, bool, string, []models.Pg_ScheduleRange_External) {

	//Insertamos los datos en Mo
	carta_scheduleranges, error_update := carta_repository.Pg_Find_ScheduleRanges(idcarta_int, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los rangos horarios de la carta, detalles: " + error_update.Error(), carta_scheduleranges
	}

	return 201, false, "", carta_scheduleranges
}

func GetCartas_Service(idbusiness int) (int, bool, string, []models.Pg_Carta_Found) {

	//Insertamos los datos en Mo
	carta_found, error_update := carta_repository.Pg_Find_Cartas(idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los rangos horarios de la carta, detalles: " + error_update.Error(), carta_found
	}

	return 201, false, "", carta_found
}

/*----------------------COPY BETWEEN MENUS----------------------*/

func AddCartaFromOther_Service(input_carta Carta, idbusiness int) (int, bool, string, int) {

	//Buscamos la carta
	idcarta_int, error_add_carta := carta_repository.Pg_Find_IniData(input_carta.FromCarta, idbusiness)
	if error_add_carta != nil {
		return 500, true, "Error en el servidor interno al intentar crear la carta, detalles: " + error_add_carta.Error(), 0
	}

	carta_elements, error_update_element := carta_repository.Pg_Find_Elements(idcarta_int.IDCarta, idbusiness)
	if error_update_element != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los elementos de la carta, detalles: " + error_update_element.Error(), 0
	}

	carta_scheduleranges, error_update_schedule := carta_repository.Pg_Find_ScheduleRanges(idcarta_int.IDCarta, idbusiness)
	if error_update_schedule != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los rangos horarios de la carta, detalles: " + error_update_schedule.Error(), 0
	}

	//Insertamos los datos en Mo
	/*idcarta, error_add_carta := carta_repository.Pg_Add(idbusiness, input_carta.Date)
	if error_add_carta != nil {
		return 500, true, "Error en el servidor interno al intentar crear la carta, detalles: " + error_add_carta.Error(), 0
	}

	error_update_element_nc := carta_repository.Pg_Update_Elements(carta_elements, idcarta, idbusiness)
	if error_update_element_nc != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los elementos, detalles: " + error_update_element_nc.Error(), 0
	}

	error_update_schedule_nc := carta_repository.Pg_Update_ScheduleRange(carta_scheduleranges, idcarta, idbusiness)
	if error_update_schedule_nc != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los rangos horarios, detalles: " + error_update_schedule_nc.Error(), 0
	}

	//Insertamos los datos en la lista de horario
	error_update_schedulelist := carta_repository.Pg_Update_ScheduleRange_List(carta_scheduleranges, idcarta, idbusiness)
	if error_update_schedulelist != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la lista de rangos horarios, detalles: " + error_update_schedulelist.Error(), 0
	}*/

	//Transaccion
	id_carta, error_update_schedulelist := carta_repository.Pg_Copy_Carta(carta_scheduleranges, carta_elements, idbusiness, input_carta.Date)
	if error_update_schedulelist != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la lista de rangos horarios, detalles: " + error_update_schedulelist.Error(), 0
	}

	return 201, false, "", id_carta
}

/*----------------------DELETE MENU----------------------*/

func DeleteCarta_Service(idbusiness int, idcarta int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_delete := carta_repository.Pg_Delete(idbusiness, idcarta)
	if error_delete != nil {
		return 500, true, "Error en el servidor interno al intentar eliminar la carta, detalles: " + error_delete.Error(), ""
	}

	return 201, false, "", "Eliminado correctamente"
}
