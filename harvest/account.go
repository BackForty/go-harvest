package harvest

type AccountService struct {
	Service
}

type Company struct {
	BaseURI            string  `json:"base_uri"`
	FullDomain         string  `json:"full_domain"`
	Name               string  `json:"name"`
	Active             bool    `json:"active"`
	WeekStartDay       string  `json:"week_start_day"`
	TimeFormat         string  `json:"hours_minutes"`
	Clock              string  `json:"clock"`
	DecimalSymbol      string  `json:"decimal_symbol"`
	ColorScheme        string  `json:"color_scheme"`
	Modules            Modules `json:"modules"`
	ThousandsSeperator string  `json:"thousands_seperator"`
}

type Account struct {
	Company Company
	Person  Person
}

func (a *AccountService) Find() (err error, account Account) {
	resourceURL := "/account/who_am_i.json"
	err = a.find(resourceURL, &account)

	return err, account
}
