package valida

import "time"

type CodeRegister struct {
	Id      int32     `json:"id"`
	Type    string    `json:"type"`
	Counter int32     `json:"counter"`
	Literal string    `json:"literal"`
	Criacao time.Time `json:"criacao"`
}
