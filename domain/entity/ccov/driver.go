package ccov

import "time"

type Driver struct {
	Document            string    `bson:"Document"`
	Name                string    `bson:"Name"`
	DriverProfile       string    `bson:"DriverProfile"`
	DateOfBirthday      time.Time `bson:"DateOfBirthday"`
	Document2           string    `bson:"Document2"`
	IssueStateDocument2 string    `bson:"IssueStateDocument2"`
	Document3           string    `bson:"Document3"`
	CategoryDocument3   string    `bson:"CategoryDocument3"`
	DueDateDocument3    time.Time `bson:"DueDateDocument3"`
	EmitterDocument3    string    `bson:"EmitterDocument3"`
	IssueStateDocument3 string    `bson:"IssueStateDocument3"`
	IssueCityDocument3  string    `bson:"IssueCityDocument3"`
	FirstDateDocument3  time.Time `bson:"FirstDateDocument3"`
	MotherName          string    `bson:"MotherName"`
	FatherName          string    `bson:"FatherName"`
	Postcode            string    `bson:"Postcode"`
	Address             string    `bson:"Address"`
	Number              string    `bson:"Number"`
	Complement          string    `bson:"Complement"`
	County              string    `bson:"County"`
	State               string    `bson:"State"`
	City                string    `bson:"City"`
	Landline            string    `bson:"Landline"`
	Contact             string    `bson:"Contact"`
	CommercialLandline  string    `bson:"CommercialLandline"`
	ReferenceLandline   string    `bson:"ReferenceLandline"`
}
