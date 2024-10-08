package inventario

import "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type Response_Notify struct {
	Error     bool          `json:"error"`
	DataError string        `json:"dataError"`
	Data      []interface{} `json:"data"`
}

type Response_Notify_test struct {
	Error     bool                   `json:"error"`
	DataError string                 `json:"dataError"`
	Data      []models.Mo_NotifyData `json:"data"`
}

type JWT struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type ResponseJWT struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT    `json:"data"`
}

type Response_Providers struct {
	Error     bool                            `json:"error"`
	DataError string                          `json:"dataError"`
	Data      []*models.Mo_Providers_Response `json:"data"`
}

type Response_StoreHouse struct {
	Error     bool                             `json:"error"`
	DataError string                           `json:"dataError"`
	Data      []*models.Mo_StoreHouse_Response `json:"data"`
}

type Response_Insumo struct {
	Error     bool                         `json:"error"`
	DataError string                       `json:"dataError"`
	Data      []*models.Mo_Insumo_Response `json:"data"`
}

type Response_Stock struct {
	Error     bool               `json:"error"`
	DataError string             `json:"dataError"`
	Data      []*models.Mo_Stock `json:"data"`
}

//NOTIFY

type Mo_Notify_Insumo struct {
	Message    string `bson:"message" json:"message,omitempty"`
	IDUser     int    `bson:"iduser" json:"iduser,omitempty"`
	Priority   int    `bson:"priority" json:"priority"`
	TypeUser   int    `bson:"typeuser" json:"typeuser,omitempty"`
	Title      string `bson:"title" json:"title,omitempty"`
	CodeNotify int    `bson:"codenotify" json:"codenotify,omitempty"`
}

type Response_Notify_Insumo struct {
	Error     bool               `json:"error"`
	DataError string             `json:"dataError"`
	Data      []Mo_Notify_Insumo `json:"data"`
}
