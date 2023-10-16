package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewFinancialLineReadRequest() FinancialLineReadRequest {
	return FinancialLineReadRequest{
		client:      c,
		queryParams: c.NewFinancialLineReadQueryParams(),
		pathParams:  c.NewFinancialLineReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewFinancialLineReadRequestBody(),
	}
}

type FinancialLineReadRequest struct {
	client      *Client
	queryParams *FinancialLineReadQueryParams
	pathParams  *FinancialLineReadPathParams
	method      string
	headers     http.Header
	requestBody FinancialLineReadRequestBody
}

func (c *Client) NewFinancialLineReadQueryParams() *FinancialLineReadQueryParams {
	selectFields, _ := utils.Fields(&FinancialLine{})
	return &FinancialLineReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type FinancialLineReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p FinancialLineReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FinancialLineReadRequest) QueryParams() *FinancialLineReadQueryParams {
	return r.queryParams
}

func (c *Client) NewFinancialLineReadPathParams() *FinancialLineReadPathParams {
	return &FinancialLineReadPathParams{}
}

type FinancialLineReadPathParams struct {
}

func (p *FinancialLineReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *FinancialLineReadRequest) PathParams() *FinancialLineReadPathParams {
	return r.pathParams
}

func (r *FinancialLineReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *FinancialLineReadRequest) Method() string {
	return r.method
}

func (s *Client) NewFinancialLineReadRequestBody() FinancialLineReadRequestBody {
	return FinancialLineReadRequestBody{}
}

type FinancialLineReadRequestBody FinancialLines

func (r *FinancialLineReadRequest) RequestBody() *FinancialLineReadRequestBody {
	return &r.requestBody
}

func (r *FinancialLineReadRequest) SetRequestBody(body FinancialLineReadRequestBody) {
	r.requestBody = body
}

func (r *FinancialLineReadRequest) NewResponseBody() *FinancialLineReadResponseBody {
	return &FinancialLineReadResponseBody{}
}

type FinancialLineReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			FinancialLine
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *FinancialLineReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("FinancialLine", r.PathParams())
}

func (r *FinancialLineReadRequest) Do() (FinancialLineReadResponseBody, error) {
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

type FinancialLines []FinancialLine

type FinancialLine struct {
	ID                    int         `json:"ID"`
	TransactionKey        interface{} `json:"TransactionKey"`
	FinancialYear         int         `json:"FinancialYear"`
	FinancialPeriod       int         `json:"FinancialPeriod"`
	SequenceNumber        int         `json:"SequenceNumber"`
	EntryNumber           string      `json:"EntryNumber"`
	Journal               string      `json:"Journal"`
	LineNumber            int         `json:"LineNumber"`
	EntryDate             edm.Date    `json:"EntryDate"`
	GLAccount             string      `json:"GLAccount"`
	Description           string      `json:"Description"`
	CurrencyCode          string      `json:"CurrencyCode"`
	CurrencyRate          int         `json:"CurrencyRate"`
	Amount                float64     `json:"Amount"`
	VATBasis              float64     `json:"VATBasis"`
	VATAmount             float64     `json:"VATAmount"`
	VATCode               string      `json:"VATCode"`
	CostCenter            string      `json:"CostCenter"`
	CostCenterDescription string      `json:"CostCenterDescription"`
	CostUnit              interface{} `json:"CostUnit"`
	CostUnitDescription   interface{} `json:"CostUnitDescription"`
	FulfillmentDate       edm.Date    `json:"FulfillmentDate"`
	DebtorNumber          interface{} `json:"DebtorNumber"`
	DebtorName            interface{} `json:"DebtorName"`
	CreditorNumber        interface{} `json:"CreditorNumber"`
	CreditorName          interface{} `json:"CreditorName"`
	Resource              int         `json:"Resource"`
	FullName              string      `json:"FullName"`
	OurReference          string      `json:"OurReference"`
	YourReference         string      `json:"YourReference"`
	Quantity              int         `json:"Quantity"`
	SerialNumber          interface{} `json:"SerialNumber"`
	ItemCode              string      `json:"ItemCode"`
	Project               interface{} `json:"Project"`
	ProjectDescription    interface{} `json:"ProjectDescription"`
	Warehouse             string      `json:"Warehouse"`
	WarehouseLocation     string      `json:"WarehouseLocation"`
	OrderNumber           interface{} `json:"OrderNumber"`
	ReportingDate         edm.Date    `json:"ReportingDate"`
	Textfield1            interface{} `json:"Textfield1"`
	Textfield2            interface{} `json:"Textfield2"`
	Textfield3            interface{} `json:"Textfield3"`
	NumberField1          int         `json:"NumberField1"`
	NumberField2          int         `json:"NumberField2"`
	TaxCode2              interface{} `json:"TaxCode2"`
	TaxCode3              interface{} `json:"TaxCode3"`
	TaxCode4              interface{} `json:"TaxCode4"`
	TaxCode5              interface{} `json:"TaxCode5"`
	TaxBasis2             int         `json:"TaxBasis2"`
	TaxBasis3             int         `json:"TaxBasis3"`
	TaxBasis4             int         `json:"TaxBasis4"`
	TaxBasis5             int         `json:"TaxBasis5"`
	TaxAmount2            int         `json:"TaxAmount2"`
	TaxAmount3            int         `json:"TaxAmount3"`
	TaxAmount4            int         `json:"TaxAmount4"`
	TaxAmount5            int         `json:"TaxAmount5"`
	OrderDebtor           interface{} `json:"OrderDebtor"`
	TransactionType       string      `json:"TransactionType"`
	EntryType             string      `json:"EntryType"`
	BankAccount           interface{} `json:"BankAccount"`
	PaymentMethod         interface{} `json:"PaymentMethod"`
	PaymentCondition      interface{} `json:"PaymentCondition"`
	PaymentReference      interface{} `json:"PaymentReference"`
	CSSDDueDate           edm.Date    `json:"CSSDDueDate"`
	CSSDDueDate2          edm.Date    `json:"CSSDDueDate2"`
	InvoiceDueDate        edm.Date    `json:"InvoiceDueDate"`
	TransactionOrigin     string      `json:"TransactionOrigin"`
	DocumentID            interface{} `json:"DocumentID"`
	DocAttachmentID       interface{} `json:"DocAttachmentID"`
	Shipment              interface{} `json:"Shipment"`
	UnitCode              string      `json:"UnitCode"`
	PriceList             interface{} `json:"PriceList"`
	Discount              interface{} `json:"Discount"`
	BankTransactionGUID   interface{} `json:"BankTransactionGUID"`
}
