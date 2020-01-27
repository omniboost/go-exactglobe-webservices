package webservices_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestFinancialLineRead(t *testing.T) {
	req := client.NewFinancialLineReadRequest()
	req.QueryParams().Top.Set(20)
	req.QueryParams().Filter.Set("EntryNumber eq '21900006'")

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
