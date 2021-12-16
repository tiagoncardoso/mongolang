package ccov

type BranchRegister struct {
	Id       int32  `bson:"_id"`
	Company  string `bson:"Company"`
	Name     string `bson:"Name"`
	Document string `bson:"Document"`
}
