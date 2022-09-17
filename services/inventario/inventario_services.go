package inventario

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	insumo_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/inventario_insumo"
	provider_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/inventario_provider"
	store_repository "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/repositories/inventario_storehouse"
)

/*----------------------------NOTIFICATION-----------------------------*/

func Notify_Ended_Service() (int, bool, string, string) {

	//ahora := time.Now()

	//if ahora.Hour() == 14 {

	data_insumos, error_add := insumo_repository.Mo_Find_Notify_Ended()
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar listar los insumos a notificar, detalles: " + error_add.Error(), ""
	}

	for _, block_of_data := range data_insumos {

		log.Println("PRIMER VALOR block_of_data.Idbusiness----->", block_of_data.Idbusiness)
		log.Println("PRIMER VALOR block_of_data.Quantity----->", block_of_data.Quantity)

		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se ha acabado el stock de " + strconv.Itoa(block_of_data.Quantity) + " insumos",
			"iduser":   strconv.Itoa(block_of_data.Idbusiness),
			"typeuser": 1,
			"priority": 1,
			"title":    "⚠️ Alerta de Insumos 📦",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}

	return 201, false, "", "Enviado correctamente"
}

func Notify_ToEnd_Service() (int, bool, string, string) {

	ahora := time.Now()

	if ahora.Hour() == 14 {

		data_insumos, error_add := insumo_repository.Mo_Find_Notify_ToEnded()
		if error_add != nil {
			return 500, true, "Error en el servidor interno al intentar listar los insumos a notificar, detalles: " + error_add.Error(), ""
		}

		for _, block_of_data := range data_insumos {

			/*--SENT NOTIFICATION--*/
			notification := map[string]interface{}{
				"message":  "Cuenta con " + strconv.Itoa(block_of_data.Quantity) + " insumos con muy poco stock",
				"iduser":   strconv.Itoa(block_of_data.Idbusiness),
				"typeuser": 1,
				"priority": 1,
				"title":    "Restoner anfitriones",
			}
			json_data, _ := json.Marshal(notification)
			http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
			/*---------------------*/
		}

	}

	return 201, false, "", "Enviado correctamente"
}

/*---------------------------------------------------------------------*/

/*----------------------CREATE DATA OF INVENTARIO----------------------*/

func AddInsumo_Service(input_insumo models.Mo_Insumo) (int, bool, string, string) {

	error_add := insumo_repository.Mo_Add(input_insumo)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar agregar el insumo, detalles: " + error_add.Error(), ""
	}

	return 201, false, "", "Insumo creado correctamente"
}

func AddProvider_Service(input_provider models.Mo_Providers) (int, bool, string, string) {

	error_add := provider_repository.Mo_Add(input_provider)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar agregar el proveedor, detalles: " + error_add.Error(), ""
	}

	return 201, false, "", "Proveedor creado correctamente"
}

func AddStorehouse_Service(input_storehouse models.Mo_StoreHouse) (int, bool, string, string) {

	error_add := store_repository.Mo_Add(input_storehouse)
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar agregar el almacén, detalles: " + error_add.Error(), ""
	}

	return 201, false, "", "Almacén creado correctamente"
}

/*----------------------UPDATE MAIN DATA----------------------*/

func UpdateInsumo_MainData_Service(idbusiness int, input_insumo models.Mo_Insumo_Response) (int, bool, string, string) {

	error_update := insumo_repository.Mo_Update_MainData(idbusiness, input_insumo)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los datos del insumo, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Insumo actualizado correctamente"
}

func UpdateInsumo_Stock_Service(idinsumo string, idbusiness int, input_insumo models.Mo_Insumo_Stock_Adjust_Requst) (int, bool, string, string) {

	error_update := insumo_repository.Mo_Update_Stock(idinsumo, idbusiness, input_insumo)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el stock del insumo, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Stock de insumo actualizado correctamente"
}

func UpdateProvider_MainData_Service(idbusiness int, input_provider models.Mo_Providers_Response) (int, bool, string, string) {

	error_update := provider_repository.Mo_Update_MainData(idbusiness, input_provider)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los datos del proveedor, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Proveedor actualizado correctamente"
}

func UpdateStoreHouse_MainData_Service(idbusiness int, input_store models.Mo_StoreHouse_Response) (int, bool, string, string) {

	error_update := store_repository.Mo_Update_MainData(idbusiness, input_store)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los datos del almacén, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Almacén actualizado correctamente"
}

/*----------------------UPDATE AVAILABILITY----------------------*/

func UpdateInsumo_Availability_Service(idinsumo string, status bool, idbusiness int) (int, bool, string, string) {

	error_update := insumo_repository.Mo_Update_Availability(status, idinsumo, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el estado del insumo, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Estado Insumo actualizado correctamente"
}

func UpdateProvider_Availability_Service(idprovider string, status bool, idbusiness int) (int, bool, string, string) {

	error_update := provider_repository.Mo_Update_Availability(status, idprovider, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el estado del proveedor, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", " Estado Proveedor actualizado correctamente"
}

func UpdateStoreHouse_Availability_Service(idstorehouse string, status bool, idbusiness int) (int, bool, string, string) {

	error_update := store_repository.Mo_Update_Availability(status, idstorehouse, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el estado del almacén, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Estado Almacén actualizado correctamente"
}

/*----------------------UPDATE SEND TO DELET----------------------*/

func UpdateInsumo_SendToDelete_Service(idinsumo string, timezone int, idbusiness int) (int, bool, string, string) {

	error_update := insumo_repository.Mo_Update_SendToDelete(idinsumo, timezone, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar recuperar el insumo de la papelera, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Insumo enviado a papelera"
}

func UpdateProvider_SendToDelete_Service(idprovider string, timezone int, idbusiness int) (int, bool, string, string) {

	error_update := provider_repository.Mo_Update_SendToDelete(idprovider, timezone, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar recuperar el proveedor de la papelera, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Proveedor enviado a papelera"
}

func UpdateStoreHouse_SendToDelete_Service(idstorehouse string, timezone int, idbusiness int) (int, bool, string, string) {

	error_update := store_repository.Mo_Update_SendToDelete(idstorehouse, timezone, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar recuperar el almacén de la papelera, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Almacén enviado a papelera"
}

/*----------------------UPDATE RECOVER----------------------*/

func UpdateInsumo_RecoverSendToDelete_Service(idinsumo string, idbusiness int) (int, bool, string, string) {

	error_update := insumo_repository.Mo_Update_RecoverSendDelete(idinsumo, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar recuperar el insumo de la papelera, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Insumo recuperado de papelera"
}

func UpdateProvider_RecoverSendToDelete_Service(idprovider string, idbusiness int) (int, bool, string, string) {

	error_update := provider_repository.Mo_Update_RecoverSendDelete(idprovider, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar recuperar el proveedor de la papelera, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Proveedor recuperado de papelera"
}

func UpdateStoreHouse_RecoverSendToDelete_Service(idstorehouse string, idbusiness int) (int, bool, string, string) {

	error_update := store_repository.Mo_Update_RecoverSendDelete(idstorehouse, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar recuperar el almacén de la papelera, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Almacen recuperado de papelera"
}

/*----------------------DELETE----------------------*/

func UpdateInsumo_Delete_Service() (string, string) {

	error_update := insumo_repository.Mo_Update_Delete()
	if error_update != nil {
		return "Error en el servidor interno al intentar eliminar el insumo, detalles: " + error_update.Error(), ""
	}

	return "", "Insumos eliminados de papelera"
}

func UpdateProvider_Delete_Service() (string, string) {

	error_update := provider_repository.Mo_Update_Delete()
	if error_update != nil {
		return "Error en el servidor interno al intentar recuperar el proveedor de la papelera, detalles: " + error_update.Error(), ""
	}

	return "", "Proveedores eliminados de papelera"
}

func UpdateStoreHouse_Delete_Service() (string, string) {

	error_update := store_repository.Mo_Update_Delete()
	if error_update != nil {
		return "Error en el servidor interno al intentar recuperar el almacén de la papelera, detalles: " + error_update.Error(), ""
	}

	return "", "Almacenes  eliminados de papelera"
}

/*----------------------FIND DATA----------------------*/

func FindInsumo_All_Service(idbusiness int, limit int64, offset int64) (int, bool, string, []*models.Mo_Insumo_Response) {

	insumos, error_find := insumo_repository.Mo_Find_All(idbusiness, limit, offset)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar listar los insumos, detalles: " + error_find.Error(), insumos
	}

	return 201, false, "", insumos
}

func FindInsumo_Stock_Service(idbusiness int, idinsumo string) (int, bool, string, []*models.Mo_Stock) {

	stock, error_find := insumo_repository.Mo_Find_Stock(idinsumo, idbusiness)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar listar el stock, detalles: " + error_find.Error(), stock
	}

	return 201, false, "", stock
}

func FindProvider_All_Service(idbusiness int, limit int64, offset int64) (int, bool, string, []*models.Mo_Providers_Response) {

	providers, error_find := provider_repository.Mo_Find_All(idbusiness, limit, offset)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar listar los proveedores, detalles: " + error_find.Error(), providers
	}

	return 201, false, "", providers
}

func FindStorehouse_All_Service(idbusiness int, limit int64, offset int64) (int, bool, string, []*models.Mo_StoreHouse_Response) {

	storehouses, error_find := store_repository.Mo_Find_All(idbusiness, limit, offset)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar listar los almacenes, detalles: " + error_find.Error(), storehouses
	}

	return 201, false, "", storehouses
}

/*----------------------FIND PAPELERA DATA----------------------*/

func FindInsumo_Papelera_Service(idbusiness int) (int, bool, string, []*models.Mo_Insumo_Response) {

	insumos, error_find := insumo_repository.Mo_Find_Papelera(idbusiness)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar listar los insumos, detalles: " + error_find.Error(), insumos
	}

	return 201, false, "", insumos
}

func FindProvider_Papelera_Service(idbusiness int) (int, bool, string, []*models.Mo_Providers_Response) {

	providers, error_find := provider_repository.Mo_Find_Papelera(idbusiness)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar listar los proveedores, detalles: " + error_find.Error(), providers
	}

	return 201, false, "", providers
}

func FindStorehouse_Papelera_Service(idbusiness int) (int, bool, string, []*models.Mo_StoreHouse_Response) {

	storehouses, error_find := store_repository.Mo_Find_Papelera(idbusiness)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar listar los almacenes, detalles: " + error_find.Error(), storehouses
	}

	return 201, false, "", storehouses
}

/*----------------------SEARCH BY NAME----------------------*/

func SearchNameInsumo_Service(idbusiness int, name string) (int, bool, string, []*models.Mo_Insumo_Response) {

	insumos, error_find := insumo_repository.Mo_Search_Name(idbusiness, name)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar buscar el insumo, detalles: " + error_find.Error(), insumos
	}

	return 201, false, "", insumos
}

func SearchNameProvider_Service(idbusiness int, name string) (int, bool, string, []*models.Mo_Providers_Response) {

	proveedor, error_find := provider_repository.Mo_Search_Name(idbusiness, name)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar buscar el provedor, detalles: " + error_find.Error(), proveedor
	}

	return 201, false, "", proveedor
}

func SearchNameStorehouse_Service(idbusiness int, name string) (int, bool, string, []*models.Mo_StoreHouse_Response) {

	almacen, error_find := store_repository.Mo_Search_Name(idbusiness, name)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar buscar el almacén, detalles: " + error_find.Error(), almacen
	}

	return 201, false, "", almacen
}
