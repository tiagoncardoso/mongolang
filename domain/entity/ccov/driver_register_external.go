package ccov

import "time"

type DriverRegisterExternal struct {
	Id              string
	ExternalId      string
	Driver          Driver
	DeviceRegisters []Vehicle
	Product         string
	ProductValue    float64
	CreationTime    time.Time
	UpdateTime      time.Time
	Status          int
	UserId          string
	ExternalUserId  string
}

func NewDriverRegisterExternal() *DriverRegisterExternal {
	return &DriverRegisterExternal{}
}