package main

import (
	"ccovdata/driver/registration"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type personModel struct {
	Name string				`json:"nome"`
	Cpf int64				`json:"cpf"`
	Criacao string			`json:"criacao"`
	CadastroGaveta string	`json:"cadastroGaveta"`
	ImgCnh string			`json:"img_cng"`
	ImgRg string			`json:"img_rg"`
	ImgEndereco string		`json:"img_endereco"`
	Uf string				`json:"uf"`
}

func main() {
	cg := fmt.Sprintf(registration.GetRegistration())

	p := &personModel{
		Name:           "Teste",
		Cpf:            99999999999,
		Criacao:        "2021-11-26 10:38",
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

	fmt.Println(string(out))
}