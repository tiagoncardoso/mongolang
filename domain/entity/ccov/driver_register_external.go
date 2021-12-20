package ccov

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DriverRegisterExternal struct {
	Id              primitive.ObjectID `bson:"_id"`
	ExternalId      string             `bson:"ExternalId"`
	Driver          Driver             `bson:"Driver"`
	DeviceRegisters []Vehicle          `bson:"DeviceRegisters"`
	Product         string             `bson:"Product"`
	ProductValue    float64            `bson:"ProductValue"`
	CreationTime    time.Time          `bson:"CreationTime"`
	UpdateTime      time.Time          `bson:"UpdateTime"`
	Status          int                `bson:"Status"`
	UserId          string             `bson:"UserId"`
	ExternalUserId  string             `bson:"ExternalUserId"`
	Company         string             `bson:"Company"`
	CompanyID       string             `bson:"CompanyID"`
	CompanyExtra    *Company           `bson:"CompanyExtra"`
	CompanyPortalId int                `bson:"CompanyPortalId"`
	RegisterExtra   *DriverRegister    `bson:"RegisterExtra"`
}

func NewDriverRegisterExternal() *DriverRegisterExternal {
	return &DriverRegisterExternal{}
}
