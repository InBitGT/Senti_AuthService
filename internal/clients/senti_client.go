package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type SentiClient struct {
	baseURL string
	key     string
	http    *http.Client
}

func NewSentiClient() *SentiClient {
	return &SentiClient{
		baseURL: os.Getenv("SENTI_SERVICE_URL"),
		key:     os.Getenv("INTERNAL_API_KEY"),
		http:    &http.Client{Timeout: 10 * time.Second},
	}
}

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

func (c *SentiClient) CreateAddress(req CreateAddressReq) (uint, error) {
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
		return 0, fmt.Errorf("senti create address failed: %s", res.Status)
	}

	var out CreateAddressResp
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return 0, err
	}
	return out.Data.ID, nil
}

func (c *SentiClient) DeleteAddress(id uint) error {
	url := fmt.Sprintf("%s/api/addresses/internal/%d", c.baseURL, id)
	httpReq, _ := http.NewRequest("DELETE", url, nil)
	httpReq.Header.Set("X-Internal-Key", c.key)

	res, err := c.http.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("senti delete address failed: %s", res.Status)
	}
	return nil
}
