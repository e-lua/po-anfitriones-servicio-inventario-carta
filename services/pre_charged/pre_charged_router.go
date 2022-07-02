package pre_charged

import (
	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"github.com/labstack/echo/v4"
)

var ImportsRouter_pg *importsRouter_pg

type importsRouter_pg struct {
}

/*----------------------UDPATE DATA CONSUME----------------------*/

func (ir *importsRouter_pg) AddPreCharged(c echo.Context) error {

	var pre_charged models.Mo_Precharged_Element

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&pre_charged)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos corretamente del proveedor, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddPreCharged_Service(pre_charged)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *importsRouter_pg) AddPreCharged_Multiple(c echo.Context) error {

	var pre_charged_multiple Request_Precarga

	err := c.Bind(&pre_charged_multiple)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos de elementos a precargar", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddPreCharged_Multiple_Service(pre_charged_multiple.AllPrecharged)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (ir *importsRouter_pg) FindPreCharged(c echo.Context) error {

	//Recibimos el nombre
	name := c.Request().URL.Query().Get("name")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindPreCharged_Service(name)
	results := ResponseListPreCharged{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}
