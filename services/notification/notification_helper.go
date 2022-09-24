package notification

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
