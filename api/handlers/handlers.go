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

	/*====================FLUJO DE INFORMACIÓN====================*/

	/*===========CARTA===========*/

	//V1 FROM V1 TO ...TO ENTITY CATEGORY
	router_category := version_1.Group("/category")
	router_category.POST("", inventario.InvetarioRouter_pg.AddCategory)
	router_category.GET("/status/:idcategory/elements", inventario.InvetarioRouter_pg.GetElementsByCategory)
	router_category.PUT("/status/:idcategory/:status", inventario.InvetarioRouter_pg.UpdateCategoryStatus)
	router_category.GET("/all", inventario.InvetarioRouter_pg.FindAllCategories)

	//V1 FROM V1 TO ...TO ENTITY ELEMENT
	router_element := version_1.Group("/element")
	router_element.POST("", inventario.InvetarioRouter_pg.AddElement)
	router_element.PUT("", inventario.InvetarioRouter_pg.UpdateElement)
	router_element.PUT("/status/:idelement/:status", inventario.InvetarioRouter_pg.UpdateElementStatus)
	router_element.GET("/:limit/:offset", inventario.InvetarioRouter_pg.FindAllElements)
	router_element.GET("/rating/:day/:limit/:offset", inventario.InvetarioRouter_pg.FindElementsRatingByDay)
	router_element.GET("/search", inventario.InvetarioRouter_pg.FindElementsRatingByName)

	//V1 FROM V1 TO ...TO ENTITY SCHEDULE RANGE
	router_schedule_range := version_1.Group("/schedulerange")
	router_schedule_range.POST("", inventario.InvetarioRouter_pg.AddScheduleRange)
	router_schedule_range.PUT("", inventario.InvetarioRouter_pg.UpdateScheduleRange)
	router_schedule_range.DELETE("/:idschedulerange", inventario.InvetarioRouter_pg.UpdateScheduleRangeStatus)
	router_schedule_range.GET("", inventario.InvetarioRouter_pg.FindAllRangoHorario)

	//V1 FROM V1 TO ...TO ENTITY TOTAL VALUES INVENTARIO
	router_total_data := version_1.Group("/totalinventario")
	router_total_data.GET("", inventario.InvetarioRouter_pg.FindAllCarta_MainData)

	/*===========CARTA DIARIA===========*/

	router_menu := version_1.Group("/menu")
	router_menu.POST("", carta.CartaRouter_pg.AddCarta)
	router_menu.PUT("", carta.CartaRouter_pg.UpdateCartaStatus)
	router_menu.GET("", carta.CartaRouter_pg.GetCartas)
	router_menu.DELETE("", carta.CartaRouter_pg.DeleteCarta)
	router_menu.GET("/:date", carta.CartaRouter_pg.GetCartaBasicData)
	router_menu.GET("/:idcarta/category", carta.CartaRouter_pg.GetCartaCategory)
	router_menu.GET("/:idcarta/category/:idcategory/elements", carta.CartaRouter_pg.GetCartaElementsByCarta)
	router_menu.PUT("/elements", carta.CartaRouter_pg.UpdateCartaElements)
	router_menu.GET("/:idcarta/elements", carta.CartaRouter_pg.GetCartaElements)
	router_menu.PUT("/onelement", carta.CartaRouter_pg.UpdateCartaOneElement)
	router_menu.PUT("/scheduleranges", carta.CartaRouter_pg.UpdateCartaScheduleRanges)
	router_menu.GET("/:idcarta/scheduleranges", carta.CartaRouter_pg.GetCartaScheduleRanges)
	/*to create an order*/
	router_menu.GET("/createorder/:date/category", carta.CartaRouter_pg.GetCategories_ToCreateOrder)
	router_menu.GET("/createorder/:date/category/:idcategory/elements", carta.CartaRouter_pg.GetElements_ToCreateOrder)
	router_menu.GET("/createorder/:date/scheduleranges", carta.CartaRouter_pg.GetSchedule_ToCreateOrder)

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
			inventario.InvetarioRouter_pg.UpdateCategory_Consumer(toCarta.IdBanner_Category_Element, toCarta.Url, toCarta.IdBusiness)
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
			inventario.InvetarioRouter_pg.UpdateElement_Consumer(toCarta.IdBanner_Category_Element, toCarta.Url, toCarta.IdBusiness)
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
			inventario.InvetarioRouter_pg.Import_OrderStadistic(export_byelement)
		}
	}()

	<-noStop3
}

func Notify_ByScheduleRange() {
	for {
		time.Sleep(48 * time.Hour)
		inventario.InvetarioRouter_pg.SearchToNotifySchedulerange()
	}
}

func Notify_ByCarta() {
	for {
		time.Sleep(24 * time.Hour)
		carta.CartaRouter_pg.SearchToNotifyCarta()
	}
}
