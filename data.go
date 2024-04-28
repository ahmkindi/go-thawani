package thawani

import (
	"net/http"
	"net/url"

	"github.com/ahmkindi/go-thawani/types/mode"
	"github.com/ahmkindi/go-thawani/types/paymentstatus"
)

type ThawaniClient struct {
	HTTPClient     *http.Client
	BaseURL        *url.URL
	APIKey         string
	PublishableKey string
}

type Product struct {
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	UnitAmount int    `json:"unit_amount"`
}

type CreateSessionReq struct {
	ClientReferenceId string      `json:"client_reference_id"`
	Mode              mode.Type   `json:"mode"`
	Products          []Product   `json:"products"`
	SuccessUrl        string      `json:"success_url"`
	CancelUrl         string      `json:"cancel_url"`
	CustomerId        string      `json:"customer_id"`
	Metadata          interface{} `json:"metadata"`
}

type CustomerData struct {
	Id               string `json:"id"`
	CustomerClientId string `json:"customer_client_id"`
}

type SessionData struct {
	SessionId         string                 `json:"session_id"`
	ClientReferenceId string                 `json:"client_reference_id"`
	CustomerId        string                 `json:"customer_id"`
	Products          []Product              `json:"products"`
	TotalAmount       int                    `json:"total_amount"`
	PaymentStatus     paymentstatus.Type     `json:"payment_status"`
	ExpiresAt         string                 `json:"expire_at"`
	CreatedAt         string                 `json:"created_at"`
	Metadata          map[string]interface{} `json:"metadata"`
}

type BasicResponse struct {
	Success     bool   `json:"success"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type CreateCustomerReq struct {
	ClientCustomerId string `json:"client_customer_id"`
}

type CreateCustomerResp struct {
	BasicResponse
	Data CustomerData `json:"data"`
}

type Session struct {
	BasicResponse
	Data SessionData `json:"data"`
}
