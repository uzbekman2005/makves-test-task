package models

type GetAllFileInfoReq struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type GetAllFileInfoRes struct {
	Count int64         `json:"count"`
	Data  []interface{} `json:"data"`
}

type GetByidReq struct {
	Id int `json:"id"`
}
