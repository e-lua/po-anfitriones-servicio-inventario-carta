package notification

import (
	"github.com/labstack/echo/v4"
)

var NotificationRouter_pg *notificationRouter_pg

type notificationRouter_pg struct {
}

/*----------------------------NOTIFICATION-----------------------------*/

func (nr *notificationRouter_pg) Notify_Ended(c echo.Context) error {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Notify_Ended_Service()
	results := Response_Notify_Insumo{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *notificationRouter_pg) Notify_ToEnd(c echo.Context) error {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Notify_ToEnd_Service()
	results := Response_Notify_Insumo{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
