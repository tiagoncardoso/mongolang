package dto

import "ccovdata/domain/entity/ccov"

type RegisterDtoInput struct {
	ID       string
	Driver   ccov.DriverRegisterExternal
	Vehicles []ccov.Vehicle
}

type RegisterDtoOutput struct {
	ID           string
	Status       string
	ErrorMessage string
}
