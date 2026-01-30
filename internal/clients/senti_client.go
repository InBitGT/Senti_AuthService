package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type CustomerClient struct {
	baseURL string
	key     string
	http    *http.Client
}

func NewCustomerClient() *CustomerClient {
	return &CustomerClient{
		baseURL: os.Getenv("SENTI_SERVICE_URL"),
		key:     os.Getenv("INTERNAL_API_KEY"),
		http:    &http.Client{Timeout: 12 * time.Second},
	}
}

/* ---------------- ADDRESS ---------------- */

type CreateAddressReq struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

type CreateAddressResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID uint `json:"id"`
	} `json:"data"`
}

func (c *CustomerClient) CreateAddress(req CreateAddressReq) (uint, error) {
	body, _ := json.Marshal(req)

	url := fmt.Sprintf("%s/api/addresses/internal", c.baseURL)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-Internal-Key", c.key)

	res, err := c.http.Do(httpReq)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return 0, fmt.Errorf("customer create address failed: %s", res.Status)
	}

	var out CreateAddressResp
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return 0, err
	}
	return out.Data.ID, nil
}

// (por ahora, tu endpoint actual borra; luego lo pasamos a soft delete)
func (c *CustomerClient) DeleteAddress(id uint) error {
	url := fmt.Sprintf("%s/api/addresses/internal/%d", c.baseURL, id)
	httpReq, _ := http.NewRequest("DELETE", url, nil)
	httpReq.Header.Set("X-Internal-Key", c.key)

	res, err := c.http.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("customer delete address failed: %s", res.Status)
	}
	return nil
}

/* ---------------- BRANCH (NUEVO) ---------------- */

type CreateBranchReq struct {
	Name        string `json:"name"`
	AddressID   uint   `json:"address_id"`
	Description string `json:"description"`
	TenantID    uint   `json:"tenant_id"`
}

type CreateBranchResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID uint `json:"id"`
	} `json:"data"`
}

// Espera que exista en CustomerService: POST /api/branches/internal
func (c *CustomerClient) CreateBranch(req CreateBranchReq) (uint, error) {
	body, _ := json.Marshal(req)

	url := fmt.Sprintf("%s/api/branches/internal", c.baseURL)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-Internal-Key", c.key)

	res, err := c.http.Do(httpReq)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return 0, fmt.Errorf("customer create branch failed: %s", res.Status)
	}

	var out CreateBranchResp
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return 0, err
	}
	return out.Data.ID, nil
}

/* ---------------- USER_BRANCH (NUEVO) ---------------- */

type CreateUserBranchReq struct {
	UserID   uint `json:"user_id"`
	BranchID uint `json:"branch_id"`
}

type CreateUserBranchResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID uint `json:"id"`
	} `json:"data"`
}

// Espera que exista en CustomerService: POST /api/user-branch/internal
func (c *CustomerClient) CreateUserBranch(req CreateUserBranchReq) (uint, error) {
	body, _ := json.Marshal(req)

	url := fmt.Sprintf("%s/api/user-branch/internal", c.baseURL)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-Internal-Key", c.key)

	res, err := c.http.Do(httpReq)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return 0, fmt.Errorf("customer create user_branch failed: %s", res.Status)
	}

	var out CreateUserBranchResp
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return 0, err
	}
	return out.Data.ID, nil
}
