package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pg_Category struct {
	IDCategory     int       `json:"id"`
	Name           string    `json:"name"`
	Elements       int       `json:"elements"`
	Available      bool      `json:"available"`
	UrlPhoto       string    `json:"url"`
	TypeFood       string    `json:"typefood"`
	SendToDelete   time.Time `json:"sendtodelete"`
	IsDelete       bool      `json:"isdeleted"`
	IsExported     bool      `json:"isexported"`
	DeletedDate    time.Time `json:"deleteddate"`
	IsSendToDelete bool      `json:"issendtodelete"`
}

type Pg_Category_Response struct {
	IDCategory   int         `json:"id"`
	Name         string      `json:"name"`
	Elements     int         `json:"elements"`
	Available    bool        `json:"available"`
	UrlPhoto     interface{} `json:"url"`
	TypeFood     string      `json:"typefood"`
	SendToDelete time.Time   `json:"sendtodelete"`
	DeletedDate  time.Time   `json:"deleteddate"`
}

type Pg_Category_ToCreate struct {
	IDCarta          int    `json:"idcarta"`
	IDCategory       int    `json:"idcategory"`
	Name             string `json:"namecategory"`
	UrlPhoto         string `json:"urlphotocategory"`
	AmountOfElements int    `json:"elements"`
}

type Pg_Category_for_search struct {
	IDCategory int    `json:"id"`
	Name       string `json:"name"`
}

type Pg_Element_for_search struct {
	IDElement int     `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	TypeMoney int     `json:"typeMoney"`
}

type Pg_Element struct {
	IDElement        int                     `json:"id"`
	IDCategory       int                     `json:"idcategory"`
	NameCategory     string                  `json:"namecategory"`
	URLPhotoCategory string                  `json:"urlphotocategory"`
	Typefood         string                  `json:"typefood"`
	Name             string                  `json:"name"`
	Price            float32                 `json:"price"`
	Description      string                  `json:"description"`
	TypeMoney        int                     `json:"typemoney"`
	UrlPhoto         interface{}             `json:"url"`
	Available        bool                    `json:"available"`
	Insumos          []Pg_Mo_Insumo_Elements `json:"insumos"`
	SendToDelete     time.Time               `json:"sendtodelete"`
	IsDelete         bool                    `json:"isdeleted"`
	IsExported       bool                    `json:"isexported"`
	DeletedDate      time.Time               `json:"deleteddate"`
	IsSendToDelete   bool                    `json:"issendtodelete"`
	Costo            float64                 `json:"costo"`
	IsAutomaticCost  bool                    `json:"isautomaticcost"`
	IsURLPrecharged  bool                    `json:"isurlprecharged"`
}

type Pg_Mo_Insumo_Elements struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name"`
	Measure        string             `json:"measure"`
	IDStoreHouse   string             `json:"idstorehouse"`
	NameStoreHouse string             `json:"namestorehouse"`
	Description    string             `json:"description"`
	Stock          []*Mo_Stock        `json:"stock"`
	Quantity       int                `json:"quantity"`
}

type Pg_Element_Response struct {
	IDElement        int                     `json:"id"`
	IDCategory       int                     `json:"idcategory"`
	NameCategory     string                  `json:"namecategory"`
	URLPhotoCategory string                  `json:"urlphotocategory"`
	Typefood         string                  `json:"typefood"`
	Name             string                  `json:"name"`
	Price            float32                 `json:"price"`
	Description      string                  `json:"description"`
	TypeMoney        int                     `json:"typemoney"`
	UrlPhoto         interface{}             `json:"url"`
	Available        bool                    `json:"available"`
	Insumos          []Pg_Mo_Insumo_Elements `json:"insumos"`
	SendToDelete     time.Time               `json:"sendtodelete"`
	IsSendToDelete   bool                    `json:"issendtodelete"`
	IsAutomaticCost  bool                    `json:"isautomaticcost"`
	Costo            float64                 `json:"costo"`
	IsURLPrecharged  bool                    `json:"isurlprecharged"`
}

type Pg_Element_Tofind struct {
	IDElement        int                     `json:"id"`
	IDCategory       int                     `json:"idcategory"`
	NameCategory     string                  `json:"namecategory"`
	URLPhotoCategory string                  `json:"urlphotocategory"`
	Typefood         string                  `json:"typefood"`
	Name             string                  `json:"name"`
	Price            float32                 `json:"price"`
	Description      string                  `json:"description"`
	TypeMoney        int                     `json:"typemoney"`
	UrlPhoto         interface{}             `json:"url"`
	Available        bool                    `json:"available"`
	Insumos          []Pg_Mo_Insumo_Elements `json:"insumos"`
	SendToDelete     time.Time               `json:"sendtodelete"`
	DeletedDate      time.Time               `json:"deleteddate"`
	IsSendToDelete   bool                    `json:"issendtodelete"`
	IsAutomaticCost  bool                    `json:"isautomaticcost"`
	Costo            float64                 `json:"costo"`
	IsURLPrecharged  bool                    `json:"isurlprecharged"`
}

type Pg_Element_WithRating struct {
	IDElement        int                     `json:"id"`
	IDCategory       int                     `json:"idcategory"`
	NameCategory     string                  `json:"namecategory"`
	URLPhotoCategory string                  `json:"urlphotocategory"`
	Typefood         string                  `json:"typefood"`
	Name             string                  `json:"name"`
	Price            float32                 `json:"price"`
	Description      string                  `json:"description"`
	TypeMoney        int                     `json:"typemoney"`
	UrlPhoto         interface{}             `json:"url"`
	Available        bool                    `json:"available"`
	Orders           int                     `json:"orders"`
	Insumos          []Pg_Mo_Insumo_Elements `json:"insumos"`
	SendToDelete     time.Time               `json:"sendtodelete"`
	IsSendToDelete   bool                    `json:"issendtodelete"`
	IsAutomaticCost  bool                    `json:"isautomaticcost"`
	Costo            float64                 `json:"costo"`
	IsURLPrecharged  bool                    `json:"isurlprecharged"`
}

type Pg_ScheduleRange struct {
	IDSchedule        int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	MinutePerFraction int    `json:"minutesperfraction"`
	StartTime         string `json:"starttime"`
	EndTime           string `json:"endtime"`
	Available         bool   `json:"available"`
	NumberOfFractions int    `json:"numberfractions"`
	TimeZone          string `json:"timezone"`
	MaxOrders         int    `json:"maxorders"`
}

type Pg_Category_Element_ScheduleRange struct {
	Category int `json:"category"`
	Element  int `json:"element"`
	Schedule int `json:"schedule"`
}

type Pg_Carta struct {
	IDMenu       int    `json:"id"`
	IDBusiness   int    `json:"business"`
	Date         string `json:"date"`
	AcceptOrders bool   `json:"acceptorders"`
	Available    bool   `json:"available"`
}

type Pg_ScheduleRangeXMenu struct {
	IDSchedule int `json:"schedule"`
	IDMenu     int `json:"menu"`
}

type Pg_ElementXMenu struct {
	IDElement int `json:"element"`
	IDMenu    int `json:"menu"`
	Stock     int `json:"stock"`
}

type Pg_ToCarta_Mqtt struct {
	IdBusiness                int    `json:"idBusiness"`
	IdBanner_Category_Element int    `json:"idBCE"`
	IdType                    int    `json:"idType"`
	Url                       string `json:"url"`
}

type Pg_Import_StadisticOrders struct {
	IdElement   int     `json:"idelement"`
	Quantity    int     `json:"quantity"`
	TotalAmount float32 `json:"totalamount"`
	TotalCost   float32 `json:"totalcost"`
	Datetime    string  `json:"datetime"`
}

type Pg_StadisticByElement struct {
	Mon int `json:"1"`
	Tue int `json:"2"`
	Wed int `json:"3"`
	Thr int `json:"4"`
	Fri int `json:"5"`
	Sat int `json:"6"`
	Sun int `json:"7"`
}

type Pg_ElementsByCategory struct {
	Element   string `json:"element"`
	Available bool   `json:"available"`
}

type Pg_Element_ToCreate struct {
	IDElement        int                     `json:"id"`
	IDBusiness       int                     `json:"idbusiness"`
	IDCategory       int                     `json:"idcategory"`
	NameCategory     string                  `json:"namecategory"`
	TypeFood         string                  `json:"typefood"`
	UrlPhotoCategory string                  `json:"urlphotocategory"`
	Name             string                  `json:"name"`
	Price            float32                 `json:"price"`
	Description      string                  `json:"description"`
	TypeMoney        int                     `json:"typemoney"`
	Stock            int                     `json:"stock"`
	UrlPhoto         string                  `json:"url"`
	Insumos          []Pg_Mo_Insumo_Elements `json:"insumos"`
	Costo            float64                 `json:"costo"`
	IsAutomaticCost  bool                    `json:"isautomaticcost"`
	IsURLPrecharged  bool                    `json:"isurlprecharged"`
}

type Pg_Schedule_ToCreate struct {
	IDSchedule     int    `json:"idschedule"`
	Date           string `json:"date"`
	Starttime      string `json:"starttime"`
	Endtime        string `json:"endtime"`
	TimeZone       string `json:"timezone"`
	MaxOrders      int    `json:"maxorders"`
	ShowToComensal string `json:"showtocomensal"`
}
