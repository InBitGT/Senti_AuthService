package auth

import "time"

type RegisterCompanyRequest struct {
	Tenant struct {
		Code    string `json:"code"`
		Name    string `json:"name"`
		Picture string `json:"picture,omitempty"`
		NIT     string `json:"nit"`
		Phone   string `json:"phone"`
		Email   string `json:"email"`
	} `json:"tenant"`

	Address struct {
		Line1      string `json:"line1"`
		Line2      string `json:"line2"`
		City       string `json:"city"`
		State      string `json:"state"`
		Country    string `json:"country"`
		PostalCode string `json:"postal_code"`
	} `json:"address"`

	Subscription struct {
		PlanID    uint       `json:"plan_id"`
		StartedAt *time.Time `json:"started_at,omitempty"`
		RenewAt   *time.Time `json:"renew_at,omitempty"`
		EndAt     *time.Time `json:"end_at,omitempty"`
	} `json:"subscription"`

	Branch struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"branch"`

	AdminUser struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		Username  string `json:"username,omitempty"`
	} `json:"admin_user"`
}

type RegisterCompanyResponse struct {
	TenantID      uint `json:"tenant_id"`
	AddressID     uint `json:"address_id"`
	BranchID      uint `json:"branch_id"`
	SuscriptionID uint `json:"suscription_id"`
	UserID        uint `json:"user_id"`
}
