package entity

import (
	"ccovdata/entity/locker_register"
	"encoding/json"
	"log"
	"os"
	"time"
)

type ValidaDriver struct {
	Name 							string										`json:"nome"`
	Cpf 							string										`json:"cpf"`
	Criacao 						time.Time									`json:"criacao"`
	CadastroGaveta 					locker_register.LockerRegister				`json:"cadastroGaveta"`
	ImgCnh 							string										`json:"img_cng"`
	ImgRg 							string										`json:"img_rg"`
	ImgEndereco 					string										`json:"img_endereco"`
	Uf 								string										`json:"uf"`
}

func (dr *ValidaDriver) BuildValidaDriver(driver *ccovDriver) string {
	dr.Name = driver.Name;
	dr.Cpf = driver.Document;
	dr.Criacao = driver.CreationTime;
	dr.CadastroGaveta = cg;
	dr.ImgCnh = "file_cnh.jpg";
	dr.ImgRg = "file_rg.png";
	dr.ImgEndereco = "file_endereco.png";
	dr.Uf = "SP";

	out, err := json.MarshalIndent(dr, "", "  ")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return string(out)
}