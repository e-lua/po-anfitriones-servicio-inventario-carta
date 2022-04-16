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
	cartadiaria "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/services/cartadiaria"
	inventario "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/services/inventario"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)

	//VERSION
	version_1 := e.Group("/v1")

	go Consumer_Category()
	go Consumer_Element()
	go Consumer_StadisticOrder()
	go Notify_ByScheduleRange()
	go Notify_ByCarta()

	//CLEAN TRASH
	go Clean_Categories()
	go Clean_Elements()
	go Clean_Providers()
	go Clean_StoreHouses()
	go Clean_Insumos()

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
	router_storehouse.GET("/:limit/:offset", inventario.InventarioRouter_pg.FindProvider_All)
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
	router_insumo.GET("/stock", inventario.InventarioRouter_pg.FindInsumo_Stock)
	router_insumo.GET("/trash", inventario.InventarioRouter_pg.FindInsumo_Papelera)
	router_insumo.GET("/search", inventario.InventarioRouter_pg.SearchNameInsumo)

	/*===========CARTA===========*/

	//V1 FROM V1 TO ...TO ENTITY CATEGORY
	router_category := version_1.Group("/category")
	router_category.POST("", carta.CartaRouter_pg.AddCategory)
	router_category.GET("/status/:idcategory/elements", carta.CartaRouter_pg.GetElementsByCategory)
	router_category.PUT("/status/:idcategory/:status", carta.CartaRouter_pg.UpdateCategoryStatus)
	router_category.PUT("/sendtrash/:idcategory/:timezone", carta.CartaRouter_pg.SendToDeleteCategory)
	router_category.PUT("/recover/:idcategory", carta.CartaRouter_pg.RecoverSendToDeleteCategory)
	router_category.GET("/all", carta.CartaRouter_pg.FindAllCategories)

	//V1 FROM V1 TO ...TO ENTITY ELEMENT
	router_element := version_1.Group("/element")
	router_element.POST("", carta.CartaRouter_pg.AddElement)
	router_element.PUT("", carta.CartaRouter_pg.UpdateElement)
	router_element.PUT("/status/:idelement/:status", carta.CartaRouter_pg.UpdateElementStatus)
	router_element.PUT("/sendtrash/:idelement/:timezone", carta.CartaRouter_pg.SendToDeleteElement)
	router_element.PUT("/recover/:idelement", carta.CartaRouter_pg.RecoverSendToDeleteElement)
	router_element.GET("/:limit/:offset", carta.CartaRouter_pg.FindAllElements)
	router_element.GET("/rating/:day/:limit/:offset", carta.CartaRouter_pg.FindElementsRatingByDay)
	router_element.GET("/search", carta.CartaRouter_pg.FindElementsRatingByName)

	//V1 FROM V1 TO ...TO ENTITY SCHEDULE RANGE
	router_schedule_range := version_1.Group("/schedulerange")
	router_schedule_range.POST("", carta.CartaRouter_pg.AddScheduleRange)
	router_schedule_range.PUT("", carta.CartaRouter_pg.UpdateScheduleRange)
	router_schedule_range.DELETE("/:idschedulerange", carta.CartaRouter_pg.UpdateScheduleRangeStatus)
	router_schedule_range.GET("", carta.CartaRouter_pg.FindAllRangoHorario)

	//V1 FROM V1 TO ...TO ENTITY TOTAL VALUES INVENTARIO
	router_total_data := version_1.Group("/totalinventario")
	router_total_data.GET("", carta.CartaRouter_pg.FindAllCarta_MainData)

	/*===========CARTA DIARIA===========*/

	router_menu := version_1.Group("/menu")
	router_menu.POST("", cartadiaria.CartaDiariaRouter_pg.AddCarta)
	router_menu.PUT("", cartadiaria.CartaDiariaRouter_pg.UpdateCartaStatus)
	router_menu.GET("", cartadiaria.CartaDiariaRouter_pg.GetCartas)
	router_menu.DELETE("", cartadiaria.CartaDiariaRouter_pg.DeleteCarta)
	router_menu.GET("/:date", cartadiaria.CartaDiariaRouter_pg.GetCartaBasicData)
	router_menu.GET("/:idcarta/category", cartadiaria.CartaDiariaRouter_pg.GetCartaCategory)
	router_menu.GET("/:idcarta/category/:idcategory/elements", cartadiaria.CartaDiariaRouter_pg.GetCartaElementsByCarta)
	router_menu.PUT("/elements", cartadiaria.CartaDiariaRouter_pg.UpdateCartaElements)
	router_menu.GET("/:idcarta/elements", cartadiaria.CartaDiariaRouter_pg.GetCartaElements)
	router_menu.PUT("/onelement", cartadiaria.CartaDiariaRouter_pg.UpdateCartaOneElement)
	router_menu.PUT("/scheduleranges", cartadiaria.CartaDiariaRouter_pg.UpdateCartaScheduleRanges)
	router_menu.GET("/:idcarta/scheduleranges", cartadiaria.CartaDiariaRouter_pg.GetCartaScheduleRanges)
	/*to create an order*/
	router_menu.GET("/createorder/:date/category", cartadiaria.CartaDiariaRouter_pg.GetCategories_ToCreateOrder)
	router_menu.GET("/createorder/:date/category/:idcategory/elements", cartadiaria.CartaDiariaRouter_pg.GetElements_ToCreateOrder)
	router_menu.GET("/createorder/:date/scheduleranges", cartadiaria.CartaDiariaRouter_pg.GetSchedule_ToCreateOrder)

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

func Consumer_Category() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {

		log.Fatal("Error connection canal")
	}

	msgs, err_consume := ch.Consume("anfitrion/category", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola")
	}

	noStop := make(chan bool)
	go func() {
		for d := range msgs {
			var toCarta models.Pg_ToCarta_Mqtt
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&toCarta)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			carta.CartaRouter_pg.UpdateCategory_Consumer(toCarta.IdBanner_Category_Element, toCarta.Url, toCarta.IdBusiness)
		}
	}()

	<-noStop
}

func Consumer_Element() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {

		log.Fatal("Error connection canal")
	}

	msgs, err_consume := ch.Consume("anfitrion/element", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola")
	}

	noStop2 := make(chan bool)
	go func() {
		for d := range msgs {
			var toCarta models.Pg_ToCarta_Mqtt
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&toCarta)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			carta.CartaRouter_pg.UpdateElement_Consumer(toCarta.IdBanner_Category_Element, toCarta.Url, toCarta.IdBusiness)
		}
	}()

	<-noStop2
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

func Notify_ByScheduleRange() {
	for {
		carta.CartaRouter_pg.SearchToNotifySchedulerange()
		time.Sleep(48 * time.Hour)
	}
}

func Notify_ByCarta() {
	for {
		cartadiaria.CartaDiariaRouter_pg.SearchToNotifyCarta()
		time.Sleep(24 * time.Hour)
	}
}

//CLEAN DATA

func Clean_Categories() {
	for {
		carta.CartaRouter_pg.UpdateCategory_Delete()
		time.Sleep(24 * time.Hour)
	}
}

func Clean_Elements() {
	for {
		log.Println("Testing deleting Elements")
		carta.CartaRouter_pg.UpdateElement_Delete()
		time.Sleep(24 * time.Hour)
	}
}

func Clean_Providers() {
	for {
		inventario.InventarioRouter_pg.UpdateProvider_Delete()
		time.Sleep(24 * time.Hour)
	}
}

func Clean_StoreHouses() {
	for {
		inventario.InventarioRouter_pg.UpdateStoreHouse_Delete()
		time.Sleep(24 * time.Hour)
	}
}

func Clean_Insumos() {
	for {
		inventario.InventarioRouter_pg.UpdateInsumo_Delete()
		time.Sleep(24 * time.Hour)
	}
}
