package models

type Balance struct {
	Id       uint32 `json:"-"`
	IdClient uint32 `json:"-"`
	Value    int64  `json:"valor"`
}
