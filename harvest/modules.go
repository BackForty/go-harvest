package harvest

type Modules struct {
	Expenses  bool `json:"expenses"`
	Invoices  bool `json:"invoices"`
	Estimates bool `json:"estimates"`
	Approval  bool `json:"approval"`
}
