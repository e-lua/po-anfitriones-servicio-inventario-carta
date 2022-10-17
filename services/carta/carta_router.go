package carta

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"github.com/labstack/echo/v4"
)

var CartaRouter_pg *cartaRouter_pg

type cartaRouter_pg struct {
}

/*----------------------CONSUMER----------------------*/

func (cr *cartaRouter_pg) UpdateCategory_Consumer(c echo.Context) error {

	var toCarta models.Pg_ToCarta_Mqtt

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&toCarta)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "Se debe enviar el nombre de la categoria, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCategory_Consumer_Service(toCarta.IdBanner_Category_Element, toCarta.Url, toCarta.IdBusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *cartaRouter_pg) UpdateElement_Consumer(c echo.Context) error {

	var toCarta models.Pg_ToCarta_Mqtt

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&toCarta)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "Se debe enviar el nombre de la categoria, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateElement_Consumer_Service(toCarta.IdBanner_Category_Element, toCarta.Url, toCarta.IdBusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *cartaRouter_pg) Import_OrderStadistic(order_stadistic []models.Pg_Import_StadisticOrders) {

	//Enviamos los datos importados a registrar
	error_order_details := Import_OrderStadistic_Service(order_stadistic)
	log.Println(error_order_details)
}

func (cr *cartaRouter_pg) UpdateCategory_Delete() {

	error_update, data := UpdateCategory_Delete_Service()
	log.Println(error_update, data)

}

func (cr *cartaRouter_pg) UpdateElement_Delete() {

	error_update, data := UpdateElement_Delete_Service()
	log.Println(error_update, data)
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:80/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

/*----------------------CREATE DATA OF INVENTARIO----------------------*/

func (cr *cartaRouter_pg) AddCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var category models.Pg_Category

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&category)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "Se debe enviar el nombre de la categoria, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(category.Name) > 50 || len(category.TypeFood) < 2 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddCategory_Service(data_idbusiness, category.Name, category.TypeFood, category.UrlPhoto)
	results := ResponseInt{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) AddElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var element models.Pg_Element

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&element)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "Se debe enviar el nombre del elemento, la categoría, el precio y el tipo de moneda, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(element.Name) > 80 || element.IDCategory < 0 || element.Price < 0 || element.TypeMoney < 0 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddElement_Service(element)
	results := ResponseInt{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) AddScheduleRange(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var scheduleRange models.Pg_ScheduleRange

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&scheduleRange)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "Se debe enviar el nombre del rango horario, minutos por fraccion, tiempo de inicio, tiempo de fin, y las ordenes maximas, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if len(scheduleRange.Name) > 20 || scheduleRange.NumberOfFractions <= 0 || scheduleRange.MinutePerFraction < 0 || len(scheduleRange.StartTime) != 5 || len(scheduleRange.EndTime) != 5 || scheduleRange.MaxOrders < 0 || scheduleRange.TimeZone == "" {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddScheduleRange_Service(data_idbusiness, scheduleRange)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) AddAutomaticDiscount(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var automaticDsicount models.Pg_AutomaticDiscount

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&automaticDsicount)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "Se debe enviar la descripción,el descuento, el grupo con mas de 1 combinación, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if automaticDsicount.Description == "" || automaticDsicount.Discount <= 0 || len(automaticDsicount.Group) < 1 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddAutomaticDiscount_Service(data_idbusiness, automaticDsicount)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------UDPATE ALL DATA OF CARTA----------------------*/

func (cr *cartaRouter_pg) GetElementsByCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Recibimos el id de la categoria
	idcategory := c.Param("idcategory")
	idcategory_int, _ := strconv.Atoi(idcategory)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetElementsByCategory_Service(data_idbusiness, idcategory_int)
	results := ResponseElementsByCategory{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) UpdateCategoryStatus(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Recibimos el id de la categoria
	idcategory := c.Param("idcategory")
	idcategory_int, _ := strconv.Atoi(idcategory)

	//Recibimos el estado
	statuscategory := c.Param("status")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCategoryStatus_Service(data_idbusiness, idcategory_int, statuscategory)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) UpdateElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var element models.Pg_Element

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&element)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "Se debe enviar el nombre del elemento, la categoría, el precio y el tipo de moneda, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if element.IDCategory < 0 || element.Price < 0 || element.TypeMoney != 0 && element.TypeMoney != 1 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateElement_Service(data_idbusiness, element)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) UpdateElementStatus(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Recibimos el id de la categoria
	idelement := c.Param("idelement")
	idelement_int, _ := strconv.Atoi(idelement)

	//Recibimos el estado
	statusElement := c.Param("status")
	status_bool_element, _ := strconv.ParseBool(statusElement)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateElementStatus_Service(data_idbusiness, idelement_int, status_bool_element)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) UpdateScheduleRange(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var scheduleRange models.Pg_ScheduleRange

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&scheduleRange)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "Se debe enviar el nombre del elemento, la categoría, el precio y el tipo de moneda, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if scheduleRange.IDSchedule < 0 || len(scheduleRange.StartTime) < 4 || len(scheduleRange.EndTime) < 4 || scheduleRange.NumberOfFractions <= 0 || len(scheduleRange.Name) > 12 || scheduleRange.MinutePerFraction < 0 || scheduleRange.MaxOrders < 0 || scheduleRange.TimeZone == "" {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateScheduleRange_Service(data_idbusiness, scheduleRange)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) UpdateAutomaticDiscount(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var automaticDsicount models.Pg_AutomaticDiscount

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&automaticDsicount)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "Se debe enviar la descripción,el descuento, el grupo con mas de 1 combinación, revise la estructura o los valores", Data: 0}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if automaticDsicount.Description == "" || automaticDsicount.Discount == 0 || len(automaticDsicount.Group) < 1 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateAutomaticDiscount_Service(data_idbusiness, automaticDsicount)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) UpdateScheduleRangeStatus(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Recibimos el id de la categoria
	idschedulerange := c.Param("idschedulerange")
	idschedulerange_int, _ := strconv.Atoi(idschedulerange)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateScheduleRangeStatus_Service(data_idbusiness, idschedulerange_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) DeleteAutomaticDiscount(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Recibimos el id de la categoria
	idautomaticdiscount := c.Param("idautomaticdiscount")
	idautomaticdiscount_int, _ := strconv.Atoi(idautomaticdiscount)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := DeleteAutomaticDiscount_Service(data_idbusiness, idautomaticdiscount_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------FIND ALL DATA OF CARTA----------------------*/

func (cr *cartaRouter_pg) FindAllCategories(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllCategories_Service(data_idbusiness)
	results := ResponseListCategory{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) FindAllElements(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Recibimos el limit
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Recibimos el limit
	offset := c.Param("offset")
	offset_int, _ := strconv.Atoi(offset)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllElements_Service(data_idbusiness, limit_int, offset_int)
	results := ResponseListElement{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) FindElementsRatingByDay(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Recibimos el day
	day := c.Param("day")
	day_int, _ := strconv.Atoi(day)
	//Recibimos el limit
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Recibimos el offset
	offset := c.Param("offset")
	offset_int, _ := strconv.Atoi(offset)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindElementsRatingByDay_Service(data_idbusiness, day_int, limit_int, offset_int)
	results := ResponseListElement_WithRating{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) FindElementsRatingByName(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Recibimos el nombre
	name := c.Request().URL.Query().Get("name")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindElementsRatingByName_Service(data_idbusiness, name)
	results := ResponseListElement{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) FindAllRangoHorario(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllRangoHorario_Service(data_idbusiness)
	results := ResponseListRangoHorario{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) FindAllAutomaticDiscount(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllAutomaticDiscount_Service(data_idbusiness)
	results := ResponseListAutomaticDiscount{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------OBTENER TODOS LOS DATOS DE CATEGORIA, ELEMENTO Y RANGO HORARIO----------------------*/

func (cr *cartaRouter_pg) FindAllCarta_MainData(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllCarta_MainData_Service(data_idbusiness)
	results := ResponseAllMainData{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------OBTENER TODOS LOS DATOS NEGOCIOS PARA NOTIFICARLOS----------------------*/

func (cr *cartaRouter_pg) SearchToNotifySchedulerange() {

	//Enviamos los datos al servicio
	status, _, dataerror, _ := SearchToNotifySchedulerange_Service()
	log.Println(strconv.Itoa(status) + " " + dataerror)
	//results := ResponseAllBusinesses{Error: boolerror, DataError: dataerror, Data: data}
}

/*----------------------ELIMINAR DATOS----------------------*/

func (cr *cartaRouter_pg) SendToDeleteCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el timezone
	timezone := c.Param("timezone")
	//Validamos los valores enviados
	if len(timezone) < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, debe enviar la zona horaria", Data: ""}
		return c.JSON(403, results)
	}

	timezone_int, _ := strconv.Atoi(timezone)

	//Recibimos el id de la categoria
	idcategory := c.Param("idcategory")
	idcategorye_int, _ := strconv.Atoi(idcategory)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SendToDeleteCategory_Service(data_idbusiness, timezone_int, idcategorye_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) RecoverSendToDeleteCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id de la categoria
	idcategory := c.Param("idcategory")
	idcategorye_int, _ := strconv.Atoi(idcategory)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := RecoverSendToDeleteCategory_Service(data_idbusiness, idcategorye_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) SendToDeleteElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el timezone
	timezone := c.Param("timezone")
	//Validamos los valores enviados
	if len(timezone) < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, debe enviar la zona horaria", Data: ""}
		return c.JSON(403, results)
	}

	timezone_int, _ := strconv.Atoi(timezone)

	//Recibimos el id de la element
	idelement := c.Param("idelement")
	idelement_int, _ := strconv.Atoi(idelement)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SendToDeleteElement_Service(data_idbusiness, timezone_int, idelement_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) RecoverSendToDeleteElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id de la element
	idelement := c.Param("idelement")
	idelement_int, _ := strconv.Atoi(idelement)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := RecoverSendToDeleteElement_Service(data_idbusiness, idelement_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------FIND PAPELERA DATA----------------------*/

func (cr *cartaRouter_pg) FindCategory_Papelera(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindCategory_Papelera_Service(data_idbusiness)
	results := ResponseListCategory{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *cartaRouter_pg) FindElement_Papelera(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindElement_Papelera_Service(data_idbusiness)
	results := ResponseListElement{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
