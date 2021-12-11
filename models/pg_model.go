package models

type Pg_Category struct {
	IDCategory int         `json:"id"`
	IDBusiness int         `json:"business"`
	Name       string      `json:"name"`
	UrlPhoto   interface{} `json:"url"`
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
	IDElement    int         `json:"id"`
	IDCategory   int         `json:"category"`
	NameCategory string      `json:"nameCategory"`
	Name         string      `json:"name"`
	Price        float32     `json:"price"`
	Description  string      `json:"description"`
	TypeMoney    int         `json:"typeMoney"`
	UrlPhoto     interface{} `json:"url"`
}

type Pg_ScheduleRange struct {
	IDSchedule        int    `json:"id"`
	IDBusiness        int    `json:"business"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	MinutePerFraction int    `json:"minutesPerFraction"`
	StartTime         string `json:"startTime"`
	EndTime           string `json:"endTime"`
	MaxOrders         int    `json:"maxOrders"`
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
	AcceptOrders bool   `json:"acceptOrders"`
	Available    bool   `json:"isAvailable"`
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
