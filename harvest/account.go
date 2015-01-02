package harvest

// AccountService wraps interaction with the API for working with Accounts
type AccountService struct {
	Service
}

// Company represents a Harvest Company
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

// Account represents a Harvest Account
type Account struct {
	Company Company
	Person  Person
}

// Find fetches account details
func (a *AccountService) Find() (account Account, err error) {
	resourceURL := "/account/who_am_i.json"
	err = a.find(resourceURL, &account)

	return account, err
}
