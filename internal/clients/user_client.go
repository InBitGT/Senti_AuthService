package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type UserClient struct {
	baseURL string
	key     string
	http    *http.Client
}

func NewUserClient() *UserClient {
	return &UserClient{
		baseURL: os.Getenv("USER_SERVICE_URL"),
		key:     os.Getenv("INTERNAL_API_KEY"),
		http:    &http.Client{Timeout: 10 * time.Second},
	}
}

type CreateAdminReq struct {
	TenantID  uint   `json:"tenant_id"`
	AddressID uint   `json:"address_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	RoleID    uint   `json:"role_id"`
}

type CreateAdminResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID uint `json:"id"`
	} `json:"data"`
}

func (c *UserClient) CreateAdmin(req CreateAdminReq) (uint, error) {
	body, _ := json.Marshal(req)

	url := fmt.Sprintf("%s/api/v1/users/internal/admin", c.baseURL)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-Internal-Key", c.key)

	res, err := c.http.Do(httpReq)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return 0, fmt.Errorf("user create admin failed: %s", res.Status)
	}

	var out CreateAdminResp
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return 0, err
	}

	return out.Data.ID, nil
}
