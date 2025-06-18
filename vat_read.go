package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewVATReadRequest() VATReadRequest {
	return VATReadRequest{
		client:      c,
		queryParams: c.NewVATReadQueryParams(),
		pathParams:  c.NewVATReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewVATReadRequestBody(),
	}
}

type VATReadRequest struct {
	client      *Client
	queryParams *VATReadQueryParams
	pathParams  *VATReadPathParams
	method      string
	headers     http.Header
	requestBody VATReadRequestBody
}

func (c *Client) NewVATReadQueryParams() *VATReadQueryParams {
	selectFields, _ := utils.Fields(&CostCenter{})
	return &VATReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type VATReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p VATReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *VATReadRequest) QueryParams() *VATReadQueryParams {
	return r.queryParams
}

func (c *Client) NewVATReadPathParams() *VATReadPathParams {
	return &VATReadPathParams{}
}

type VATReadPathParams struct {
}

func (p *VATReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VATReadRequest) PathParams() *VATReadPathParams {
	return r.pathParams
}

func (r *VATReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *VATReadRequest) Method() string {
	return r.method
}

func (s *Client) NewVATReadRequestBody() VATReadRequestBody {
	return VATReadRequestBody{}
}

type VATReadRequestBody CostCenters

func (r *VATReadRequest) RequestBody() *VATReadRequestBody {
	return &r.requestBody
}

func (r *VATReadRequest) SetRequestBody(body VATReadRequestBody) {
	r.requestBody = body
}

func (r *VATReadRequest) NewResponseBody() *VATReadResponseBody {
	return &VATReadResponseBody{}
}

type VATReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			VAT
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *VATReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("VAT", r.PathParams())
}

func (r *VATReadRequest) Do() (VATReadResponseBody, error) {
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

type VAT struct {
	AccountDescription              string      `json:"AccountDescription"`
	AccountGuid                     string      `json:"AccountGuid"`
	AccountName                     string      `json:"AccountName"`
	AccountStatus                   string      `json:"AccountStatus"`
	AmountMaximum                   float64     `json:"AmountMaximum"`
	AmountMinimum                   float64     `json:"AmountMinimum"`
	Autofatturacode                 interface{} `json:"Autofatturacode,omitempty"`
	Calculationbasis                string      `json:"Calculationbasis"`
	CompanyCode                     string      `json:"CompanyCode"`
	CompanyName                     string      `json:"CompanyName"`
	CountryCode                     string      `json:"CountryCode"`
	CreatedDate                     string      `json:"CreatedDate"`
	Creator                         int         `json:"Creator"`
	CreatorFullName                 string      `json:"CreatorFullName"`
	Creditor                        string      `json:"Creditor"`
	CreditorCode                    string      `json:"CreditorCode"`
	Description                     string      `json:"Description"`
	Description1                    string      `json:"Description1"`
	Description2                    string      `json:"Description2"`
	Description3                    string      `json:"Description3"`
	Description4                    string      `json:"Description4"`
	Excluded                        string      `json:"Excluded"`
	ExtraDutyPercent                float64     `json:"ExtraDutyPercent"`
	GLToBeClaimed                   string      `json:"GLToBeClaimed,omitempty"`
	GLToBeClaimedDescription        string      `json:"GLToBeClaimedDescription,omitempty"`
	GLToClaim                       string      `json:"GLToClaim,omitempty"`
	GLToClaimDescription            string      `json:"GLToClaimDescription,omitempty"`
	GLToClaimSuspense               string      `json:"GLToClaimSuspense,omitempty"`
	GLToClaimSuspenseDescription    string      `json:"GLToClaimSuspenseDescription,omitempty"`
	GLToPay                         string      `json:"GLToPay,omitempty"`
	GLToPayDescription              string      `json:"GLToPayDescription,omitempty"`
	GLToPaySuspense                 string      `json:"GLToPaySuspense,omitempty"`
	GLToPaySuspenseDescription      string      `json:"GLToPaySuspenseDescription,omitempty"`
	ID                              int         `json:"ID"`
	InvoiceRegister                 int         `json:"InvoiceRegister"`
	InvestmentTaxAccount            string      `json:"InvestmentTaxAccount,omitempty"`
	InvestmentTaxBasisAccount       string      `json:"InvestmentTaxBasisAccount,omitempty"`
	InvestmentTaxBasisOffsetAccount string      `json:"InvestmentTaxBasisOffsetAccount,omitempty"`
	InvestmentTaxPercent            float64     `json:"InvestmentTaxPercent,omitempty"`
	ModifiedDate                    string      `json:"ModifiedDate"`
	Modifier                        int         `json:"Modifier"`
	ModifierFullName                string      `json:"ModifierFullName"`
	OffsetAccount                   string      `json:"OffsetAccount,omitempty"`
	PaymentPeriod                   string      `json:"PaymentPeriod,omitempty"`
	Percentage                      float64     `json:"Percentage"`
	PercentageDescription           string      `json:"PercentageDescription"`
	PerpertualService               bool        `json:"PerpertualService"`
	PurchaseType                    string      `json:"PurchaseType"`
	Rent                            string      `json:"Rent,omitempty"`
	RoundingScheme                  string      `json:"RoundingScheme,omitempty"`
	SalesPurchase                   string      `json:"SalesPurchase"`
	Taxkey                          string      `json:"Taxkey"`
	Taxsubkey                       string      `json:"Taxsubkey"`
	Taxtype                         string      `json:"Taxtype"`
	UseCashSystem                   bool        `json:"UseCashSystem,omitempty"`
	VATApplicable                   int         `json:"VATApplicable"`
	VATCharged                      bool        `json:"VATCharged,omitempty"`
	VATCode                         string      `json:"VATCode"`
	VATCodeCreditNote               string      `json:"VATCodeCreditNote,omitempty"`
	VATExemption                    bool        `json:"VATExemption,omitempty"`
	VATListing                      string      `json:"VATListing,omitempty"`
	VatNonDeductibleAccount         string      `json:"VatNonDeductibleAccount,omitempty"`
	VatNonDeductiblePercent         float64     `json:"VatNonDeductiblePercent,omitempty"`
	VATTransCodeCreditNote          string      `json:"VATTransCodeCreditNote,omitempty"`
}
