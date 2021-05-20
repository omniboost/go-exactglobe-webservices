package webservices_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestFinancialHeaderCreate(t *testing.T) {
	lines := []string{
		`{
        "Amount": 3,
        "CostCenter": "",
        "CreditorNumber": null,
        "Description": "test",
        "EntryDate": "/Date(1581465600000)/",
        "FinancialPeriod": 6,
        "FinancialYear": 2020,
        "GLAccount": "140200001",
        "VATAmount": 0,
        "VATCode": "0"
		}`,
		`{
        "Amount": -3,
        "CostCenter": "005000",
        "CreditorNumber": null,
        "Description": "test",
        "EntryDate": "/Date(1581465600000)/",
        "FinancialPeriod": 6,
        "FinancialYear": 2020,
        "GLAccount": "370300000",
        "VATAmount": 0,
        "VATCode": "0"
		}`,
	}

	transactionKey := ""
	for _, line := range lines {
		lineRequest := client.NewFinancialLineCreateRequest()

		err := json.Unmarshal([]byte(line), lineRequest.RequestBody())
		lineRequest.RequestBody().TransactionKey = transactionKey
		if err != nil {
			t.Error(err)
			return
		}

		resp, err := lineRequest.Do()
		if err != nil {
			t.Error(err)
			return
		}

		if transactionKey == "" {
			transactionKey = resp.D.TransactionKey
		}
	}

	header := `{
		"TransactionKey": "",
		"Journal": "190"
	}`

	req := client.NewFinancialHeaderCreateRequest()
	err := json.Unmarshal([]byte(header), req.RequestBody())
	req.RequestBody().TransactionKey = transactionKey
	if err != nil {
		t.Error(err)
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
		return
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
