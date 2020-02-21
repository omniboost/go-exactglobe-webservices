package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewFinancialLineCreateRequest() FinancialLineCreateRequest {
	return FinancialLineCreateRequest{
		client:      c,
		queryParams: c.NewFinancialLineCreateQueryParams(),
		pathParams:  c.NewFinancialLineCreatePathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewFinancialLineCreateRequestBody(),
	}
}

type FinancialLineCreateRequest struct {
	client      *Client
	queryParams *FinancialLineCreateQueryParams
	pathParams  *FinancialLineCreatePathParams
	method      string
	headers     http.Header
	requestBody FinancialLineCreateRequestBody
}

func (c *Client) NewFinancialLineCreateQueryParams() *FinancialLineCreateQueryParams {
	selectFields, _ := utils.Fields(&Project{})
	return &FinancialLineCreateQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type FinancialLineCreateQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p FinancialLineCreateQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FinancialLineCreateRequest) QueryParams() *FinancialLineCreateQueryParams {
	return r.queryParams
}

func (c *Client) NewFinancialLineCreatePathParams() *FinancialLineCreatePathParams {
	return &FinancialLineCreatePathParams{}
}

type FinancialLineCreatePathParams struct {
}

func (p *FinancialLineCreatePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *FinancialLineCreateRequest) PathParams() *FinancialLineCreatePathParams {
	return r.pathParams
}

func (r *FinancialLineCreateRequest) SetMethod(method string) {
	r.method = method
}

func (r *FinancialLineCreateRequest) Method() string {
	return r.method
}

func (s *Client) NewFinancialLineCreateRequestBody() FinancialLineCreateRequestBody {
	return FinancialLineCreateRequestBody{}
}

type FinancialLineCreateRequestBody NewFinancialLine

func (r *FinancialLineCreateRequest) RequestBody() *FinancialLineCreateRequestBody {
	return &r.requestBody
}

func (r *FinancialLineCreateRequest) SetRequestBody(body FinancialLineCreateRequestBody) {
	r.requestBody = body
}

func (r *FinancialLineCreateRequest) NewResponseBody() *FinancialLineCreateResponseBody {
	return &FinancialLineCreateResponseBody{}
}

type FinancialLineCreateResponseBody struct {
	D struct {
		Next     string       `json:"__next"`
		MetaData edm.MetaData `json:"__metadata"`

		ID                    int         `json:"ID,omitempty"`
		TransactionKey        string      `json:"TransactionKey"`
		FinancialYear         interface{} `json:"FinancialYear"`
		FinancialPeriod       interface{} `json:"FinancialPeriod"`
		SequenceNumber        interface{} `json:"SequenceNumber"`
		EntryNumber           interface{} `json:"EntryNumber"`
		Journal               string      `json:"Journal,omitempty"`
		LineNumber            interface{} `json:"LineNumber"`
		EntryDate             interface{} `json:"EntryDate"`
		GLAccount             string      `json:"GLAccount"`
		Description           interface{} `json:"Description"`
		CurrencyCode          interface{} `json:"CurrencyCode"`
		CurrencyRate          interface{} `json:"CurrencyRate"`
		Amount                float64     `json:"Amount"`
		VATBasis              interface{} `json:"VATBasis"`
		VATAmount             interface{} `json:"VATAmount"`
		VATCode               interface{} `json:"VATCode"`
		CostCenter            string      `json:"CostCenter"`
		CostCenterDescription interface{} `json:"CostCenterDescription"`
		CostUnit              string      `json:"CostUnit"`
		CostUnitDescription   interface{} `json:"CostUnitDescription"`
		FulfillmentDate       interface{} `json:"FulfillmentDate"`
		DebtorNumber          interface{} `json:"DebtorNumber"`
		DebtorName            interface{} `json:"DebtorName"`
		CreditorNumber        interface{} `json:"CreditorNumber"`
		CreditorName          interface{} `json:"CreditorName"`
		Resource              int         `json:"Resource"`
		FullName              interface{} `json:"FullName"`
		OurReference          string      `json:"OurReference"`
		YourReference         string      `json:"YourReference"`
		Quantity              interface{} `json:"Quantity"`
		SerialNumber          interface{} `json:"SerialNumber"`
		ItemCode              interface{} `json:"ItemCode"`
		Project               string      `json:"Project,omitempty"`
		ProjectDescription    interface{} `json:"ProjectDescription"`
		Warehouse             interface{} `json:"Warehouse"`
		WarehouseLocation     interface{} `json:"WarehouseLocation"`
		OrderNumber           interface{} `json:"OrderNumber"`
		ReportingDate         interface{} `json:"ReportingDate"`
		Textfield1            interface{} `json:"Textfield1"`
		Textfield2            interface{} `json:"Textfield2"`
		Textfield3            interface{} `json:"Textfield3"`
		NumberField1          interface{} `json:"NumberField1"`
		NumberField2          interface{} `json:"NumberField2"`
		TaxCode2              interface{} `json:"TaxCode2"`
		TaxCode3              interface{} `json:"TaxCode3"`
		TaxCode4              interface{} `json:"TaxCode4"`
		TaxCode5              interface{} `json:"TaxCode5"`
		TaxBasis2             interface{} `json:"TaxBasis2"`
		TaxBasis3             interface{} `json:"TaxBasis3"`
		TaxBasis4             interface{} `json:"TaxBasis4"`
		TaxBasis5             interface{} `json:"TaxBasis5"`
		TaxAmount2            interface{} `json:"TaxAmount2"`
		TaxAmount3            interface{} `json:"TaxAmount3"`
		TaxAmount4            interface{} `json:"TaxAmount4"`
		TaxAmount5            interface{} `json:"TaxAmount5"`
		OrderDebtor           interface{} `json:"OrderDebtor"`
		TransactionType       interface{} `json:"TransactionType"`
		EntryType             interface{} `json:"EntryType"`
		BankAccount           interface{} `json:"BankAccount"`
		PaymentMethod         interface{} `json:"PaymentMethod"`
		PaymentCondition      interface{} `json:"PaymentCondition"`
		PaymentReference      interface{} `json:"PaymentReference"`
		CSSDDueDate           interface{} `json:"CSSDDueDate"`
		CSSDDueDate2          interface{} `json:"CSSDDueDate2"`
		InvoiceDueDate        interface{} `json:"InvoiceDueDate"`
		TransactionOrigin     interface{} `json:"TransactionOrigin"`
		DocumentID            interface{} `json:"DocumentID"`
		DocAttachmentID       interface{} `json:"DocAttachmentID"`
		Shipment              interface{} `json:"Shipment"`
		UnitCode              interface{} `json:"UnitCode"`
		PriceList             interface{} `json:"PriceList"`
		Discount              interface{} `json:"Discount"`
		BankTransactionGUID   interface{} `json:"BankTransactionGUID"`
	} `json:"d"`
}

func (r *FinancialLineCreateRequest) URL() url.URL {
	return r.client.GetEndpointURL("FinancialLine", r.PathParams())
}

func (r *FinancialLineCreateRequest) Do() (FinancialLineCreateResponseBody, error) {
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

type NewFinancialLine struct {
	TransactionKey string       `json:"TransactionKey,omitempty"`
	Journal        string       `json:"Journal"`
	GLAccount      string       `json:"GLAccount"`
	CostCenter     string       `json:"CostCenter,omitempty"`
	CostUnit       string       `json:"CostUnit,omitempty"`
	Resource       int          `json:"Resource"`
	Project        string       `json:"Project,omitempty"`
	Amount         float64      `json:"Amount"`
	DebtorNumber   string       `json:"DebtorNumber,omitempty"`
	CreditorNumber string       `json:"CreditorNumber,omitempty"`
	VATAmount      float64      `json:"VATAmount"`
	VATCode        string       `json:"VATCode"`
	Description    string       `json:"Description"`
	EntryDate      edm.DateTime `json:"EntryDate"`
}
