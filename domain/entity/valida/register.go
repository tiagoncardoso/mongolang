package valida

import "time"

type Register struct {
	ID               int64     `json:"id"`
	ProfissionalId   int64     `json:"profissional_id"`
	MotorId          int64     `json:"motor_id"`
	Carreta1Id       int64     `json:"carreta1_id"`
	Carreta2Id       int64     `json:"carreta2_id"`
	Carreta3Id       int64     `json:"carreta3_id"`
	ViagemId         int64     `json:"viagem_id"`
	TipoCadastro     string    `json:"tipo_cadastro"`
	Validado         bool      `json:"validado"`
	ValidadeCadastro string    `json:"validade_cadastro"`
	Status           string    `json:"status"`
	Criacao          time.Time `json:"criacao"`
	Criminal         bool      `json:"criminal"`
	Plus             bool      `json:"plus"`
}

func NewRegister(profissionalId int64, motorId int64, carreta1Id int64, carreta2Id int64, carreta3Id int64, viagemId int64) *Register {
	return &Register{
		ProfissionalId: profissionalId,
		MotorId:        motorId,
		Carreta1Id:     carreta1Id,
		Carreta2Id:     carreta2Id,
		Carreta3Id:     carreta3Id,
		ViagemId:       viagemId,
		Validado:       true,
		TipoCadastro:   "PROFISSIONAL_VEICULO",
		Criacao:        time.Now(),
		Criminal:       false,
	}
}

func (re *Register) SetPlus(isPlus bool) {
	re.Plus = isPlus
}

func (re *Register) SetRegisterValidity(creation time.Time, validity time.Time, score int) {
	if score <= 10 {
		re.ValidadeCadastro = creation.Format("YYYY-MM-DD")
	}
}

func (re *Register) SetRegisterValidation(score int) {
	if score <= 10 {
		re.Status = "DIVERGENTE"
	} else {
		re.Status = "ADEQUADO"
	}
}
