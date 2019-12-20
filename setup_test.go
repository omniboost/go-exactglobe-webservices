package webservices_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	webservices "github.com/omniboost/go-exactglobe-webservices"
)

var (
	client *webservices.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("GLOBE_BASE_URL")
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		log.Fatal(err)
	}

	databaseName := os.Getenv("GLOBE_DATABASE_NAME")
	databaseServerName := os.Getenv("GLOBE_DATABASE_SERVER_NAME")
	username := os.Getenv("GLOBE_USERNAME")
	password := os.Getenv("GLOBE_PASSWORD")
	debug := os.Getenv("GLOBE_DEBUG")
	log.Println(baseURLString, databaseName, username, password)

	client = webservices.NewClient(nil, *baseURL, databaseName, databaseServerName, username, password)
	if debug != "" {
		client.SetDebug(true)
	}
	client.SetDisallowUnknownFields(true)
	m.Run()
}
