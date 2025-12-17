package auth

type RegisterCompanyRequest struct {
	Tenant struct {
		Code  string `json:"code"`
		Name  string `json:"name"`
		NIT   string `json:"nit"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	} `json:"tenant"`

	Address struct {
		Line1      string `json:"line1"`
		Line2      string `json:"line2"`
		City       string `json:"city"`
		State      string `json:"state"`
		Country    string `json:"country"`
		PostalCode string `json:"postal_code"`
	} `json:"address"`

	AdminUser struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
	} `json:"admin_user"`
}

type RegisterCompanyResponse struct {
	TenantID  uint `json:"tenant_id"`
	AddressID uint `json:"address_id"`
	UserID    uint `json:"user_id"`
}
