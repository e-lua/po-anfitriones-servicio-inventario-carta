package cartadiaria

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var CartaDiariaRouter_pg *cartaDiariaRouter_pg

type cartaDiariaRouter_pg struct {
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

func GetAddress(idbusiness string) (int, bool, string, B_Address) {
	//Obtenemos los datos del auth
	respuesta_2, _ := http.Get("http://c-busqueda.restoner-api.fun:6850/v1/business/address?idbusiness=" + idbusiness)
	var get_respuesta_2 ResponseAddress
	error_decode_respuesta_2 := json.NewDecoder(respuesta_2.Body).Decode(&get_respuesta_2)
	if error_decode_respuesta_2 != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la dirección, detalles: " + error_decode_respuesta_2.Error(), get_respuesta_2.Data
	}
	return 200, false, "", get_respuesta_2.Data
}

/*----------------------CREATE DATA OF MENU----------------------*/

func (cdr *cartaDiariaRouter_pg) AddCarta(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)

	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: boolerror, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var carta Carta

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, detalles: " + err.Error(), Data: 0}
		return c.JSON(403, results)
	}

	if !carta.WannaCopy {
		//Enviamos los datos al servicio
		status, boolerror, dataerror, data := AddCarta_Service(carta, data_idbusiness)
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		//Enviamos los datos al servicio
		status, boolerror, dataerror, data := AddCartaFromOther_Service(carta, data_idbusiness)
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}

}

func (cdr *cartaDiariaRouter_pg) UpdateCartaStatus(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var carta_status CartaStatus

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta_status)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCartaStatus_Service(carta_status, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) UpdateCartaElements(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Obtenemos los datos de latitud y longitud
	status, boolerror, dataerror, data_address := GetAddress(strconv.Itoa(data_idbusiness))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}

	//Instanciamos una variable del modelo CartaElements
	var carta_elements CartaElements_WithAction

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta_elements)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCartaElements_Service(carta_elements, data_idbusiness, data_address.Latitude, data_address.Longitude)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) UpdateCartaOneElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el JWT
	stock_string := c.Request().URL.Query().Get("stock")
	idelement_string := c.Request().URL.Query().Get("idelement")
	idcarta_string := c.Request().URL.Query().Get("idcarta")

	//Convertimos a int
	stock, _ := strconv.Atoi(stock_string)
	idcarta, _ := strconv.Atoi(idelement_string)
	idelement, _ := strconv.Atoi(idcarta_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCartaOneElement_Service(stock, idelement, idcarta, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) UpdateCartaScheduleRanges(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo CartaSchedule
	var carta_schedule CartaSchedule

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta_schedule)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCartaScheduleRanges_Service(carta_schedule, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------GET DATA OF MENU----------------------*/

func (cdr *cartaDiariaRouter_pg) GetCartaBasicData(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el limit
	date := c.Param("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaBasicData_Service(date, data_idbusiness)
	results := ResponseCartaBasicData{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartaCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el limit
	idcarta := c.Param("idcarta")
	idcarta_int, _ := strconv.Atoi(idcarta)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaCategory_Service(idcarta_int, data_idbusiness)
	results := ResponseCartaCategory{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartaElementsByCarta(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el idcarta
	idcarta := c.Param("idcarta")
	idcarta_int, _ := strconv.Atoi(idcarta)

	//Recibimos el idcategory
	idcategory := c.Param("idcategory")
	idcategory_int, _ := strconv.Atoi(idcategory)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaElementsByCarta_Service(idcarta_int, data_idbusiness, idcategory_int)
	results := ResponseCartaElements{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartaElements(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el limit
	idcarta := c.Param("idcarta")
	idcarta_int, _ := strconv.Atoi(idcarta)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaElements_Service(idcarta_int, data_idbusiness)
	results := ResponseCartaElements{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartaScheduleRanges(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el limit
	idcarta := c.Param("idcarta")
	idcarta_int, _ := strconv.Atoi(idcarta)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaScheduleRanges_Service(idcarta_int, data_idbusiness)
	results := ResponseCartaSchedule{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartas(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartas_Service(data_idbusiness)
	results := ResponseCartas{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------GET DATA OF MENU----------------------*/

func (cdr *cartaDiariaRouter_pg) DeleteCarta(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Category
	var carta_status CartaStatus

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta_status)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := DeleteCarta_Service(data_idbusiness, carta_status.IDCarta)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------GET DATA TO CREATE ORDER----------------------*/

func (cdr *cartaDiariaRouter_pg) GetCategories_ToCreateOrder(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCategories_ToCreateOrder_Service(date, data_idbusiness)
	results := ResponseCartaCategory_ToCreate{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetElements_ToCreateOrder(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el id de la categoria
	idcategory := c.Param("idcategory")
	idcategory_int, _ := strconv.Atoi(idcategory)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetElements_ToCreateOrder_Service(date, data_idbusiness, idcategory_int)
	results := ResponseCartaElements_ToCreate{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetSchedule_ToCreateOrder(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetSchedule_ToCreateOrder_Service(date, data_idbusiness)
	results := ResponseCartaSchedule_ToCreate{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------OBTENER TODOS LOS DATOS NEGOCIOS PARA NOTIFICARLOS----------------------*/

func (cdr *cartaDiariaRouter_pg) SearchToNotifyCarta() {

	//Enviamos los datos al servicio
	status, _, dataerror, _ := SearchToNotifyCarta_Service()
	log.Println(strconv.Itoa(status) + " " + dataerror)
}

/*----------------------DELETE----------------------*/

func (cdr *cartaDiariaRouter_pg) Delete_Vencidas() {

	error_delete, data := Delete_Vencidas_Service()
	log.Fatal(error_delete, data)

}

/*----------------------SEARCH TEXT----------------------*/

func (cr *cartaDiariaRouter_pg) SearchByNameAndDescription(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Recibimos el text
	text := c.Param("text")

	//Recibimos el limit
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Recibimos el limit
	offset := c.Param("offset")
	offset_int, _ := strconv.Atoi(offset)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SearchByNameAndDescription_Service(date, idbusiness_int, text, limit_int, offset_int)
	results := ResponseCartaElements_Searched{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
