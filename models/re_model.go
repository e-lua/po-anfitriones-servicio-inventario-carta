package models

type Re_SetGetCode struct {
	IdBusiness int                               `json:"idbusiness"`
	CartaData  Pg_Category_Element_ScheduleRange `json:"cartadata"`
}
