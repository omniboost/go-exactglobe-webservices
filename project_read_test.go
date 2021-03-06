package webservices_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestProjectRead(t *testing.T) {
	req := client.NewProjectReadRequest()
	req.QueryParams().Top.Set(1)
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
	log.Println(string(b))
}
