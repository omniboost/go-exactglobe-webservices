package webservices_test

import (
	"encoding/json"
	"fmt"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestAccountRead(t *testing.T) {
	req := client.NewAccountReadRequest()
	req.QueryParams().Top.Set(1)
	// req.QueryParams().Filter.Set("ID eq guid'{e4e12d2a-409b-4447-93ad-360efc59a4e5}'")
	// req.QueryParams().Filter.Set("AccountCode eq '           80004153'")
	// req.QueryParams().Filter.Set("AccountCode eq '           965200000'")
	// req.QueryParams().Filter.Set("ID eq guid'{e4e12d2a-409b-4447-93ad-360efc59a4e5}'")
	// req.QueryParams().Select.Add("CompanyName")
	// req.QueryParams().Select.Add("FirstName")
	// req.QueryParams().Select.Add("LastName")
	// req.QueryParams().Select.Add("FullName")

	// req.PathParams().Date = time.Date(2018, 12, 5, 0, 0, 0, 0, time.UTC)
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
