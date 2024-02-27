package services

import (
	"bytes"
	"encoding/json"
	"errors"

	"net/http"
	"os"

	"github.com/mikaellpc4/go-payments-api/models"
)

type CreatePix struct {
	Platform models.Platform
	Amount   int
	Document string
	Name     string
}

type CreatePixResponse struct {
	Message string `json:"message"`
	Data    PixResponseData
}

type PixResponseData struct {
	QRCodeURL          string `json:"qr_code_url"`
	Favoured           string `json:"favoured"`
	ExternalReference  string `json:"external_reference"`
	ProcessorReference string `json:"processor_reference"`
	Checkout           string `json:"checkout"`
	PixKey             string `json:"pix_key"`
	Expiration         int    `json:"expiration"`
}

func CreatePIX(createPix CreatePix) (CreatePixResponse, error) {
	tenantId := os.Getenv("PAYMENT_TENANT_ID")

	body, _ := json.Marshal(map[string]interface{}{
		"value":     createPix.Amount,
		"name":      createPix.Name,
		"document":  createPix.Document,
		"tenant_id": tenantId,
	})

	apiUrl := os.Getenv("PAYMENT_API_URL")
	acessToken := os.Getenv("PAYMENT_ACCESS_TOKEN")
	company := os.Getenv("PAYMENT_COMPANY")
	email := os.Getenv("PAYMENT_EMAIL")

	request, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(body))
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("User-Agent", company+" ("+email+")")
	request.Header.Add("Authorization", acessToken)

	client := http.Client{}
	res, err := client.Do(request)

	var createPixResponse CreatePixResponse

	if err != nil {
		return createPixResponse, err
	}

	defer res.Body.Close()

	jsonError := json.NewDecoder(res.Body).Decode(&createPixResponse)

	if res.StatusCode != http.StatusOK {
		return createPixResponse, errors.New(createPixResponse.Message)
	}

	if jsonError != nil {
		return createPixResponse, jsonError
	}

	return createPixResponse, nil
}
