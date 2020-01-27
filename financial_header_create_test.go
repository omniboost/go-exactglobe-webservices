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
			"TransactionKey": "",
			"Journal": "90",
			"GLAccount": "0010",
			"CostCenter": "VERKOOP",
			"CostUnit": "BINNEN",
			"Amount": "100"
		}`,
		`{
			"TransactionKey": "",
			"Journal": "90",
			"GLAccount": "0010",
			"CostCenter": "VERKOOP",
			"CostUnit": "BINNEN",
			"Amount": "-100"
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
		"Journal": "90"
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
