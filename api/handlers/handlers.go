package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"
	carta "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/services/flujo_de_informacion"
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

	/*====================FLUJO DE INFORMACIÓN====================*/

	/*===========CARTA===========*/

	//V1 FROM V1 TO ...TO ENTITY CATEGORY
	router_category := version_1.Group("/category")
	router_category.POST("", carta.CartaRouter_pg.AddCategory)
	router_category.PUT("", carta.CartaRouter_pg.UpdateCategory)
	router_category.GET("", carta.CartaRouter_pg.FindAllCategories)

	//V1 FROM V1 TO ...TO ENTITY ELEMENT
	router_element := version_1.Group("/element")
	router_element.POST("", carta.CartaRouter_pg.AddElement)
	router_element.PUT("", carta.CartaRouter_pg.UpdateElement)
	router_element.GET("/:limit/:offset", carta.CartaRouter_pg.FindAllElements)

	//V1 FROM V1 TO ...TO ENTITY ELEMENT
	router_schedule_range := version_1.Group("/schedulerange")
	router_schedule_range.POST("", carta.CartaRouter_pg.AddScheduleRange)
	router_schedule_range.PUT("", carta.CartaRouter_pg.UpdateScheduleRange)
	router_schedule_range.GET("", carta.CartaRouter_pg.FindAllRangoHorario)

	//V1 FROM V1 TO ...TO ENTITY ELEMENT
	router_total_data := version_1.Group("/totalcartvalues")
	router_total_data.GET("", carta.CartaRouter_pg.FindAllCarta_MainData)

	//V1 FROM V1 TO ...TO ENTITY CARTA
	router_carta := version_1.Group("/carta")
	router_carta.POST("", carta.CartaRouter_pg.AddCarta)

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
		defer ch.Close()
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
		defer ch.Close()
		log.Fatal("Error connection canal")
	}

	msgs, err_consume := ch.Consume("anfitrion/element", "", true, false, false, false, nil)
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
			carta.CartaRouter_pg.UpdateElement_Consumer(toCarta.IdBanner_Category_Element, toCarta.Url, toCarta.IdBusiness)
		}
	}()

	<-noStop
}
