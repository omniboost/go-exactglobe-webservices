package webservices_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestFinancialHeaderRead(t *testing.T) {
	req := client.NewFinancialHeaderReadRequest()
	req.QueryParams().Top.Set(1000)
	// req.QueryParams().Filter.Set("ID eq 14639")
	// req.QueryParams().Filter.Set("Journal eq ' 95' and FinancialYear eq 2021 and FinancialPeriod eq 7")
	// req.QueryParams().Filter.Set("Journal eq '809' and ReportingDate gt DateTime'2022-05-16T00:00:00Z' and ReportingDate lt DateTime'2022-05-17T00:00:00Z'")
	req.QueryParams().Filter.Set("EntryNumber eq '23700001'")

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))

	// for _, r := range resp.D.Results {
	// 	req := client.NewFinancialLineReadRequest()
	// 	req.QueryParams().Top.Set(20)
	// 	req.QueryParams().Filter.Set(fmt.Sprintf("EntryNumber eq '%s'", r.EntryNumber))
	// 	// req.QueryParams().Filter.Set("Journal eq ' 95' and GLAccount eq '140300002' and ReportingDate gt DateTime'2020-06-06T00:00:00Z'")

	// 	resp, err := req.Do()
	// 	if err != nil {
	// 		t.Error(err)
	// 	}

	// 	b, _ := json.MarshalIndent(resp, "", "  ")
	// 	log.Println(string(b))
	// }
}
