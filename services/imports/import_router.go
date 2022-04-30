package imports

import (
	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	"github.com/labstack/echo/v4"
)

var ImportsRouter_pg *importsRouter_pg

type importsRouter_pg struct {
}

/*----------------------UDPATE DATA CONSUME----------------------*/

func (ir *importsRouter_pg) UpdateElementStock(c echo.Context) error {

	var input_insumos []models.Mqtt_Import_InsumoStock

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&input_insumos)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, detalles: " + err.Error(), Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateInsumoStock_Service(input_insumos)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
