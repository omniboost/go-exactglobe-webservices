package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewFinancialHeaderReadRequest() FinancialHeaderReadRequest {
	return FinancialHeaderReadRequest{
		client:      c,
		queryParams: c.NewFinancialHeaderReadQueryParams(),
		pathParams:  c.NewFinancialHeaderReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewFinancialHeaderReadRequestBody(),
	}
}

type FinancialHeaderReadRequest struct {
	client      *Client
	queryParams *FinancialHeaderReadQueryParams
	pathParams  *FinancialHeaderReadPathParams
	method      string
	headers     http.Header
	requestBody FinancialHeaderReadRequestBody
}

func (c *Client) NewFinancialHeaderReadQueryParams() *FinancialHeaderReadQueryParams {
	selectFields, _ := utils.Fields(&FinancialHeader{})
	return &FinancialHeaderReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type FinancialHeaderReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p FinancialHeaderReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FinancialHeaderReadRequest) QueryParams() *FinancialHeaderReadQueryParams {
	return r.queryParams
}

func (c *Client) NewFinancialHeaderReadPathParams() *FinancialHeaderReadPathParams {
	return &FinancialHeaderReadPathParams{}
}

type FinancialHeaderReadPathParams struct {
}

func (p *FinancialHeaderReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *FinancialHeaderReadRequest) PathParams() *FinancialHeaderReadPathParams {
	return r.pathParams
}

func (r *FinancialHeaderReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *FinancialHeaderReadRequest) Method() string {
	return r.method
}

func (s *Client) NewFinancialHeaderReadRequestBody() FinancialHeaderReadRequestBody {
	return FinancialHeaderReadRequestBody{}
}

type FinancialHeaderReadRequestBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			FinancialHeader
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *FinancialHeaderReadRequest) RequestBody() *FinancialHeaderReadRequestBody {
	return &r.requestBody
}

func (r *FinancialHeaderReadRequest) SetRequestBody(body FinancialHeaderReadRequestBody) {
	r.requestBody = body
}

func (r *FinancialHeaderReadRequest) NewResponseBody() *FinancialHeaderReadResponseBody {
	return &FinancialHeaderReadResponseBody{}
}

type FinancialHeaderReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			FinancialHeader
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *FinancialHeaderReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("FinancialHeader", r.PathParams())
}

func (r *FinancialHeaderReadRequest) Do() (FinancialHeaderReadResponseBody, error) {
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

type FinancialHeader struct {
	ID                     int         `json:"ID"`
	FinancialYear          int         `json:"FinancialYear"`
	FinancialPeriod        int         `json:"FinancialPeriod"`
	Journal                string      `json:"Journal"`
	JournalType            string      `json:"JournalType"`
	JournalTypeDescription string      `json:"JournalTypeDescription"`
	JournalAccount         string      `json:"JournalAccount"`
	SequenceNumber         string      `json:"SequenceNumber"`
	EntryNumber            string      `json:"EntryNumber"`
	EntryDate              edm.Date    `json:"EntryDate"`
	ReportingDate          edm.Date    `json:"ReportingDate"`
	Description            string      `json:"Description"`
	CostCenter             interface{} `json:"CostCenter"`
	CostCenterDescription  interface{} `json:"CostCenterDescription"`
	CostUnit               interface{} `json:"CostUnit"`
	CostUnitDescription    interface{} `json:"CostUnitDescription"`
	Resource               int         `json:"Resource"`
	FullName               interface{} `json:"FullName"`
	DebtorNumber           interface{} `json:"DebtorNumber"`
	DebtorName             interface{} `json:"DebtorName"`
	CreditorNumber         interface{} `json:"CreditorNumber"`
	Creditorname           interface{} `json:"Creditorname"`
	PaymentMethod          interface{} `json:"PaymentMethod"`
	PaymentCondition       interface{} `json:"PaymentCondition"`
	SelectionCode          interface{} `json:"SelectionCode"`
	OurReference           string      `json:"OurReference"`
	YourReference          string      `json:"YourReference"`
	FulfillmentDate        string      `json:"FulfillmentDate"`
	InvoiceDueDate         interface{} `json:"InvoiceDueDate"`
	CurrencyCode           string      `json:"CurrencyCode"`
	CurrencyRate           int         `json:"CurrencyRate"`
	AngloSaxonNotation     interface{} `json:"AngloSaxonNotation"`
	Amount                 float64     `json:"Amount"`
	Project                interface{} `json:"Project"`
	ProjectDescription     interface{} `json:"ProjectDescription"`
	PaymentReference       interface{} `json:"PaymentReference"`
	Status                 string      `json:"Status"`
	StatusDescription      string      `json:"StatusDescription"`
	TextField1             interface{} `json:"TextField1"`
	TextField2             interface{} `json:"TextField2"`
	TextField3             interface{} `json:"TextField3"`
	NumberField1           int         `json:"NumberField1"`
	NumberField2           int         `json:"NumberField2"`
	CSSDDueDate            interface{} `json:"CSSDDueDate"`
	CSSDDueDate2           interface{} `json:"CSSDDueDate2"`
	CSSDAmount             string      `json:"CSSDAmount"`
	CSSDAmount1            int         `json:"CSSDAmount1"`
	CSSDAmount2            int         `json:"CSSDAmount2"`
	BankAccount            interface{} `json:"BankAccount"`
	OpeningBalance         int         `json:"OpeningBalance"`
	ClosingBalance         int         `json:"ClosingBalance"`
	TransactionKey         interface{} `json:"TransactionKey"`
	EntryLines             string      `json:"EntryLines"`
	PaymentTermLines       string      `json:"PaymentTermLines"`
	BankStatementLines     string      `json:"BankStatementLines"`
	BankStatementNumber    interface{} `json:"BankStatementNumber"`
	PaymentStatus          string      `json:"PaymentStatus"`
	PaymentReceived        float64     `json:"PaymentReceived"`
	PaymentReceivedTC      float64     `json:"PaymentReceivedTC"`
	OutstandingAmount      float64     `json:"OutstandingAmount"`
	OutstandingAmountTC    float64     `json:"OutstandingAmountTC"`
	MandateReference       interface{} `json:"MandateReference"`
	SysGUID                string      `json:"SysGuid"`
	VoidReason             interface{} `json:"VoidReason"`
}
