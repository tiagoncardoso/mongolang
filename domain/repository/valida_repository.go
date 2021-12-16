package repository

import "ccovdata/domain/entity/valida"

type DriverRepository interface {
	Insert() (error, valida.DriverRegister)
}

type VehicleRepository interface {
	Insert() (error, valida.VehicleRegister)
}

type VehicleOwnerRepository interface {
	Insert() (error, VehicleOwnerRepository)
}

type TravelRepository interface {
	Insert() (error, valida.TravelRegister)
}

type MainRegisterRepository interface {
	Insert() (error, valida.MainRegister)
}

type CodeRepository interface {
	Insert() (error, valida.CodeRegister)
}

type ResultRepository interface {
	Insert() (error, valida.ResultRegister)
}
