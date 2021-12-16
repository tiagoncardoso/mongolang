package ccov

import "time"

type Vehicle struct {
	Document           string    `bson:"Document"`
	Name               string    `bson:"Name"`
	Document2          string    `bson:"Document2"`
	NameDocument2      string    `bson:"NameDocument2"`
	DueDocument2       time.Time `bson:"DueDocument2"`
	Postcode           string    `bson:"Postcode"`
	State              string    `bson:"State"`
	City               string    `bson:"City"`
	Address            string    `bson:"Address"`
	Number             string    `bson:"Number"`
	County             string    `bson:"County"`
	Complement         string    `bson:"Complement"`
	Landline           string    `bson:"Landline"`
	Contact            string    `bson:"Contact"`
	CommercialLandline string    `bson:"CommercialLandline"`
	Plate              string    `bson:"Plate"`
	PostcodePlate      string    `bson:"PostcodePlate"`
	StatePlate         string    `bson:"StatePlate"`
	CityPlate          string    `bson:"CityPlate"`
	Document3          string    `bson:"Document3"`
	DeviceBrand        string    `bson:"DeviceBrand"`
	DeviceModel        string    `bson:"DeviceModel"`
	Document4          string    `bson:"Document4"`
	Year               string    `bson:"Year"`
	Color              string    `bson:"Color"`
}
