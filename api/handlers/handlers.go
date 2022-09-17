package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	carta "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/services/carta"
	exportfile "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/services/exportfile"
	imports "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/services/imports"
	inventario "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/services/inventario"
	pre_charged "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/services/pre_charged"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)

	//VERSION
	version_1 := e.Group("/v1")

	go Consumer_StadisticOrder()
	go Consumer_StockInsumos()

	//NOTIFICATIONS
	go Notify_ByScheduleRange()
	go Notify_InsumoStatus()

	//CLEAN TRASH
	go Clean_Data()
	/*====================FLUJO DE INFORMACIÓN====================*/

	/*===========INVENTARIO===========*/

	//V1 FROM V1 TO ...TO ENTITY PROVIDER
	router_provider := version_1.Group("/provider")
	router_provider.POST("", inventario.InventarioRouter_pg.AddProvider)
	router_provider.PUT("", inventario.InventarioRouter_pg.UpdateProvider_MainData)
	router_provider.PUT("/status/:idprovider/:status", inventario.InventarioRouter_pg.UpdateProvider_Availability)
	router_provider.PUT("/sendtrash/:idprovider/:timezone", inventario.InventarioRouter_pg.UpdateProvider_SendToDelete)
	router_provider.PUT("/recover/:idprovider", inventario.InventarioRouter_pg.UpdateProvider_RecoverSendToDelete)
	router_provider.GET("/:limit/:offset", inventario.InventarioRouter_pg.FindProvider_All)
	router_provider.GET("/trash", inventario.InventarioRouter_pg.FindProvider_Papelera)
	router_provider.GET("/search", inventario.InventarioRouter_pg.SearchNameProvider)

	//V1 FROM V1 TO ...TO ENTITY STOREHOUSE
	router_storehouse := version_1.Group("/storehouse")
	router_storehouse.POST("", inventario.InventarioRouter_pg.AddStorehouse)
	router_storehouse.PUT("", inventario.InventarioRouter_pg.UpdateStoreHouse_MainData)
	router_storehouse.PUT("/status/:idstorehouse/:status", inventario.InventarioRouter_pg.UpdateStoreHouse_Availability)
	router_storehouse.PUT("/sendtrash/:idstorehouse/:timezone", inventario.InventarioRouter_pg.UpdateStoreHouse_SendToDelete)
	router_storehouse.PUT("/recover/:idstorehouse", inventario.InventarioRouter_pg.UpdateStoreHouse_RecoverSendToDelete)
	router_storehouse.GET("/:limit/:offset", inventario.InventarioRouter_pg.FindStorehouse_All)
	router_storehouse.GET("/trash", inventario.InventarioRouter_pg.FindStorehouse_Papelera)
	router_storehouse.GET("/search", inventario.InventarioRouter_pg.SearchNameStorehouse)

	//V1 FROM V1 TO ...TO ENTITY INSUMO
	router_insumo := version_1.Group("/insumo")
	router_insumo.POST("", inventario.InventarioRouter_pg.AddInsumo)
	router_insumo.PUT("", inventario.InventarioRouter_pg.UpdateInsumo_MainData)
	router_insumo.PUT("/stock/:idinsumo", inventario.InventarioRouter_pg.UpdateInsumo_Stock)
	router_insumo.PUT("/status/:idinsumo/:status", inventario.InventarioRouter_pg.UpdateInsumo_Availability)
	router_insumo.PUT("/sendtrash/:idinsumo/:timezone", inventario.InventarioRouter_pg.UpdateInsumo_SendToDelete)
	router_insumo.PUT("/recover/:idinsumo", inventario.InventarioRouter_pg.UpdateInsumo_RecoverSendToDelete)
	router_insumo.GET("/:limit/:offset", inventario.InventarioRouter_pg.FindInsumo_All)
	router_insumo.GET("/stock/:idinsumo", inventario.InventarioRouter_pg.FindInsumo_Stock)
	router_insumo.GET("/trash", inventario.InventarioRouter_pg.FindInsumo_Papelera)
	router_insumo.GET("/search", inventario.InventarioRouter_pg.SearchNameInsumo)
	router_insumo.GET("/sendtoemail", exportfile.ExportfileRouter_pg.ExportFile_Insumo)

	/*===========CARTA===========*/

	//V1 FROM V1 TO ...TO ENTITY CATEGORY
	router_category := version_1.Group("/category")
	router_category.POST("", carta.CartaRouter_pg.AddCategory)
	router_category.GET("/status/:idcategory/elements", carta.CartaRouter_pg.GetElementsByCategory)
	router_category.PUT("/status/:idcategory/:status", carta.CartaRouter_pg.UpdateCategoryStatus)
	router_category.PUT("/sendtrash/:idcategory/:timezone", carta.CartaRouter_pg.SendToDeleteCategory)
	router_category.PUT("/recover/:idcategory", carta.CartaRouter_pg.RecoverSendToDeleteCategory)
	router_category.POST("/image", carta.CartaRouter_pg.UpdateCategory_Consumer)
	router_category.GET("/all", carta.CartaRouter_pg.FindAllCategories)
	router_category.GET("/trash", carta.CartaRouter_pg.FindCategory_Papelera)

	//V1 FROM V1 TO ...TO ENTITY ELEMENT
	router_element := version_1.Group("/element")
	router_element.POST("", carta.CartaRouter_pg.AddElement)
	router_element.PUT("", carta.CartaRouter_pg.UpdateElement)
	router_element.PUT("/status/:idelement/:status", carta.CartaRouter_pg.UpdateElementStatus)
	router_element.PUT("/sendtrash/:idelement/:timezone", carta.CartaRouter_pg.SendToDeleteElement)
	router_element.PUT("/recover/:idelement", carta.CartaRouter_pg.RecoverSendToDeleteElement)
	router_element.POST("/image", carta.CartaRouter_pg.UpdateElement_Consumer)
	router_element.GET("/:limit/:offset", carta.CartaRouter_pg.FindAllElements)
	router_element.GET("/rating/:day/:limit/:offset", carta.CartaRouter_pg.FindElementsRatingByDay)
	router_element.GET("/search", carta.CartaRouter_pg.FindElementsRatingByName)
	router_element.GET("/precharged", pre_charged.ImportsRouter_pg.FindPreCharged)
	router_element.GET("/trash", carta.CartaRouter_pg.FindElement_Papelera)
	router_element.GET("/sendtoemail", exportfile.ExportfileRouter_pg.ExportFile_Element)

	//V1 FROM V1 TO ...TO ENTITY SCHEDULE RANGE
	router_schedule_range := version_1.Group("/schedulerange")
	router_schedule_range.POST("", carta.CartaRouter_pg.AddScheduleRange)
	router_schedule_range.PUT("", carta.CartaRouter_pg.UpdateScheduleRange)
	router_schedule_range.DELETE("/:idschedulerange", carta.CartaRouter_pg.UpdateScheduleRangeStatus)
	router_schedule_range.GET("", carta.CartaRouter_pg.FindAllRangoHorario)

	//V1 FROM V1 TO ...TO ENTITY PRECHARGED
	router_precharged := version_1.Group("/pre-charged-element")
	router_precharged.POST("", pre_charged.ImportsRouter_pg.AddPreCharged_Multiple)

	//V1 FROM V1 TO ...TO ENTITY TOTAL VALUES INVENTARIO
	router_total_data := version_1.Group("/totalinventario")
	router_total_data.GET("", carta.CartaRouter_pg.FindAllCarta_MainData)

	//Abrimos el puerto
	PORT := os.Getenv("PORT")
	//Si dice que existe PORT
	if PORT == "" {
		PORT = "6500"
	}

	//cors son los permisos que se le da a la API
	//para que sea accesibl esde cualquier lugar
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func index(c echo.Context) error {
	return c.JSON(401, "Acceso no autorizado")
}

func Consumer_StadisticOrder() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {

		log.Fatal("Error connection canal")
	}

	msgs, err_consume := ch.Consume("anfitrion/ordersperelement", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola")
	}

	noStop3 := make(chan bool)
	go func() {
		for d := range msgs {
			var export_byelement []models.Pg_Import_StadisticOrders
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&export_byelement)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			carta.CartaRouter_pg.Import_OrderStadistic(export_byelement)
		}
	}()

	<-noStop3
}

func Consumer_StockInsumos() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {

		log.Fatal("Error connection canal")
	}

	msgs, err_consume := ch.Consume("anfitrion/stock_insumo_add", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola")
	}

	noStopInsumoStatus := make(chan bool)
	go func() {
		for d := range msgs {
			var import_elements []models.Mqtt_Import_InsumoStock
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&import_elements)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			imports.ImportsRouter_pg.UpdateElementStock(import_elements)
		}
	}()

	<-noStopInsumoStatus
}

//NOTIFICATION

func Notify_ByScheduleRange() {

	noStop_NotifySchedule := make(chan bool)
	go func() {
		for {
			time.Sleep(48 * time.Hour)
			carta.CartaRouter_pg.SearchToNotifySchedulerange()
		}
	}()

	<-noStop_NotifySchedule
}

func Notify_InsumoStatus() {

	noStop_NotifySchedule := make(chan bool)
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			inventario.InventarioRouter_pg.Notify_Ended()
			inventario.InventarioRouter_pg.Notify_ToEnd()
		}
	}()

	<-noStop_NotifySchedule
}

//CLEAN DATA

func Clean_Data() {

	noStop_Clean_Data := make(chan bool)
	go func() {
		for {
			time.Sleep(2 * time.Hour)
			//Limpiar categorias
			carta.CartaRouter_pg.UpdateCategory_Delete()
			time.Sleep(5 * time.Minute)
			//Limpiar elementos
			carta.CartaRouter_pg.UpdateElement_Delete()
			time.Sleep(10 * time.Hour)
			//Limpiar proveedores
			inventario.InventarioRouter_pg.UpdateProvider_Delete()
			time.Sleep(15 * time.Hour)
			//Limpiar almacenes
			inventario.InventarioRouter_pg.UpdateStoreHouse_Delete()
			time.Sleep(20 * time.Hour)
			//Limpiar insumos
			inventario.InventarioRouter_pg.UpdateInsumo_Delete()
		}
	}()
	<-noStop_Clean_Data
}
