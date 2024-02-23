package models

import "time"

type Transaction struct {
	Id          uint32    `json:"-"`
	IdClient    uint32    `json:"-"`
	Value       int64     `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}
