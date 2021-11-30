package driver

import (
	mondrv "ccovdata/ccovdb"
	"ccovdata/driver/registration"
	"encoding/json"
	"log"
	"os"
	"time"
)

type driverModel struct {
	Name 							string				`json:"nome"`
	Cpf 							string				`json:"cpf"`
	Criacao 						time.Time			`json:"criacao"`
	CadastroGaveta 					json.RawMessage		`json:"cadastroGaveta"`
	ImgCnh 							string				`json:"img_cng"`
	ImgRg 							string				`json:"img_rg"`
	ImgEndereco 					string				`json:"img_endereco"`
	Uf 								string				`json:"uf"`
}

func BuildValidaDriver(driver *mondrv.Driver) string {
	cg := registration.GetRegistration()

	p := &driverModel{
		Name:           driver.Name,
		Cpf:            driver.Document,
		Criacao:        driver.CreationTime,
		CadastroGaveta: cg,
		ImgCnh:         "file_cnh.jpg",
		ImgRg:          "file_rg.png",
		ImgEndereco:    "file_endereco.png",
		Uf:             "SP",
	}

	out, err := json.MarshalIndent(p, "", "  ")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return string(out)
}