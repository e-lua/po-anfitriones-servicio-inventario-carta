package models

import "time"

type Mqtt_Element_With_Stock struct {
	IDElement        int       `json:"id"`
	IDCarta          int       `json:"idcarta"`
	IDBusiness       int       `json:"idbusiness"`
	IDCategory       int       `json:"idcategory"`
	Typefood         string    `json:"typefood"`
	NameCategory     string    `json:"namecategory"`
	UrlPhotoCategory string    `json:"urlphotocategory"`
	Name             string    `json:"name"`
	Price            float32   `json:"price"`
	Description      string    `json:"description"`
	TypeMoney        int       `json:"typemoney"`
	Stock            int       `json:"stock"`
	UrlPhoto         string    `json:"url"`
	AvailableOrders  bool      `json:"isavailableorders"`
	IsExported       bool      `json:"isexported"`
	Date             string    `json:"date"`
	DeletedDate      time.Time `json:"deleteddate"`
}

type Mqtt_Element_With_Stock_export struct {
	IdCarta             int           `json:"idcarta"`
	Elements_with_stock []interface{} `json:"elementswithstock"`
}
