package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type device struct {
	Code     int8   `bson:"Code"`
	Document string `bson:"Document"`
	Plate    string `bson:"Plate"`
}

type document struct {
	DocumentType int8               `bson:"DocumentType"`
	DocumentId   primitive.ObjectID `bson:"DocumentId"`
	DocumentName string             `bson:"DocumentName"`
}

type reference struct {
	Phone string `bson:"Phone"`
	Name  string `bson:"Name"`
}

type CcovDriver struct {
	ID                             primitive.ObjectID `bson:"_id"`
	Delete                         bool               `bson:"Delete"`
	UserId                         primitive.ObjectID `bson:"UserId"`
	CompanyId                      primitive.ObjectID `bson:"CompanyId"`
	Company                        string             `bson:"Company"`
	BranchId                       primitive.ObjectID `bson:"BranchId"`
	Branch                         string             `bson:"Branch"`
	DriverQueryId                  primitive.ObjectID `bson:"DriverQueryId"`
	Status                         int8               `bson:"Status"`
	CreationTime                   time.Time          `bson:"CreationTime"`
	UpdateTime                     time.Time          `bson:"UpdateTime"`
	SubmissionTime                 time.Time          `bson:"SubmissionTime"`
	Document                       string             `bson:"Document"`
	DriverProfileId                primitive.ObjectID `bson:"DriverProfileId"`
	DriverProfile                  string             `bson:"DriverProfile"`
	Name                           string             `bson:"Name"`
	DateOfBirthday                 time.Time          `bson:"DateOfBirthday"`
	Document2                      string             `bson:"Document2"`
	IssueStateDocument2            string             `bson:"IssueStateDocument2"`
	IssueStateDocument2Description string             `bson:"IssueStateDocument2Description"`
	Document3                      string             `bson:"Document3"`
	CategoryDocument3              int32              `bson:"CategoryDocument3"`
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
	Contact                        string             `bson:"Contact"`
	CommercialLandline             string             `bson:"CommercialLandline"`
	CommercialContact              string             `bson:"CommercialContact"`
	ReferenceLandline              string             `bson:"ReferenceLandline"`
	ReferenceContact               string             `bson:"ReferenceContact"`
	DeviceRegisters                []device           `bson:"DeviceRegisters"`
	Comments                       string             `bson:"Comments"`
	UserEmployee                   string             `bson:"UserEmployee"`
	ProductValue                   float64            `bson:"ProductValue"`
	ProductId                      primitive.ObjectID `bson:"ProductId"`
	Product                        string             `bson:"Product"`
	RegisterSpecial                bool               `bson:"RegisterSpecial"`
	Score                          int8               `bson:"Score"`
	Attachments                    []string           `bson:"Attachments"`
	References                     []reference        `bson:"References"`
	Documents                      []document         `bson:"Documents"`
	BlockedDriver                  bool               `bson:"BlockedDriver"`
}

func NewCcovDriver() *CcovDriver {
	return &CcovDriver{}
}
