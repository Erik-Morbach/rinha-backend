package models

type Client struct {
	Id    uint32 `json:"-"`
	Name  string `json:"nome"`
	Limit int64  `json:"limite"`
}
