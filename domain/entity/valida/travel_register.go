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
	TipoCarga        []int     `json:"tipo_carga"`
	ValorCarga       float64   `json:"valor_carga"`
	OrigemCidadeTxt  string    `json:"origem_cidade_txt"`
	DestinoCidadeTxt string    `json:"destino_cidade_txt"`
	Criacao          time.Time `json:"criacao"`
}

func NewTravelRegister(createdAt time.Time) *TravelRegister {
	return &TravelRegister{
		OrigemPaisId:    1,
		OrigemUfId:      21,
		OrigemCidadeId:  3980,
		DestinoPaisId:   1,
		DestinoUfId:     20,
		DestinoCidadeId: 3829,
		Criacao:         createdAt,
	}
}

func (tr *TravelRegister) SetValorCarga(valor float64) {
	tr.ValorCarga = valor
}

func (tr *TravelRegister) SetChargerType(chargerType string) {
	var tipoCarga []int

	switch chargerType {
	case "ALGODAO":
	case "Algodão":
		tipoCarga = append(tipoCarga, 7)
		break
	case "Soja":
		tipoCarga = append(tipoCarga, 170)
		break

	default:
		tipoCarga = append(tipoCarga, 162)
		break
	}
	tipoCarga = append(tipoCarga, 162)

	tr.TipoCarga = tipoCarga
}
