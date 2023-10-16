package webservices_test

import (
	"encoding/json"
	"fmt"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGeneralLedgerAccountRead(t *testing.T) {
	req := client.NewGeneralLedgerAccountReadRequest()
	// req.QueryParams().Top.Set(1)
	req.QueryParams().Filter.Set("GeneralLedgerCode eq '140200003' or GeneralLedgerCode eq '140200050' or GeneralLedgerCode eq '140200090' or GeneralLedgerCode eq '190100050' or GeneralLedgerCode eq '199100050' or GeneralLedgerCode eq '199200000' or GeneralLedgerCode eq '199200051' or GeneralLedgerCode eq '199200052' or GeneralLedgerCode eq '370100001' or GeneralLedgerCode eq '370100002' or GeneralLedgerCode eq '370300000' or GeneralLedgerCode eq '520000000' or GeneralLedgerCode eq '612000000' or GeneralLedgerCode eq '801000000' or GeneralLedgerCode eq '805000000' or GeneralLedgerCode eq '807000000' or GeneralLedgerCode eq '807100000' or GeneralLedgerCode eq '808000000' or GeneralLedgerCode eq '811100000' or GeneralLedgerCode eq '831000000' or GeneralLedgerCode eq '832000000' or GeneralLedgerCode eq '833000000' or GeneralLedgerCode eq '841000000' or GeneralLedgerCode eq '841100000' or GeneralLedgerCode eq '841500000' or GeneralLedgerCode eq '842200000' or GeneralLedgerCode eq '843000000' or GeneralLedgerCode eq '843100000' or GeneralLedgerCode eq '843400000'")
	// req.QueryParams().Filter.Set("EntryNumber eq '20950144'")
	// req.QueryParams().Filter.Set("Journal eq '802' ReportingDate gt DateTime'2022-05-01T00:00:00Z'")

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
