package intercom

// A ContactCompany represents a Company association for a Contact.
type ContactCompany struct {
	CompanyID string `json:"company_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Remove    *bool  `json:"remove,omitempty"`
}

// UserCompany is an alias for ContactCompany for backwards compatibility.
type UserCompany = ContactCompany

// MakeContactCompaniesFromCompanies converts a slice of Companies to ContactCompanies.
func MakeContactCompaniesFromCompanies(companies []Company) []ContactCompany {
	contactCompanies := make([]ContactCompany, len(companies))
	for i := 0; i < len(companies); i++ {
		contactCompanies[i] = ContactCompany{
			CompanyID: companies[i].CompanyID,
			Name:      companies[i].Name,
			Remove:    companies[i].Remove,
		}
	}
	return contactCompanies
}

// MakeUserCompaniesFromCompanies is an alias for MakeContactCompaniesFromCompanies for backwards compatibility.
func MakeUserCompaniesFromCompanies(companies []Company) []ContactCompany {
	return MakeContactCompaniesFromCompanies(companies)
}
