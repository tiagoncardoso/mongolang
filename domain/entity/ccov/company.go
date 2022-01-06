package ccov

type Company struct {
	Document string `bson:"Document"`
	Name     string `bson:"Name"`
}

func NewCompany() *Company {
	return &Company{}
}

func (com *Company) CompanyIsVerttice() {
	com.Document = "35097794000138"
}
