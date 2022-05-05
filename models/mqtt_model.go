package models

import "time"

type Mqtt_Element_With_Stock struct {
	IDElement        int                     `json:"id"`
	IDCarta          int                     `json:"idcarta"`
	IDBusiness       int                     `json:"idbusiness"`
	IDCategory       int                     `json:"idcategory"`
	Typefood         string                  `json:"typefood"`
	NameCategory     string                  `json:"namecategory"`
	UrlPhotoCategory string                  `json:"urlphotocategory"`
	Name             string                  `json:"name"`
	Price            float32                 `json:"price"`
	Description      string                  `json:"description"`
	TypeMoney        int                     `json:"typemoney"`
	Stock            int                     `json:"stock"`
	UrlPhoto         string                  `json:"url"`
	AvailableOrders  bool                    `json:"isavailableorders"`
	IsExported       bool                    `json:"isexported"`
	Date             string                  `json:"date"`
	DeletedDate      time.Time               `json:"deleteddate"`
	Insumos          []Pg_Mo_Insumo_Elements `bson:"insumos"  json:"insumos"`
	Costo            float64                 `json:"costo"`
}

type Mqtt_Element_With_Stock_export struct {
	IdCarta             int           `json:"idcarta"`
	Elements_with_stock []interface{} `json:"elementswithstock"`
}

type Mqtt_Import_InsumoStock struct {
	Quantity int                     `json:"quantity"`
	Insumos  []Pg_Mo_Insumo_Elements `json:"insumos"`
}

//TESTING

type Mqtt_Insumo_Elements struct {
	Name           string      `json:"name"`
	Measure        string      `json:"measure"`
	IDStoreHouse   string      `json:"idstorehouse"`
	NameStoreHouse string      `json:"namestorehouse"`
	Description    string      `json:"description"`
	Stock          []*Mo_Stock `json:"stock"`
	Quantity       int         `json:"quantity"`
}

type Mqtt_Element struct {
	IDElement        int                    `json:"id"`
	IDCategory       int                    `json:"idcategory"`
	NameCategory     string                 `json:"namecategory"`
	URLPhotoCategory string                 `json:"urlphotocategory"`
	Typefood         string                 `json:"typefood"`
	Name             string                 `json:"name"`
	Price            float32                `json:"price"`
	Description      string                 `json:"description"`
	TypeMoney        int                    `json:"typemoney"`
	UrlPhoto         interface{}            `json:"url"`
	Available        bool                   `json:"available"`
	Insumos          []Mqtt_Insumo_Elements `json:"insumos"`
	SendToDelete     time.Time              `json:"sendtodelete"`
	IsDelete         bool                   `json:"isdeleted"`
	IsExported       bool                   `json:"isexported"`
	DeletedDate      time.Time              `json:"deleteddate"`
	IsSendToDelete   bool                   `json:"issendtodelete"`
	Costo            float64                `json:"costo"`
}
