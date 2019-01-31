package webservices_test

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	webservices "github.com/omniboost/go-exactglobe-webservices"
)

func TestProjectRead(t *testing.T) {
	baseURLString := os.Getenv("GLOBE_BASE_URL")
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		t.Error(err)
	}
	databaseName := os.Getenv("GLOBE_DATABASE_NAME")
	databaseServerName := os.Getenv("GLOBE_DATABASE_SERVER_NAME")
	username := os.Getenv("GLOBE_USERNAME")
	password := os.Getenv("GLOBE_PASSWORD")
	log.Println(baseURLString, databaseName, username, password)

	client := webservices.NewClient(nil, *baseURL, databaseName, databaseServerName, username, password)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

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
