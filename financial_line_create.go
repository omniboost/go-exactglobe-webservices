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
		TransactionKey        string      `json:"TransactionKey,omitempty"`
		FinancialYear         interface{} `json:"FinancialYear,omitempty"`
		FinancialPeriod       interface{} `json:"FinancialPeriod,omitempty"`
		SequenceNumber        interface{} `json:"SequenceNumber,omitempty"`
		EntryNumber           interface{} `json:"EntryNumber,omitempty"`
		Journal               string      `json:"Journal,omitempty"`
		LineNumber            interface{} `json:"LineNumber,omitempty"`
		EntryDate             interface{} `json:"EntryDate,omitempty"`
		GLAccount             string      `json:"GLAccount,omitempty"`
		Description           interface{} `json:"Description,omitempty"`
		CurrencyCode          interface{} `json:"CurrencyCode,omitempty"`
		CurrencyRate          interface{} `json:"CurrencyRate,omitempty"`
		Amount                float64     `json:"Amount,omitempty"`
		VATBasis              interface{} `json:"VATBasis,omitempty"`
		VATAmount             interface{} `json:"VATAmount,omitempty"`
		VATCode               interface{} `json:"VATCode,omitempty"`
		CostCenter            string      `json:"CostCenter,omitempty"`
		CostCenterDescription interface{} `json:"CostCenterDescription,omitempty"`
		CostUnit              string      `json:"CostUnit,omitempty"`
		CostUnitDescription   interface{} `json:"CostUnitDescription,omitempty"`
		FulfillmentDate       interface{} `json:"FulfillmentDate,omitempty"`
		DebtorNumber          string      `json:"DebtorNumber,omitempty"`
		DebtorName            interface{} `json:"DebtorName,omitempty"`
		CreditorNumber        interface{} `json:"CreditorNumber,omitempty"`
		CreditorName          interface{} `json:"CreditorName,omitempty"`
		Resource              int         `json:"Resource,omitempty"`
		FullName              interface{} `json:"FullName,omitempty"`
		OurReference          string      `json:"OurReference,omitempty"`
		YourReference         string      `json:"YourReference,omitempty"`
		Quantity              interface{} `json:"Quantity,omitempty"`
		SerialNumber          interface{} `json:"SerialNumber,omitempty"`
		ItemCode              interface{} `json:"ItemCode,omitempty"`
		Project               string      `json:"Project,omitempty"`
		ProjectDescription    interface{} `json:"ProjectDescription,omitempty"`
		Warehouse             interface{} `json:"Warehouse,omitempty"`
		WarehouseLocation     interface{} `json:"WarehouseLocation,omitempty"`
		OrderNumber           interface{} `json:"OrderNumber,omitempty"`
		ReportingDate         interface{} `json:"ReportingDate,omitempty"`
		Textfield1            interface{} `json:"Textfield1,omitempty"`
		Textfield2            interface{} `json:"Textfield2,omitempty"`
		Textfield3            interface{} `json:"Textfield3,omitempty"`
		NumberField1          interface{} `json:"NumberField1,omitempty"`
		NumberField2          interface{} `json:"NumberField2,omitempty"`
		TaxCode2              interface{} `json:"TaxCode2,omitempty"`
		TaxCode3              interface{} `json:"TaxCode3,omitempty"`
		TaxCode4              interface{} `json:"TaxCode4,omitempty"`
		TaxCode5              interface{} `json:"TaxCode5,omitempty"`
		TaxBasis2             interface{} `json:"TaxBasis2,omitempty"`
		TaxBasis3             interface{} `json:"TaxBasis3,omitempty"`
		TaxBasis4             interface{} `json:"TaxBasis4,omitempty"`
		TaxBasis5             interface{} `json:"TaxBasis5,omitempty"`
		TaxAmount2            interface{} `json:"TaxAmount2,omitempty"`
		TaxAmount3            interface{} `json:"TaxAmount3,omitempty"`
		TaxAmount4            interface{} `json:"TaxAmount4,omitempty"`
		TaxAmount5            interface{} `json:"TaxAmount5,omitempty"`
		OrderDebtor           interface{} `json:"OrderDebtor,omitempty"`
		TransactionType       interface{} `json:"TransactionType,omitempty"`
		EntryType             interface{} `json:"EntryType,omitempty"`
		BankAccount           interface{} `json:"BankAccount,omitempty"`
		PaymentMethod         interface{} `json:"PaymentMethod,omitempty"`
		PaymentCondition      interface{} `json:"PaymentCondition,omitempty"`
		PaymentReference      interface{} `json:"PaymentReference,omitempty"`
		CSSDDueDate           interface{} `json:"CSSDDueDate,omitempty"`
		CSSDDueDate2          interface{} `json:"CSSDDueDate2,omitempty"`
		InvoiceDueDate        interface{} `json:"InvoiceDueDate,omitempty"`
		TransactionOrigin     interface{} `json:"TransactionOrigin,omitempty"`
		DocumentID            interface{} `json:"DocumentID,omitempty"`
		DocAttachmentID       interface{} `json:"DocAttachmentID,omitempty"`
		Shipment              interface{} `json:"Shipment,omitempty"`
		UnitCode              interface{} `json:"UnitCode,omitempty"`
		PriceList             interface{} `json:"PriceList,omitempty"`
		Discount              interface{} `json:"Discount,omitempty"`
		BankTransactionGUID   interface{} `json:"BankTransactionGUID,omitempty"`
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
	TransactionKey string   `json:"TransactionKey,omitempty"`
	Journal        string   `json:"Journal"`
	GLAccount      string   `json:"GLAccount"`
	CostCenter     string   `json:"CostCenter,omitempty"`
	CostUnit       string   `json:"CostUnit,omitempty"`
	Resource       int      `json:"Resource"`
	Project        string   `json:"Project,omitempty"`
	Amount         float64  `json:"Amount"`
	DebtorNumber   string   `json:"DebtorNumber,omitempty"`
	CreditorNumber string   `json:"CreditorNumber,omitempty"`
	VATAmount      float64  `json:"VATAmount"`
	VATCode        string   `json:"VATCode"`
	Description    string   `json:"Description"`
	EntryDate      edm.Date `json:"EntryDate"`
	ReportingDate  edm.Date `json:"ReportingDate"`
	Quantity       float64  `json:"Quantity"`
	Textfield1     string   `json:"Textfield1,omitempty"`
	Textfield2     string   `json:"Textfield2,omitempty"`
	Textfield3     string   `json:"Textfield3,omitempty"`
	Textfield4     string   `json:"Textfield4,omitempty"`
	Textfield5     string   `json:"Textfield5,omitempty"`
	DocumentID     string   `json:"DocumentID,omitempty"`
}
