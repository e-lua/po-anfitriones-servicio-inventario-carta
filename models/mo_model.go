package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mo_Providers struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Type             string             `json:"type"`
	IDBusiness       int                `json:"idbusiness"`
	CreatedDate      time.Time          `json:"createdDate"`
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
	IsDelete         bool               `json:"isdeleted"`
	IsExported       bool               `json:"isexported"`
	DeletedDate      time.Time          `json:"deleteddate"`
	IsSendToDelete   bool               `json:"issendtodelete"`
}

type Mo_Providers_Response struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Type             string             `json:"type"`
	IDBusiness       int                `json:"idbusiness"`
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
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name"`
	IDBusiness     int                `json:"idbusiness"`
	CreatedDate    time.Time          `json:"createdDate"`
	Description    string             `json:"description"`
	IsExported     bool               `json:"isexported"`
	Available      bool               `json:"available"`
	SendToDelete   time.Time          `json:"sendtodelete"`
	IsDelete       bool               `json:"isdeleted"`
	DeletedDate    time.Time          `json:"deleteddate"`
	IsSendToDelete bool               `json:"issendtodelete"`
}

type Mo_StoreHouse_Response struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name         string             `json:"name"`
	IDBusiness   int                `json:"idbusiness"`
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
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name"`
	IDBusiness     int                `json:"idbusiness"`
	CreatedDate    time.Time          `json:"createdDate"`
	Measure        string             `json:"measure"`
	StoreHouse     Mo_StoreHouse      `json:"storehouse"`
	Description    string             `json:"description"`
	Stock          []*Mo_Stock        `json:"stock"`
	Available      bool               `json:"available"`
	IsDelete       bool               `json:"isdeleted"`
	IsExported     bool               `json:"isexported"`
	SendToDelete   time.Time          `json:"sendtodelete"`
	DeletedDate    time.Time          `json:"deleteddate"`
	IsSendToDelete bool               `json:"issendtodelete"`
}

type Mo_Insumo_Response struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name         string             `json:"name"`
	IDBusiness   int                `json:"idbusiness"`
	Measure      string             `json:"measure"`
	StoreHouse   Mo_StoreHouse      `json:"storehouse"`
	Description  string             `json:"description"`
	Stock        []*Mo_Stock        `json:"stock"`
	Available    bool               `json:"available"`
	SendToDelete time.Time          `json:"sendtodelete"`
}
