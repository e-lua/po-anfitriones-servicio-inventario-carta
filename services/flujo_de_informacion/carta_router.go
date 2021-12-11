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

func (cr *cartaRouter_pg) UpdateCategory_Consumer(idcategory int, urlphoto string, idbusiness int) {

	//Enviamos los datos al servicio
	error_update_category := UpdateCategory_Consumer_Service(idcategory, urlphoto, idbusiness)
	if error_update_category != nil {
		log.Fatal(error_update_category)
	}
}

func (cr *cartaRouter_pg) UpdateElement_Consumer(idelement int, urlphoto string, idbusiness int) {

	//Enviamos los datos al servicio
	error_update_element := UpdateElement_Consumer_Service(idelement, urlphoto, idbusiness)
	if error_update_element != nil {
		log.Fatal(error_update_element)
	}
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://147.182.242.142:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

/*----------------------CREATE DATA OF CARTA----------------------*/

func (cr *cartaRouter_pg) AddCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "Token incorrecto", Data: 0}
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
	if len(category.Name) > 15 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddCategory_Service(data_idbusiness, category.Name)
	results := ResponseInt{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) AddElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, _ := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
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
	if len(element.Name) > 25 || element.IDCategory < 0 || element.Price < 0 || element.TypeMoney < 0 {
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
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "Token incorrecto", Data: 0}
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
	if len(scheduleRange.Name) > 12 || scheduleRange.MinutePerFraction < 0 || len(scheduleRange.StartTime) < 5 || len(scheduleRange.EndTime) < 5 || scheduleRange.MaxOrders < 0 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddScheduleRange_Service(data_idbusiness, scheduleRange)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------UDPATE ALL DATA OF CARTA----------------------*/

func (cr *cartaRouter_pg) UpdateCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "Token incorrecto", Data: 0}
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
	if category.IDCategory < 0 || len(category.Name) > 15 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCategory_Service(data_idbusiness, category)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) UpdateElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "Token incorrecto", Data: 0}
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
	if element.IDElement < 0 || len(element.Name) > 25 || element.IDCategory < 0 || element.Price < 0 || element.TypeMoney < 0 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateElement_Service(data_idbusiness, element)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (cr *cartaRouter_pg) UpdateScheduleRange(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "Token incorrecto", Data: 0}
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
	if scheduleRange.IDSchedule < 0 || len(scheduleRange.Name) > 12 || scheduleRange.MinutePerFraction < 0 || len(scheduleRange.StartTime) < 5 || len(scheduleRange.EndTime) < 5 || scheduleRange.MaxOrders < 0 {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateScheduleRange_Service(data_idbusiness, scheduleRange)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------FIND ALL DATA OF CARTA----------------------*/

func (cr *cartaRouter_pg) FindAllCategories(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "Token incorrecto", Data: 0}
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
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "Token incorrecto", Data: 0}
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

func (cr *cartaRouter_pg) FindAllRangoHorario(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllRangoHorario_Service(data_idbusiness)
	results := ResponseListRangoHorario{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------OBTENER TODOS LOS DATOS DE CATEGORIA, ELEMENTO Y RANGO HORARIO----------------------*/

func (cr *cartaRouter_pg) FindAllCarta_MainData(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: true, DataError: "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAllCarta_MainData_Service(data_idbusiness)
	results := ResponseAllMainData{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}
