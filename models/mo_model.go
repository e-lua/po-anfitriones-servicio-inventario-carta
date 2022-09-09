package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mo_Providers struct {
	Type             int       `json:"type"`
	IDBusiness       int       `json:"idbusiness"`
	CreatedDate      time.Time `json:"createdDate"`
	Number           string    `json:"number"`
	ProviderName     string    `json:"providername"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	NameContact      string    `json:"namecontact"`
	Address          string    `json:"address"`
	ReferenceAddress string    `json:"referenceaddress"`
	Description      string    `json:"description"`
	Available        bool      `json:"available"`
	SendToDelete     time.Time `json:"sendtodelete"`
	IsDeleted        bool      `json:"isdeleted"`
	IsExported       bool      `json:"isexported"`
	DeletedDate      time.Time `json:"deleteddate"`
	IsSendToDelete   bool      `json:"issendtodelete"`
}

type Mo_Providers_Response struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Type             int                `json:"type"`
	Number           string             `json:"number"`
	ProviderName     string             `json:"providername"`
	Email            string             `json:"email"`
	Phone            string             `json:"phone"`
	NameContact      string             `json:"namecontact"`
	Address          string             `json:"address"`
	ReferenceAddress string             `json:"referenceaddress"`
	Description      string             `json:"description"`
	Available        bool               `json:"available"`
	SendToDelete     time.Time          `json:"sendtodelete"`
}

type Mo_StoreHouse struct {
	Name           string    `json:"name"`
	IDBusiness     int       `json:"idbusiness"`
	CreatedDate    time.Time `json:"createdDate"`
	Description    string    `json:"description"`
	IsExported     bool      `json:"isexported"`
	Available      bool      `json:"available"`
	SendToDelete   time.Time `json:"sendtodelete"`
	IsDeleted      bool      `json:"isdeleted"`
	DeletedDate    time.Time `json:"deleteddate"`
	IsSendToDelete bool      `json:"issendtodelete"`
}

type Mo_StoreHouse_Response struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Available    bool               `json:"available"`
	SendToDelete time.Time          `json:"sendtodelete"`
}

type Mo_Stock struct {
	Price        float64   `json:"price"`
	IdProvider   string    `json:"idprovider"`
	TimeZone     string    `json:"timezone"`
	CreatedDate  time.Time `json:"createdDate"`
	Quantity     int       `json:"quantity"`
	ProviderName string    `json:"providername"`
}

type Mo_Insumo struct {
	Name           string      `json:"name"`
	IDBusiness     int         `json:"idbusiness"`
	CreatedDate    time.Time   `json:"createdDate"`
	Measure        string      `json:"measure"`
	IDStoreHouse   string      `json:"idstorehouse"`
	NameStoreHouse string      `json:"namestorehouse"`
	Description    string      `json:"description"`
	Stock          []*Mo_Stock `json:"stock"`
	OutputStock    int         `json:"outputstock"`
	Available      bool        `json:"available"`
	IsDeleted      bool        `json:"isdeleted"`
	IsExported     bool        `json:"isexported"`
	SendToDelete   time.Time   `json:"sendtodelete"`
	DeletedDate    time.Time   `json:"deleteddate"`
	IsSendToDelete bool        `json:"issendtodelete"`
}

type Mo_Insumo_Response struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name"`
	Measure        string             `json:"measure"`
	IDStoreHouse   string             `json:"idstorehouse"`
	NameStoreHouse string             `json:"namestorehouse"`
	Description    string             `json:"description"`
	Stock          []Mo_Stock         `json:"stock"`
	OutputStock    int                `json:"outputstock"`
	Available      bool               `json:"available"`
	SendToDelete   time.Time          `json:"sendtodelete"`
}

type Mo_Insumo_Stock_Adjust_Requst struct {
	AmountToAdjust int        `json:"outputstock"`
	IsAdjust       bool       `json:"isadjust"`
	Stock          []Mo_Stock `json:"stock"`
}

type Mo_Precharged_Element struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Mo_In_2_NotifyData struct {
	Key   string `bson:"Key" json:"Key"`
	Value int    `bson:"Value" json:"Value"`
}

type Mo_Insumo_NotifyData struct {
	NotifyData []Mo_In_2_NotifyData
}
