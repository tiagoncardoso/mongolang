package ccov

type Company struct {
	Document string `json:"Document"`
}

func NewCompany() *Company {
	return &Company{}
}

func (com *Company) CompanyIsVerttice() {
	com.Document = "35097794000138"
}
