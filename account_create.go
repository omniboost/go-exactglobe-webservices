package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewAccountCreateRequest() AccountCreateRequest {
	return AccountCreateRequest{
		client:      c,
		queryParams: c.NewAccountCreateQueryParams(),
		pathParams:  c.NewAccountCreatePathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: AccountCreateRequestBody{},
	}
}

type AccountCreateRequest struct {
	client      *Client
	queryParams *AccountCreateQueryParams
	pathParams  *AccountCreatePathParams
	method      string
	headers     http.Header
	requestBody AccountCreateRequestBody
}

func (c *Client) NewAccountCreateQueryParams() *AccountCreateQueryParams {
	return &AccountCreateQueryParams{}
}

type AccountCreateQueryParams struct{}

func (p AccountCreateQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountCreateRequest) QueryParams() *AccountCreateQueryParams {
	return r.queryParams
}

func (c *Client) NewAccountCreatePathParams() *AccountCreatePathParams {
	return &AccountCreatePathParams{}
}

type AccountCreatePathParams struct{}

func (p *AccountCreatePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountCreateRequest) PathParams() *AccountCreatePathParams {
	return r.pathParams
}

func (r *AccountCreateRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountCreateRequest) Method() string {
	return r.method
}

func (r *AccountCreateRequest) NewRequestBody() AccountCreateRequestBody {
	return AccountCreateRequestBody{}
}

type AccountCreateRequestBody NewAccount

func (r *AccountCreateRequest) RequestBody() *AccountCreateRequestBody {
	return &r.requestBody
}

func (r *AccountCreateRequest) SetRequestBody(body AccountCreateRequestBody) {
	r.requestBody = body
}

func (r *AccountCreateRequest) NewResponseBody() *AccountCreateResponseBody {
	return &AccountCreateResponseBody{}
}

type AccountCreateResponseBody struct {
	D struct {
		Next     string       `json:"__next"`
		MetaData edm.MetaData `json:"__metadata"`

		Account
	} `json:"d"`
}

func (r *AccountCreateRequest) URL() url.URL {
	return r.client.GetEndpointURL("Account", r.PathParams())
}

func (r *AccountCreateRequest) Do() (AccountCreateResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

type NewAccount struct {
	AccountCode            string `json:"AccountCode"`
	AccountStatus          string `json:"AccountStatus"`
	CustomerType           string `json:"CustomerType"`
	CustomerStatus         string `json:"CustomerStatus"`
	AccountName            string `json:"AccountName"`
	CountryOfOrigin        string `json:"CountryOfOrigin"`
	CountryOfAssembly      string `json:"CountryOfAssembly"`
	PhoneNumber            string `json:"PhoneNumber"`
	FaxNumber              string `json:"FaxNumber"`
	ContactTitle           string `json:"ContactTitle,omitempty"`
	ContactFirstName       string `json:"ContactFirstName,omitempty"`
	ContactLastName        string `json:"ContactLastName,omitempty"`
	ContactEmail           string `json:"ContactEmail,omitempty"`
	ContactGender          string `json:"ContactGender,omitempty"`
	VATNumber              string `json:"VATNumber"`
	DeliveryAddressline1   string `json:"DeliveryAddressline1,omitempty"`
	DeliveryAddressline2   string `json:"DeliveryAddressline2,omitempty"`
	DeliveryAddressline3   string `json:"DeliveryAddressline3,omitempty"`
	DeliveryPostCode       string `json:"DeliveryPostCode,omitempty"`
	DeliveryCity           string `json:"DeliveryCity,omitempty"`
	DeliveryCountry        string `json:"DeliveryCountry,omitempty"`
	InvoiceAddress         string `json:"InvoiceAddress,omitempty"`
	InvoiceAddress2        string `json:"InvoiceAddress2,omitempty"`
	InvoiceAddress3        string `json:"InvoiceAddress3,omitempty"`
	InvoiceAddressPostCode string `json:"InvoiceAddressPostCode,omitempty"`
	InvoiceCity            string `json:"InvoiceCity,omitempty"`
	InvoiceCountry         string `json:"InvoiceCountry,omitempty"`
	MailingAddress1        string `json:"MailingAddress1,omitempty"`
	MailingAddress2        string `json:"MailingAddress2,omitempty"`
	MailingAddress3        string `json:"MailingAddress3,omitempty"`
	MailingPostcode        string `json:"MailingPostcode,omitempty"`
	MailingCity            string `json:"MailingCity,omitempty"`
	MailingCountry         string `json:"MailingCountry,omitempty"`
	PostalAddressLine1     string `json:"PostalAddressLine1,omitempty"`
	PostalAddressLine2     string `json:"PostalAddressLine2,omitempty"`
	PostalAddressLine3     string `json:"PostalAddressLine3,omitempty"`
	PostalPostCode         string `json:"PostalPostCode,omitempty"`
	PostalCity             string `json:"PostalCity,omitempty"`
	PostalCountry          string `json:"PostalCountry,omitempty"`
	DebtorNumber           string `json:"DebtorNumber,omitempty"`
	DebtorCode             string `json:"DebtorCode,omitempty"`
}
