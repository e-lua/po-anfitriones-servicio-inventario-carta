package notification

import (
	"strconv"

	insumo_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/inventario_insumo"
)

/*----------------------------NOTIFICATION-----------------------------*/

func Notify_Ended_Service() (int, bool, string, []Mo_Notify_Insumo) {

	var notificaciones []Mo_Notify_Insumo

	data_insumos, error_add := insumo_repository.Mo_Find_Notify_Ended()
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar listar los insumos a notificar, detalles: " + error_add.Error(), notificaciones
	}

	for _, block_of_data := range data_insumos {

		var notificacion Mo_Notify_Insumo

		notificacion.Message = "Se ha acabado el stock de " + strconv.Itoa(block_of_data.Quantity) + " insumos"
		notificacion.IDUser = block_of_data.Idbusiness
		notificacion.TypeUser = 1
		notificacion.Priority = 1
		notificacion.Title = "⚠️ Alerta de Insumos 📦"

		notificaciones = append(notificaciones, notificacion)

	}

	return 201, false, "", notificaciones
}

func Notify_ToEnd_Service() (int, bool, string, []Mo_Notify_Insumo) {

	var notificaciones []Mo_Notify_Insumo

	data_insumos, error_add := insumo_repository.Mo_Find_Notify_ToEnded()
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar listar los insumos a notificar, detalles: " + error_add.Error(), notificaciones
	}

	for _, block_of_data := range data_insumos {

		var notificacion Mo_Notify_Insumo

		notificacion.Message = "Cuenta con " + strconv.Itoa(block_of_data.Quantity) + " insumos con muy poco stock"
		notificacion.IDUser = block_of_data.Idbusiness
		notificacion.TypeUser = 1
		notificacion.Priority = 1
		notificacion.Title = "Alerta de Insumos 📦"

		notificaciones = append(notificaciones, notificacion)

	}

	return 201, false, "", notificaciones
}
