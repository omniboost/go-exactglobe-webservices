package webservices_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestFinancialHeaderRead(t *testing.T) {
	req := client.NewFinancialHeaderReadRequest()
	req.QueryParams().Top.Set(10)
	// req.QueryParams().Filter.Set("ID eq 14639")
	// req.QueryParams().Filter.Set("Journal eq '95'")

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
