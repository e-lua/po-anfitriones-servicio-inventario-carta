package carta

import "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
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

type JWT struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type ResponseObjectId struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type Carta struct {
	Date      string `json:"date"`
	WannaCopy bool   `json:"wannacopy"`
	FromCarta int    `json:"fromcarta"`
}

type CartaStatus struct {
	IDCarta   int  `json:"idcarta"`
	Available bool `json:"available"`
	Visible   bool `json:"visible"`
}

type CartaElements struct {
	IDCarta  int                                     `json:"idcarta"`
	Elements []models.Pg_Element_With_Stock_External `json:"elements"`
}

type CartaSchedule struct {
	IDCarta        int                                `json:"idcarta"`
	ScheduleRanges []models.Pg_ScheduleRange_External `json:"elements"`
}

type ResponseCartaBasicData struct {
	Error     bool                     `json:"error"`
	DataError string                   `json:"dataError"`
	Data      models.Pg_Carta_External `json:"data"`
}

type ResponseCartaElements struct {
	Error     bool                                    `json:"error"`
	DataError string                                  `json:"dataError"`
	Data      []models.Pg_Element_With_Stock_External `json:"data"`
}

type ResponseCartaSchedule struct {
	Error     bool                               `json:"error"`
	DataError string                             `json:"dataError"`
	Data      []models.Pg_ScheduleRange_External `json:"data"`
}
