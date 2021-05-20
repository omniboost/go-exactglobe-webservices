package webservices_test

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/Azure/go-ntlmssp"
	webservices "github.com/omniboost/go-exactglobe-webservices"
	// ntlm "github.com/vadimi/go-http-ntlm"
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

	client = webservices.NewClient(nil, *baseURL, databaseName, databaseServerName)
	if debug != "" {
		client.SetDebug(true)
	}
	client.SetDisallowUnknownFields(true)
	client.SetBeforeRequestDo(func(httpClient *http.Client, req *http.Request, body interface{}) {
		tr := &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			ForceAttemptHTTP2: false,
		}

		ntlmTransport := ntlmssp.Negotiator{
			RoundTripper: tr,
		}

		httpClient.Transport = ntlmTransport

		// Set credentials
		req.SetBasicAuth(username, password)
	})
	m.Run()
}
