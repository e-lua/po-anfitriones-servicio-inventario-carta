package pre_charged

import models "github.com/Aphofisis/po-anfitrion-servicio-inventario-carta/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type ResponseListPreCharged struct {
	Error     bool                            `json:"error"`
	DataError string                          `json:"dataError"`
	Data      []*models.Mo_Precharged_Element `json:"data"`
}
