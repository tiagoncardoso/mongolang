package valida

import "time"

type TravelRegister struct {
	Id               int32     `json:"id"`
	OrigemPaisId     int16     `json:"origem_pais_id"`
	OrigemUfId       int32     `json:"origem_uf_id"`
	OrigemCidadeId   int32     `json:"origem_cidade_id"`
	DestinoPaisId    int16     `json:"destino_pais_id"`
	DestinoUfId      int32     `json:"destino_uf_id"`
	DestinoCidadeId  int32     `json:"destino_cidade_id"`
	TipoCarga        []int32   `json:"tipo_carga"`
	ValorCarga       float64   `json:"valor_carga"`
	OrigemCidadeTxt  string    `json:"origem_cidade_txt"`
	DestinoCidadeTxt string    `json:"destino_cidade_txt"`
	Criacao          time.Time `json:"criacao"`
}
