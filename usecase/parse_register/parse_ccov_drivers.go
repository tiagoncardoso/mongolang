package parse_register

import (
	"ccovdata/domain/factory"
	"ccovdata/usecase/dto"
	"fmt"
	"strconv"
)

type ParseRegister struct {
	Repository factory.DriverRegisterExternalRepository
}

func NewParseRegister(repository factory.DriverRegisterExternalRepository) *ParseRegister {
	return &ParseRegister{Repository: repository}
}

func (pr *ParseRegister) ProcessCcovRegisterExternal(registerInput []*dto.RegisterDtoInput) {
	var output = dto.RegisterDtoOutput{}
	for i, v := range registerInput {
		output = dto.RegisterDtoOutput{
			ID:           "Item " + strconv.Itoa(i),
			Status:       "certo",
			ErrorMessage: "nada",
		}

		pr.PrintRegister(v)

		fmt.Println(output)
	}
}

func (pr *ParseRegister) PrintRegister(register *dto.RegisterDtoInput) {
	fmt.Print("%s\n", register)
}
