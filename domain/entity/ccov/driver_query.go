package ccov

import "time"

type DriverQuery struct {
	Id               string         `bson:"_id"`
	BranchId         string         `bson:"BranchId"`
	Branch           BranchRegister `bson:"Branch"`
	Status           int16          `bson:"Status"`
	DriverDocument   string         `bson:"DriverDocument"`
	Trucks           []Vehicle      `bson:"Trucks"`
	OriginCity       string         `bson:"OriginCity"`
	OriginState      string         `bson:"OriginState"`
	DestinationCity  string         `bson:"DestinationCity"`
	DestinationState string         `bson:"DestinationState"`
	Product          string         `bson:"Product"`
	ProductValue     int32          `bson:"ProductValue"`
	CreationTime     time.Time      `bson:"CreationTime"`
}
