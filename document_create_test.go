package webservices_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestDocumentCreate(t *testing.T) {
	req := client.NewDocumentCreateRequest()

	rb := `{
		"DocumentCode": "00000667",
		"DocumentName": "Leon Bogaert",
		"DocumentStatus": "S",
		"ContactEmail": "leon@omniboost.io",
		"ContactFirstName": "Leon",
		"ContactGender": "M",
		"ContactLastName": "Bogaert",
		"ContactTitle": "LS",
		"CustomerType": "C",
		"CustomerStatus": "A",
		"DebtorCode": "00000667",
		"DebtorNumber": "00000667",
		"FaxNumber": null,
		"InvoiceAddress": "Stadhuisplein 3",
		"InvoiceAddress2": "",
		"InvoiceAddress3": "",
		"InvoiceAddressPostCode": "",
		"InvoiceCity": "Terneuzen",
		"InvoiceCountry": "NL",
		"PhoneNumber": "",
		"VATNumber": null
	}`

	err := json.Unmarshal([]byte(rb), req.RequestBody())
	if err != nil {
		t.Error(err)
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}


