package webservices_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGeneralLedgerAccountRead(t *testing.T) {
	req := client.NewGeneralLedgerAccountReadRequest()
	req.QueryParams().Top.Set(20)
	// req.QueryParams().Filter.Set("EntryNumber eq '20950144'")
	// req.QueryParams().Filter.Set("Journal eq ' 95' and GLAccount eq '140300002' and ReportingDate gt DateTime'2020-06-06T00:00:00Z'")

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
