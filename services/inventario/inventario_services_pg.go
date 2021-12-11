package carta

import (
	"log"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	category_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/category"
	element_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/element"
	general_carta_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/general"
	schedule_range_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/schedule_range"
)

/*----------------------CONSUMER----------------------*/

func UpdateCategory_Consumer_Service(idcategory int, urlphoto string, idbusiness int) error {

	error_add_business := category_repository.Pg_Update_UrlPhoto(idcategory, urlphoto, idbusiness)

	if error_add_business != nil {
		log.Fatal(error_add_business)
	}

	return nil
}

func UpdateElement_Consumer_Service(idelement int, urlphoto string, idbusiness int) error {

	error_add_business := element_repository.Pg_Update_UrlPhoto(idelement, urlphoto, idbusiness)

	if error_add_business != nil {
		log.Fatal(error_add_business)
	}

	return nil
}

/*----------------------CREATE DATA OF CARTA----------------------*/

func AddCategory_Service(idbusiness int, input_name_category string) (int, bool, string, int) {

	//Agregamos la categoria
	idcategory, error_add := category_repository.Pg_Add(idbusiness, input_name_category)
	if error_add != nil {
		return 404, true, "Error en el servidor interno al intentar agregar la categoria, detalles: " + error_add.Error(), 0
	}

	return 201, false, "", idcategory
}

func AddElement_Service(input_element models.Pg_Element) (int, bool, string, int) {

	//Agregamos la categoria
	idelement, error_add := element_repository.Pg_Add(input_element)
	if error_add != nil {
		return 404, true, "Error en el servidor interno al intentar agregar el elemento, detalles: " + error_add.Error(), 0
	}

	return 201, false, "", idelement
}

func AddScheduleRange_Service(idbusiness int, input_schedule models.Pg_ScheduleRange) (int, bool, string, string) {

	//Agregamos la categoria
	error_add := schedule_range_repository.Pg_Add(idbusiness, input_schedule)
	if error_add != nil {
		return 404, true, "Error en el servidor interno al intentar agregar el rango horario, detalles: " + error_add.Error(), ""
	}

	return 201, false, "", "Rango horario creado correctamente"
}

/*----------------------UDPATE ALL DATA OF CARTA----------------------*/

func UpdateElement_Service(idbusiness int, input_element models.Pg_Element) (int, bool, string, string) {

	//Agregamos la categoria
	error_udpate := element_repository.Pg_Update_Data(input_element, idbusiness)
	if error_udpate != nil {
		return 404, true, "Error en el servidor interno al intentar actualizar el elemento, detalles: " + error_udpate.Error(), ""
	}

	return 201, false, "", "Elemento actualizado correctamente"
}

func UpdateScheduleRange_Service(idbusiness int, input_schedulerange models.Pg_ScheduleRange) (int, bool, string, string) {

	//Agregamos la categoria
	error_udpate := schedule_range_repository.Pg_Update_Data(input_schedulerange, idbusiness)
	if error_udpate != nil {
		return 404, true, "Error en el servidor interno al intentar actualizar el rango horario, detalles: " + error_udpate.Error(), ""
	}

	return 201, false, "", "Rango horario actualizado correctamente"
}

/*----------------------FIND ALL DATA OF CARTA----------------------*/

func FindAllCategories_Service(input_idbusiness int) (int, bool, string, []models.Pg_Category) {

	//Agregamos la categoria
	lista_category, error_add := category_repository.Pg_Find_All(input_idbusiness)
	if error_add != nil {
		return 404, true, "Error en el servidor interno al intentar listar las categorías de este negocio, detalles: " + error_add.Error(), lista_category
	}

	return 201, false, "", lista_category
}

func FindAllElements_Service(input_idbusiness int, input_limit int, input_offset int) (int, bool, string, []models.Pg_Element) {

	//Agregamos la categoria
	lista_Elemento, error_add := element_repository.Pg_Find_All(input_idbusiness, input_limit, input_offset)
	if error_add != nil {
		return 404, true, "Error en el servidor interno al ntentar listar los elementos de este negocio, detalles: " + error_add.Error(), lista_Elemento
	}

	return 201, false, "", lista_Elemento
}

func FindAllRangoHorario_Service(input_idbusiness int) (int, bool, string, []models.Pg_ScheduleRange) {

	//Agregamos la categoria
	lista_RangoHorario, error_add := schedule_range_repository.Pg_Find_All(input_idbusiness)
	if error_add != nil {
		return 404, true, "Error en el servidor interno al ntentar listar los elementos de este negocio, detalles: " + error_add.Error(), lista_RangoHorario
	}

	return 201, false, "", lista_RangoHorario
}

/*----------------------OBTENER TODOS LOS DATOS DE CATEGORIA, ELEMENTO Y RANGO HORARIO----------------------*/

func FindAllCarta_MainData_Service(input_idbusiness int) (int, bool, string, models.Pg_Category_Element_ScheduleRange) {

	//Agregamos la categoria
	all_main_Data, error_add := general_carta_repository.Pg_Find_Main_Data(input_idbusiness)
	if error_add != nil {
		return 404, true, "Error en el servidor interno al ntentar listar los elementos de este negocio, detalles: " + error_add.Error(), all_main_Data
	}

	return 201, false, "", all_main_Data
}
