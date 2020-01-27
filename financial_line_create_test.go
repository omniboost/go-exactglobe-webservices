package webservices_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestFinancialLineCreate(t *testing.T) {
	req := client.NewFinancialLineCreateRequest()

	rb := `{
		"TransactionKey": "",
		"Journal": "90",
		"GLAccount": "0010",
		"CostCenter": "VERKOOP",
		"CostUnit": "BINNEN",
		"Resource": "1034",
		"Project": "20003",
		"Amount": "100"
	}`

	err := json.Unmarshal([]byte(rb), req.RequestBody())
	if err != nil {
		t.Error(err)
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
