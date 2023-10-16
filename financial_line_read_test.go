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
	req.QueryParams().Filter.Set("EntryNumber eq '30090281'")
	// req.QueryParams().Filter.Set("Journal eq '809' and ReportingDate gt DateTime'2022-04-30T00:00:00Z' and GLAccount eq '    97400' and Description eq 'Kamers OOO (002)'")

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
