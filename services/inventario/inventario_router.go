package inventario

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"github.com/labstack/echo/v4"
)

var InventarioRouter_pg *inventarioRouter_pg

type inventarioRouter_pg struct {
}

/*----------------------CONSUMER----------------------*/

func (ir *inventarioRouter_pg) UpdateInsumo_Delete() {

	error_update, data := UpdateInsumo_Delete_Service()
	log.Fatal(error_update, data)

}

func (ir *inventarioRouter_pg) UpdateProvider_Delete() {

	error_update, data := UpdateProvider_Delete_Service()
	log.Fatal(error_update, data)
}

func (ir *inventarioRouter_pg) UpdateStoreHouse_Delete() {

	error_update, data := UpdateStoreHouse_Delete_Service()
	log.Fatal(error_update, data)
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

func GetJWTRol(jwt string) (int, bool, string, int, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:80/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0, 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness, get_respuesta.Data.IdRol
}

/*----------------------CREATE DATA OF INVENTARIO----------------------*/

func (ir *inventarioRouter_pg) AddProvider(c echo.Context) error {

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

	//Instanciamos una variable del modelo Proveedor
	var provider models.Mo_Providers

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&provider)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos corretamente del proveedor, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	provider.CreatedDate = time.Now()
	provider.IDBusiness = data_idbusiness
	provider.Available = true

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddProvider_Service(provider)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) AddStorehouse(c echo.Context) error {

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

	//Instanciamos una variable del modelo Almacen
	var storehouse models.Mo_StoreHouse

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&storehouse)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos corretamente del proveedor, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	storehouse.CreatedDate = time.Now()
	storehouse.IDBusiness = data_idbusiness
	storehouse.Available = true

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddStorehouse_Service(storehouse)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) AddInsumo(c echo.Context) error {

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

	//Instanciamos una variable del modelo Insumo
	var insumo models.Mo_Insumo

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&insumo)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos corretamente del proveedor, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	insumo.CreatedDate = time.Now()
	insumo.IDBusiness = data_idbusiness
	insumo.Available = true
	insumo.OutputStock = 0

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddInsumo_Service(insumo)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------UPDATE MAIN DATA----------------------*/

func (ir *inventarioRouter_pg) UpdateProvider_MainData(c echo.Context) error {

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

	//Instanciamos una variable del modelo Proveedor
	var provider models.Mo_Providers_Response

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&provider)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos corretamente del proveedor, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateProvider_MainData_Service(data_idbusiness, provider)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) UpdateStoreHouse_MainData(c echo.Context) error {

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

	//Instanciamos una variable del modelo Alamacen
	var storehouse models.Mo_StoreHouse_Response

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&storehouse)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos corretamente del proveedor, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateStoreHouse_MainData_Service(data_idbusiness, storehouse)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) UpdateInsumo_MainData(c echo.Context) error {

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

	//Instanciamos una variable del modelo Insumo
	var insumo models.Mo_Insumo_Response

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&insumo)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos corretamente del proveedor, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateInsumo_MainData_Service(data_idbusiness, insumo)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) UpdateInsumo_Stock(c echo.Context) error {

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

	//Recibimos el id de la insumo
	idinsumo := c.Param("idinsumo")

	//Instanciamos una variable del modelo Insumo
	var insumo_stock models.Mo_Insumo_Stock_Adjust_Requst

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&insumo_stock)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos corretamente del proveedor, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateInsumo_Stock_Service(idinsumo, data_idbusiness, insumo_stock)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------UPDATE AVAILABILITY----------------------*/

func (ir *inventarioRouter_pg) UpdateProvider_Availability(c echo.Context) error {

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

	//Recibimos el id de la proveedor
	idprovider := c.Param("idprovider")

	//Recibimos el estado
	status_provider := c.Param("status")
	status_bool, _ := strconv.ParseBool(status_provider)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateProvider_Availability_Service(idprovider, status_bool, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) UpdateStoreHouse_Availability(c echo.Context) error {

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

	//Recibimos el id de la storehouse
	idstorehouse := c.Param("idstorehouse")

	//Recibimos el estado
	status_provider := c.Param("status")
	status_bool, _ := strconv.ParseBool(status_provider)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateStoreHouse_Availability_Service(idstorehouse, status_bool, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) UpdateInsumo_Availability(c echo.Context) error {

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

	//Recibimos el id de la storehouse
	idinsumo := c.Param("idinsumo")

	//Recibimos el estado
	status_provider := c.Param("status")
	status_bool, _ := strconv.ParseBool(status_provider)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateInsumo_Availability_Service(idinsumo, status_bool, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------UPDATE SEND TO DELET----------------------*/

func (ir *inventarioRouter_pg) UpdateProvider_SendToDelete(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness, idrol := GetJWTRol(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if idrol != 1 && idrol != 3 {
		results := Response{Error: true, DataError: "Este rol no esta autorizado para esta acción", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id de la proveedor
	idprovider := c.Param("idprovider")

	//Recibimos el timezone
	timezone := c.Param("timezone")
	//Validamos los valores enviados
	if len(timezone) < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, debe enviar la zona horaria", Data: ""}
		return c.JSON(403, results)
	}

	timezone_int, _ := strconv.Atoi(timezone)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateProvider_SendToDelete_Service(idprovider, timezone_int, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) UpdateStoreHouse_SendToDelete(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness, idrol := GetJWTRol(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if idrol != 1 && idrol != 3 {
		results := Response{Error: true, DataError: "Este rol no esta autorizado para esta acción", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id de la storehouse
	idstorehouse := c.Param("idstorehouse")

	//Recibimos el timezone
	timezone := c.Param("timezone")
	//Validamos los valores enviados
	if len(timezone) < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, debe enviar la zona horaria", Data: ""}
		return c.JSON(403, results)
	}

	timezone_int, _ := strconv.Atoi(timezone)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateStoreHouse_SendToDelete_Service(idstorehouse, timezone_int, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) UpdateInsumo_SendToDelete(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness, idrol := GetJWTRol(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if idrol != 1 && idrol != 3 {
		results := Response{Error: true, DataError: "Este rol no esta autorizado para esta acción", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id de la insumo
	idinsumo := c.Param("idinsumo")

	//Recibimos el timezone
	timezone := c.Param("timezone")
	//Validamos los valores enviados
	if len(timezone) < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, debe enviar la zona horaria", Data: ""}
		return c.JSON(403, results)
	}

	timezone_int, _ := strconv.Atoi(timezone)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateInsumo_SendToDelete_Service(idinsumo, timezone_int, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------UPDATE RECOVER----------------------*/

func (ir *inventarioRouter_pg) UpdateProvider_RecoverSendToDelete(c echo.Context) error {

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

	//Recibimos el id de la proveedor
	idprovider := c.Param("idprovider")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateProvider_RecoverSendToDelete_Service(idprovider, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) UpdateStoreHouse_RecoverSendToDelete(c echo.Context) error {

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

	//Recibimos el id de la storehouse
	idstorehouse := c.Param("idstorehouse")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateStoreHouse_RecoverSendToDelete_Service(idstorehouse, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) UpdateInsumo_RecoverSendToDelete(c echo.Context) error {

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

	//Recibimos el id de la insumo
	idinsumo := c.Param("idinsumo")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateInsumo_RecoverSendToDelete_Service(idinsumo, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------FIND DATA----------------------*/

func (ir *inventarioRouter_pg) FindProvider_All(c echo.Context) error {

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

	//Recibimos el id de la proveedor
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)

	//Recibimos el id de la proveedor
	offset := c.Param("offset")
	offset_int, _ := strconv.Atoi(offset)

	if offset_int == 0 {
		offset_int = 1
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindProvider_All_Service(data_idbusiness, int64(limit_int), int64(offset_int))
	results := Response_Providers{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) FindStorehouse_All(c echo.Context) error {

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

	//Recibimos el id de la proveedor
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)

	//Recibimos el id de la proveedor
	offset := c.Param("offset")
	offset_int, _ := strconv.Atoi(offset)

	if offset_int == 0 {
		offset_int = 1
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindStorehouse_All_Service(data_idbusiness, int64(limit_int), int64(offset_int))
	results := Response_StoreHouse{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) FindInsumo_All(c echo.Context) error {

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

	//Recibimos el id de la proveedor
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)

	//Recibimos el id de la proveedor
	offset := c.Param("offset")
	offset_int, _ := strconv.Atoi(offset)

	if offset_int == 0 {
		offset_int = 1
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindInsumo_All_Service(data_idbusiness, int64(limit_int), int64(offset_int))
	results := Response_Insumo{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) FindInsumo_Stock(c echo.Context) error {

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

	//Recibimos el id de la proveedor
	idinsumo := c.Param("idinsumo")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindInsumo_Stock_Service(data_idbusiness, idinsumo)
	results := Response_Stock{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------FIND PAPELERA DATA----------------------*/

func (ir *inventarioRouter_pg) FindProvider_Papelera(c echo.Context) error {

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
	status, boolerror, dataerror, data := FindProvider_Papelera_Service(data_idbusiness)
	results := Response_Providers{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) FindStorehouse_Papelera(c echo.Context) error {

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
	status, boolerror, dataerror, data := FindStorehouse_Papelera_Service(data_idbusiness)
	results := Response_StoreHouse{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) FindInsumo_Papelera(c echo.Context) error {

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
	status, boolerror, dataerror, data := FindInsumo_Papelera_Service(data_idbusiness)
	results := Response_Insumo{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------SEARCH BY NAME----------------------*/

func (ir *inventarioRouter_pg) SearchNameProvider(c echo.Context) error {

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

	//Recibimos el nombre
	name := c.Request().URL.Query().Get("name")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SearchNameProvider_Service(data_idbusiness, name)
	results := Response_Providers{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) SearchNameStorehouse(c echo.Context) error {

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

	//Recibimos el nombre
	name := c.Request().URL.Query().Get("name")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SearchNameStorehouse_Service(data_idbusiness, name)
	results := Response_StoreHouse{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *inventarioRouter_pg) SearchNameInsumo(c echo.Context) error {

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

	//Recibimos el nombre
	name := c.Request().URL.Query().Get("name")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SearchNameInsumo_Service(data_idbusiness, name)
	results := Response_Insumo{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
