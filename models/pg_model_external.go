package models

import "time"

type Pg_Carta_External struct {
	IDCarta            int       `json:"id"`
	Date               time.Time `json:"date"`
	AvailableForOrders bool      `json:"availablefororders"`
	Visible            bool      `json:"visible"`
	Elements           int       `json:"elements"`
	ScheduleRanges     int       `json:"scheduleranges"`
}

type Pg_Category_External struct {
	IDCategory       int    `json:"idcategory"`
	Name             string `json:"namecategory"`
	UrlPhoto         string `json:"urlphotocategory"`
	AmountOfElements int    `json:"elements"`
}

type Pg_Element_With_Stock_External struct {
	IDElement        int     `json:"id"`
	IDCarta          int     `json:"idcarta"`
	IDBusiness       int     `json:"idbusiness"`
	IDCategory       int     `json:"idcategory"`
	NameCategory     string  `json:"namecategory"`
	UrlPhotoCategory string  `json:"urlphotocategory"`
	Name             string  `json:"name"`
	Price            float32 `json:"price"`
	Description      string  `json:"description"`
	TypeMoney        int     `json:"typemoney"`
	Stock            int     `json:"stock"`
	UrlPhoto         string  `json:"url"`
}

type Pg_ScheduleRange_External struct {
	IDSchedule        int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	MinutePerFraction int    `json:"minutesperfraction"`
	StartTime         string `json:"starttime"`
	EndTime           string `json:"endtime"`
	NumberOfFractions int    `json:"numberfractions"`
	MaxOrders         int    `json:"maxOrders"`
}
