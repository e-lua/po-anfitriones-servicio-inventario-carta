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
	FromCarta string `json:"fromcarta"`
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
	ScheduleRanges []models.Pg_ScheduleRange_External `json:"schedule"`
}

type ResponseCartaBasicData struct {
	Error     bool                     `json:"error"`
	DataError string                   `json:"dataError"`
	Data      models.Pg_Carta_External `json:"data"`
}

type ResponseCartaCategory struct {
	Error     bool                          `json:"error"`
	DataError string                        `json:"dataError"`
	Data      []models.Pg_Category_External `json:"data"`
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

type ResponseCartas struct {
	Error     bool                    `json:"error"`
	DataError string                  `json:"dataError"`
	Data      []models.Pg_Carta_Found `json:"data"`
}

type ResponseCartaCategoryAndElement struct {
	Error     bool                    `json:"error"`
	DataError string                  `json:"dataError"`
	Data      CartaCategoryAndElement `json:"data"`
}

type CartaCategoryAndElement struct {
	IDCarta    int                                     `json:"icarta"`
	Categories []models.Pg_Category_External           `json:"categories"`
	Elements   []models.Pg_Element_With_Stock_External `json:"elements"`
}

/*===============TESTEANDO===============*/

type CartaElements_WithAction struct {
	IDCarta            int                                            `json:"idcarta"`
	ElementsWithAction []models.Pg_Element_With_Stock_External_Action `json:"elements"`
}
