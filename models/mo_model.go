package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mo_Carta_Ini struct {
	Date       time.Time `bson:"date" json:"date,omitempty"`
	IDBusiness int       `bson:"idbusiness" json:"idbusiness,omitempty"`
}

type Mo_Carta_Available_Visible struct {
	IDCarta            string `bson:"idcarta" json:"idcarta,omitempty"`
	IDBusiness         int    `bson:"idbusiness" json:"idbusiness,omitempty"`
	AvailableForOrders bool   `bson:"availablefororders" json:"availablefororders,omitempty"`
	Visible            bool   `bson:"visible" json:"visible,omitempty"`
}

type Mo_Carta struct {
	IDCarta            primitive.ObjectID      `bson:"_id" json:"_id,omitempty"`
	Date               time.Time               `bson:"date" json:"date,omitempty"`
	UpdatedDate        time.Time               `bson:"updatedDate" json:"updatedDate,omitempty"`
	IDBusiness         int                     `bson:"idbusiness" json:"idbusiness,omitempty"`
	AvailableForOrders bool                    `bson:"availablefororders" json:"availablefororders,omitempty"`
	Visible            bool                    `bson:"visible" json:"visible,omitempty"`
	Categories         []Mo_Category           `bson:"categories" json:"categories,omitempty"`
	Elements           []Mo_Element_With_Stock `bson:"elements" json:"elements,omitempty"`
	ScheduleRange      []Mo_ScheduleRange      `bson:"scheduleranges" json:"scheduleranges,omitempty"`
}

type Mo_Category struct {
	IDCategory       int         `json:"id"`
	Name             string      `json:"name"`
	AmountOfElements int         `json:"elements"`
	UrlPhoto         interface{} `json:"urlphotocategory"`
}

type Mo_Element_With_Stock struct {
	IDElement   int         `json:"id"`
	IDCategory  int         `json:"category"`
	Name        string      `json:"name"`
	Price       float32     `json:"price"`
	Description string      `json:"description"`
	TypeMoney   int         `json:"typemoney"`
	Stock       int         `json:"stock"`
	UrlPhoto    interface{} `json:"url"`
}

type Mo_ScheduleRange struct {
	IDSchedule        int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	MinutePerFraction int    `json:"minutesperfraction"`
	StartTime         string `json:"starttime"`
	EndTime           string `json:"endtime"`
	NumberOfFractions int    `json:"numberfractions"`
	MaxOrders         int    `json:"maxOrders"`
}
