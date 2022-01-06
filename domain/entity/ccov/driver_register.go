package ccov

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DriverRegister struct {
	Id                             primitive.ObjectID `bson:"_id"`
	Delete                         bool               `bson:"Delete"`
	DriverExternalId               primitive.ObjectID `bson:"DriverExternalId"`
	CreationTime                   time.Time          `bson:"CreationTime"`
	Document                       string             `bson:"Document"`
	DriverProfile                  string             `bson:"DriverProfile"`
	Name                           string             `bson:"Name"`
	DateOfBirthday                 time.Time          `bson:"DateOfBirthday"`
	Document2                      string             `bson:"Document2"`
	IssueStateDocument2            string             `bson:"IssueStateDocument2"`
	IssueStateDocument2Description string             `bson:"IssueStateDocument2Description"`
	Document3                      string             `bson:"Document3"`
	CategoryDocument3              int                `bson:"CategoryDocument3"`
	DueDateDocument3               time.Time          `bson:"DueDateDocument3"`
	EmitterDocument3               string             `bson:"EmitterDocument3"`
	IssueStateDocument3            string             `bson:"IssueStateDocument3"`
	IssueStateDocument3Description string             `bson:"IssueStateDocument3Description"`
	IssueCityDocument3             string             `bson:"IssueCityDocument3"`
	IssueCityDocument3Description  string             `bson:"IssueCityDocument3Description"`
	FirstDateDocument3             time.Time          `bson:"FirstDateDocument3"`
	MotherName                     string             `bson:"MotherName"`
	FatherName                     string             `bson:"FatherName"`
	Postcode                       string             `bson:"Postcode"`
	Address                        string             `bson:"Address"`
	Number                         string             `bson:"Number"`
	Complement                     string             `bson:"Complement"`
	County                         string             `bson:"County"`
	State                          string             `bson:"State"`
	StateDescription               string             `bson:"StateDescription"`
	City                           string             `bson:"City"`
	CityDescription                string             `bson:"CityDescription"`
	Landline                       string             `bson:"Landline"`
	CommercialLandline             string             `bson:"CommercialLandline"`
	ReferenceLandline              string             `bson:"ReferenceLandline"`
	Protocol                       string             `bson:"Protocol"`
	ProductValue                   float64            `bson:"ProductValue"`
	DeviceRegisters                []Vehicle          `bson:"DeviceRegisters"`
	Product                        string             `bson:"Product"`
	RegisterSpecial                bool               `bson:"RegisterSpecial"`
	Score                          int                `bson:"Score"`
	CompletionTime                 time.Time          `bson:"CompletionTime"`
	ValidityTime                   time.Time          `bson:"ValidityTime"`
	BlockedDriver                  bool               `bson:"BlockedDriver"`
	ExternalId                     string             `bson:"ExternalId"`
	Company                        string             `bson:"Company"`
	CompanyID                      string             `bson:"CompanyID"`
	CompanyExtra                   *Company           `bson:"CompanyExtra"`
	CompanyPortalId                int                `bson:"CompanyPortalId"`
	RegisterExtra                  *DriverRegister    `bson:"RegisterExtra"`
}

func NewDriverRegister() *DriverRegister {
	return &DriverRegister{}
}

func (dr *DriverRegister) IsPlus() bool {
	productTest := dr.Product == "Algodão" || dr.Product == "ALGODAO"
	scoreTest := dr.RegisterExtra.Score > 10

	return productTest && scoreTest
}
