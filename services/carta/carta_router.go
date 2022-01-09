package carta

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var CartaRouter_pg *cartaRouter_pg

type cartaRouter_pg struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

/*----------------------CREATE DATA OF MENU----------------------*/

func (cr *cartaRouter_pg) AddCarta(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: 0}
		return c.JSON(status, results)

	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: boolerror, DataError: "Token incorrecto", Data: 0}
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

func (cr *cartaRouter_pg) UpdateCartaStatus(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) UpdateCartaElements(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
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
	status, boolerror, dataerror, data := UpdateCartaElements_Service(carta_elements, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *cartaRouter_pg) UpdateCartaOneElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) UpdateCartaScheduleRanges(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) GetCartaBasicData(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el limit
	date := c.Param("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaBasicData_Service(date, data_idbusiness)
	results := ResponseCartaBasicData{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *cartaRouter_pg) GetCartaCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) GetCartaElementsByCarta(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) GetCartaElements(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) GetCartaScheduleRanges(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) GetCartas(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartas_Service(data_idbusiness)
	results := ResponseCartas{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------GET DATA OF MENU----------------------*/

func (cr *cartaRouter_pg) DeleteCarta(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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
