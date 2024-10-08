package carta

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	carta_automaticdiscount_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/carta_automaticdiscount"
	category_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/carta_category"
	element_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/carta_element"
	general_carta_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/carta_general"
	ordersstadistic_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/carta_orders-stadistic"
	schedule_range_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/carta_schedule_range"
)

/*----------------------CONSUMER----------------------*/

func UpdateCategory_Consumer_Service(idcategory int, urlphoto string, idbusiness int) (int, bool, string, string) {

	error_add_business := category_repository.Pg_Update_UrlPhoto(idcategory, urlphoto, idbusiness)

	if error_add_business != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la imagen de la categoria, detalles: " + error_add_business.Error(), ""
	}

	return 201, false, "", "Imagen actualizada correctamente"
}

func UpdateElement_Consumer_Service(idelement int, urlphoto string, idbusiness int) (int, bool, string, string) {

	error_add_business := element_repository.Pg_Update_UrlPhoto(idelement, urlphoto, idbusiness)

	if error_add_business != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la imagen del elemento, detalles: " + error_add_business.Error(), ""
	}

	return 201, false, "", "Imagen actualizada correctamente"
}

func Import_OrderStadistic_Service(orders_stadistic []models.Pg_Import_StadisticOrders) error {

	err_add_ordersstadistic := ordersstadistic_repository.Pg_Insert_OrderStadistic(orders_stadistic)
	if err_add_ordersstadistic != nil {
		log.Println(err_add_ordersstadistic)
	}

	return nil
}

/*----------------------CREATE DATA OF CARTA----------------------*/

func AddCategory_Service(idbusiness int, input_name_category string, input_typefood_category string, urlphoto string) (int, bool, string, int) {

	//Agregamos la categoria
	idcategory, error_add := category_repository.Pg_Add(idbusiness, input_name_category, input_typefood_category, urlphoto)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar agregar la categoria, detalles: " + error_add.Error(), 0
	}

	return 201, false, "", idcategory
}

func AddElement_Service(input_element models.Pg_Element) (int, bool, string, int) {

	//Agregamos el elemento
	idelement, error_add := element_repository.Pg_Add(input_element)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar agregar el elemento, detalles: " + error_add.Error(), 0
	}

	return 201, false, "", idelement
}

func AddScheduleRange_Service(idbusiness int, input_schedule models.Pg_ScheduleRange) (int, bool, string, string) {

	//Agregamos el rango horario
	error_add := schedule_range_repository.Pg_Add(idbusiness, input_schedule)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar agregar el rango horario, detalles: " + error_add.Error(), ""
	}

	return 201, false, "", "Rango horario creado correctamente"
}

func AddAutomaticDiscount_Service(idbusiness int, input_automaticdiscount models.Pg_AutomaticDiscount) (int, bool, string, string) {

	//Agregamos el rango horario
	error_add := carta_automaticdiscount_repository.Pg_Add(idbusiness, input_automaticdiscount)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar agregar el descuento automatico, detalles: " + error_add.Error(), ""
	}

	return 201, false, "", "Descuento automatico creado correctamente"
}

/*----------------------UDPATE ALL DATA OF INVENTARIO----------------------*/

func GetElementsByCategory_Service(idbusiness int, idcategory int) (int, bool, string, ElementsByCategory) {

	var elements_by_category ElementsByCategory

	elements, quantity, error_status_true := element_repository.Pg_Find_ByCategory(idbusiness, idcategory)
	if error_status_true != nil {
		return 500, true, "Error en el servidor interno al intentar buscar los elementos de esta categoria, detalles: " + error_status_true.Error(), elements_by_category
	}

	elements_by_category.Element = elements
	elements_by_category.Quantity = quantity

	return 201, false, "", elements_by_category
}

func UpdateCategoryStatus_Service(idbusiness int, idcategory int, statuscategory string) (int, bool, string, string) {

	if statuscategory == "true" {
		error_status_true := category_repository.Pg_Update_AvailableToTrue(idcategory, idbusiness)
		if error_status_true != nil {
			return 500, true, "Error en el servidor interno al intentar cambiar el estado de la categoria, detalles: " + error_status_true.Error(), ""
		}
	} else {
		error_status_false := category_repository.Pg_Update_AvailableToFalse(idcategory, idbusiness)
		if error_status_false != nil {
			return 500, true, "Error en el servidor interno al intentar cambiar el estado de la categoria, detalles: " + error_status_false.Error(), ""
		}
	}

	return 201, false, "", "Actualizado a :" + statuscategory
}

func UpdateElement_Service(idbusiness int, input_element models.Pg_Element) (int, bool, string, string) {

	//Agregamos la categoria
	error_udpate := element_repository.Pg_Update_Data(input_element, idbusiness)
	if error_udpate != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el elemento, detalles: " + error_udpate.Error(), ""
	}

	return 201, false, "", "Elemento actualizado correctamente"
}

func UpdateElementStatus_Service(idbusiness int, idelement int, statuselement bool) (int, bool, string, string) {

	error_status := element_repository.Pg_Update_Available(statuselement, idelement, idbusiness)
	if error_status != nil {
		return 500, true, "Error en el servidor interno al intentar cambiar el estado del elemento, detalles: " + error_status.Error(), ""
	}

	return 201, false, "", "Estado de elemento actualizado correctamente"
}

func UpdateScheduleRange_Service(idbusiness int, input_schedulerange models.Pg_ScheduleRange) (int, bool, string, string) {

	//Agregamos la categoria
	error_udpate := schedule_range_repository.Pg_Update_Data(input_schedulerange, idbusiness)
	if error_udpate != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el rango horario, detalles: " + error_udpate.Error(), ""
	}

	return 201, false, "", "Rango horario actualizado correctamente"
}

func UpdateScheduleRangeStatus_Service(idbusiness int, idschedule int) (int, bool, string, string) {

	error_status := schedule_range_repository.Pg_Delete(idschedule, idbusiness)
	if error_status != nil {
		return 500, true, "Error en el servidor interno al intentar eliminar el rango horario, detalles: " + error_status.Error(), ""
	}

	return 201, false, "", "Rango horario eliminado correctamente"
}

func UpdateAutomaticDiscount_Service(idbusiness int, input_automaticdiscount models.Pg_AutomaticDiscount) (int, bool, string, string) {

	//Agregamos la categoria
	error_udpate := carta_automaticdiscount_repository.Pg_Update_Data(input_automaticdiscount, idbusiness)
	if error_udpate != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el descuento automatico, detalles: " + error_udpate.Error(), ""
	}

	return 201, false, "", "Descuento automatico actualizado correctamente"
}

func DeleteAutomaticDiscount_Service(idbusiness int, iddiscount int) (int, bool, string, string) {

	error_status := carta_automaticdiscount_repository.Pg_Delete(iddiscount, idbusiness)
	if error_status != nil {
		return 500, true, "Error en el servidor interno al intentar eliminar el descuento automatico, detalles: " + error_status.Error(), ""
	}

	return 201, false, "", "Descuento automatico eliminado correctamente"
}

/*----------------------FIND ALL DATA OF INVENTARIO----------------------*/

func FindAllCategories_Service(input_idbusiness int) (int, bool, string, []models.Pg_Category_Response) {

	//Agregamos la categoria
	lista_category, error_add := category_repository.Pg_Find_All(input_idbusiness)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar listar las categorías de este negocio, detalles: " + error_add.Error(), lista_category
	}

	return 201, false, "", lista_category
}

func FindAllElements_Service(input_idbusiness int, input_limit int, input_offset int) (int, bool, string, []models.Pg_Element_Tofind) {

	//Agregamos la categoria
	lista_Elemento, error_add := element_repository.Pg_Find_All(input_idbusiness, input_limit, input_offset)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al ntentar listar los elementos de este negocio, detalles: " + error_add.Error(), lista_Elemento
	}

	return 201, false, "", lista_Elemento
}

func FindElementsRatingByDay_Service(input_idbusiness int, input_dayint int, input_limit int, input_offset int) (int, bool, string, []models.Pg_Element_WithRating) {

	//Agregamos la categoria
	lista_Elemento, error_add := element_repository.Pg_Find_ByDayWeek(input_idbusiness, input_dayint, input_limit, input_offset)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar listar el rating de los elementos de este negocio, detalles: " + error_add.Error(), lista_Elemento
	}

	return 201, false, "", lista_Elemento
}

func FindElementsRatingByName_Service(input_idbusiness int, name string) (int, bool, string, []models.Pg_Element_Tofind) {

	//Agregamos la categoria
	lista_Elemento, error_add := element_repository.Pg_Find_ByName(input_idbusiness, name)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar buscar el elementopor su nombre, detalles: " + error_add.Error(), lista_Elemento
	}

	return 201, false, "", lista_Elemento
}

func FindAllRangoHorario_Service(input_idbusiness int) (int, bool, string, []models.Pg_ScheduleRange) {

	//Agregamos la categoria
	lista_RangoHorario, error_add := schedule_range_repository.Pg_Find_All(input_idbusiness)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al ntentar listar los elementos de este negocio, detalles: " + error_add.Error(), lista_RangoHorario
	}

	return 201, false, "", lista_RangoHorario
}

func FindAllAutomaticDiscount_Service(input_idbusiness int) (int, bool, string, []models.Pg_AutomaticDiscount) {

	//Agregamos la categoria
	lista_AutomaticDiscount, error_add := carta_automaticdiscount_repository.Pg_Find_All(input_idbusiness)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar listar los los descuentos automaticos, detalles: " + error_add.Error(), lista_AutomaticDiscount
	}

	return 201, false, "", lista_AutomaticDiscount
}

/*----------------------OBTENER TODOS LOS DATOS DE CATEGORIA, ELEMENTO Y RANGO HORARIO----------------------*/

func FindAllCarta_MainData_Service(input_idbusiness int) (int, bool, string, models.Pg_Category_Element_ScheduleRange_AutomaticDiscount) {

	//Primero en la memoria cache
	carta_data_re, error_find_re := general_carta_repository.Re_Get_DataCard_Business(input_idbusiness)
	if error_find_re != nil || carta_data_re.CartaData.Category <= 0 {

		//Agregamos la categoria
		all_main_Data, error_add := general_carta_repository.Pg_Find_Main_Data(input_idbusiness)
		if error_add != nil {
			return 500, true, "Error en el servidor interno al intentar buscar losd atos de la carta, detalles: " + error_add.Error(), all_main_Data
		}

		error_add_re := general_carta_repository.Re_Set_DataCard_Business(input_idbusiness, all_main_Data)
		if error_add_re != nil {
			return 500, true, "Error en el servidor interno al intentar agregar los datos de la carta a la memoria cache, detalles: " + error_add_re.Error(), all_main_Data
		}

		return 201, false, "", all_main_Data

	}

	return 201, false, "", carta_data_re.CartaData
}

/*----------------------OBTENER TODOS LOS DATOS NEGOCIOS PARA NOTIFICARLOS----------------------*/

func SearchToNotifySchedulerange_Service() (int, bool, string, []int) {

	//Agregamos la categoria
	all_business, quantity, error_add := schedule_range_repository.Pg_SearchToNotify()
	if error_add != nil {
		return 500, true, "Error en el servidor interno al ntentar listar los negocios con datos de los rangos horarios a no notificar, detalles: " + error_add.Error(), all_business
	}

	if quantity > 0 {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":      "El rango horario es el horario de atención. Por ejemplo, horario mañana de 8:00 a 12:00, o, turno tarde de 14:00 a 18:00",
			"multipleuser": all_business,
			"typeuser":     6,
			"priority":     1,
			"title":        "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}

	return 201, false, "", all_business
}

/*----------------------ELIMINAR DATOS----------------------*/

func SendToDeleteCategory_Service(input_idbusiness int, timezone int, input_idcategory int) (int, bool, string, string) {

	error_sendtodelete := category_repository.Pg_Update_SendToDelete(input_idcategory, timezone, input_idbusiness)
	if error_sendtodelete != nil {
		return 500, true, "Error en el servidor interno al intentar enviar la categoria a la papelera, detalles: " + error_sendtodelete.Error(), ""
	}

	return 201, false, "", "Categoria enviada a papelera correctamente"
}

func RecoverSendToDeleteCategory_Service(input_idbusiness int, input_idcategory int) (int, bool, string, string) {

	error_recoversendtodelete := category_repository.Pg_Update_RecoverSendDelete(input_idcategory, input_idbusiness)
	if error_recoversendtodelete != nil {
		return 500, true, "Error en el servidor interno al intentar recuperar la categoria de la papelera, detalles: " + error_recoversendtodelete.Error(), ""
	}

	return 201, false, "", "Categoria recuperada correctamente"
}

func SendToDeleteElement_Service(input_idbusiness int, timezone int, input_idelement int) (int, bool, string, string) {

	error_sendtodelete := element_repository.Pg_Update_SendToDelete(input_idelement, timezone)
	if error_sendtodelete != nil {
		return 500, true, "Error en el servidor interno al intentar enviar la categoria a la papelera, detalles: " + error_sendtodelete.Error(), ""
	}

	return 201, false, "", "Element enviado a papelera correctamente"
}

func RecoverSendToDeleteElement_Service(input_idbusiness int, input_idelement int) (int, bool, string, string) {

	error_recoversendtodelete := element_repository.Pg_Update_RecoverSendDelete(input_idelement)
	if error_recoversendtodelete != nil {
		return 500, true, "Error en el servidor interno al intentar recuperar la categoria de la papelera, detalles: " + error_recoversendtodelete.Error(), ""
	}

	return 201, false, "", "Elemento recuperado correctamente"
}

/*----------------------DELETE----------------------*/

func UpdateCategory_Delete_Service() (string, string) {

	error_update := category_repository.Pg_Update_Delete()
	if error_update != nil {
		return "Error en el servidor interno al intentar eliminar la categoria, detalles: " + error_update.Error(), ""
	}

	return "", "Categoria eliminada correctamente"
}

func UpdateElement_Delete_Service() (string, string) {

	error_update := element_repository.Pg_Update_Delete()
	if error_update != nil {
		return "Error en el servidor interno al intentar eliminar el elemento, detalles: " + error_update.Error(), ""
	}

	return "", "Elemento eliminado correctamente"
}

/*----------------------FIND PAPELERA DATA----------------------*/

func FindCategory_Papelera_Service(idbusiness int) (int, bool, string, []models.Pg_Category_Response) {

	categorias, error_find := category_repository.Pg_Find_Papelera(idbusiness)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar listar las categorias, detalles: " + error_find.Error(), categorias
	}

	return 201, false, "", categorias
}

func FindElement_Papelera_Service(idbusiness int) (int, bool, string, []models.Pg_Element_Tofind) {

	providers, error_find := element_repository.Pg_Find_Papelera(idbusiness)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar listar los elementos, detalles: " + error_find.Error(), providers
	}

	return 201, false, "", providers
}
