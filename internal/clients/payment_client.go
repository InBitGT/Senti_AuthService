package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type PaymentClient struct {
	baseURL string
	key     string
	http    *http.Client
}

func NewPaymentClient() *PaymentClient {
	return &PaymentClient{
		baseURL: os.Getenv("PAYMENT_SERVICE_URL"),
		key:     os.Getenv("INTERNAL_API_KEY"),
		http:    &http.Client{Timeout: 12 * time.Second},
	}
}

type CreateSuscriptionReq struct {
	TenantID  uint       `json:"tenant_id"`
	PlanID    uint       `json:"plan_id"`
	StartedAt time.Time  `json:"started_at"`
	RenewAt   time.Time  `json:"renew_at"`
	EndAt     *time.Time `json:"end_at,omitempty"`
}

type CreateSuscriptionResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID uint `json:"id"`
	} `json:"data"`
}

// Espera que exista en PaymentService: POST /api/suscriptions/internal
func (c *PaymentClient) CreateSuscription(req CreateSuscriptionReq) (uint, error) {
	body, _ := json.Marshal(req)

	url := fmt.Sprintf("%s/api/suscriptions/internal", c.baseURL)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-Internal-Key", c.key)

	res, err := c.http.Do(httpReq)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return 0, fmt.Errorf("payment create suscription failed: %s", res.Status)
	}

	var out CreateSuscriptionResp
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return 0, err
	}
	return out.Data.ID, nil
}
