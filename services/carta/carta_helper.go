package carta

import "time"

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
	Date      time.Time `json:"date"`
	WannaCopy bool      `json:"wannacopy"`
	FromCarta int       `json:"fromcarta"`
}

type CartaStatus struct {
	IDCarta   int  `json:"idcarta"`
	Available bool `json:"available"`
	Visible   bool `json:"visible"`
}
