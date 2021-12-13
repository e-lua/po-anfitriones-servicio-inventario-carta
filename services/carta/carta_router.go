package carta

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

var CartaRouter_pg *cartaRouter_pg

type cartaRouter_pg struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://143.198.75.79:5000/v1/trylogin?jwt=" + jwt)
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
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: 0}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddCarta_Service(carta, data_idbusiness)
	results := ResponseInt{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

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
