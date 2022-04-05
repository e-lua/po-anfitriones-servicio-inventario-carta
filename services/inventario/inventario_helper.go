package inventario

import "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type ResponseElementsByCategory struct {
	Error     bool               `json:"error"`
	DataError string             `json:"dataError"`
	Data      ElementsByCategory `json:"data"`
}

type ElementsByCategory struct {
	Element  []models.Pg_ElementsByCategory `json:"elements"`
	Quantity int                            `json:"quantity"`
}

type ResponseInt struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      int    `json:"data"`
}

type ResponseJWT struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT    `json:"data"`
}

type ResponseListCategory struct {
	Error     bool                 `json:"error"`
	DataError string               `json:"dataError"`
	Data      []models.Pg_Category `json:"data"`
}

type ResponseListCategoryForSearch struct {
	Error     bool                            `json:"error"`
	DataError string                          `json:"dataError"`
	Data      []models.Pg_Category_for_search `json:"data"`
}

type ResponseCategory struct {
	Error     bool               `json:"error"`
	DataError string             `json:"dataError"`
	Data      models.Pg_Category `json:"data"`
}

type Category struct {
	IdBusiness int    `json:"id"`
	Name       string `json:"name"`
	UrlPhoto   string `json:"url"`
}

type ResponseListElement struct {
	Error     bool                `json:"error"`
	DataError string              `json:"dataError"`
	Data      []models.Pg_Element `json:"data"`
}

type ResponseListElement_WithRating struct {
	Error     bool                           `json:"error"`
	DataError string                         `json:"dataError"`
	Data      []models.Pg_Element_WithRating `json:"data"`
}

type Response_StadisticElement struct {
	Error     bool          `json:"error"`
	DataError string        `json:"dataError"`
	Data      []interface{} `json:"data"`
}

type ResponseListElementForSearch struct {
	Error     bool                           `json:"error"`
	DataError string                         `json:"dataError"`
	Data      []models.Pg_Element_for_search `json:"data"`
}

type ResponseElement struct {
	Error     bool              `json:"error"`
	DataError string            `json:"dataError"`
	Data      models.Pg_Element `json:"data"`
}

type Element struct {
	IdElement int     `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	UrlPhoto  string  `json:"url"`
}

type ResponseListRangoHorario struct {
	Error     bool                      `json:"error"`
	DataError string                    `json:"dataError"`
	Data      []models.Pg_ScheduleRange `json:"data"`
}

type ResponseAllMainData struct {
	Error     bool                                     `json:"error"`
	DataError string                                   `json:"dataError"`
	Data      models.Pg_Category_Element_ScheduleRange `json:"data"`
}

type ResponseRangoHorario struct {
	Error     bool                    `json:"error"`
	DataError string                  `json:"dataError"`
	Data      models.Pg_ScheduleRange `json:"data"`
}

type RangoHorario struct {
	IdRangoHorario    int    `json:"id"`
	Name              string `json:"name"`
	MinutePerFraction int    `json:"minutesPerFraction"`
	StartTime         string `json:"startTIme"`
	EndTime           string `json:"endTime"`
	MaxOrders         int    `json:"maxOrders"`
}

type JWT struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type CategoryForSearch struct {
	Name string `json:"name"`
}

//Notify

type ResponseAllBusinesses struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      []int  `json:"data"`
}
