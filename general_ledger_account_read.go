package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
	uuid "github.com/satori/go.uuid"
)

func (c *Client) NewGeneralLedgerAccountReadRequest() GeneralLedgerAccountReadRequest {
	return GeneralLedgerAccountReadRequest{
		client:      c,
		queryParams: c.NewGeneralLedgerAccountReadQueryParams(),
		pathParams:  c.NewGeneralLedgerAccountReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGeneralLedgerAccountReadRequestBody(),
	}
}

type GeneralLedgerAccountReadRequest struct {
	client      *Client
	queryParams *GeneralLedgerAccountReadQueryParams
	pathParams  *GeneralLedgerAccountReadPathParams
	method      string
	headers     http.Header
	requestBody GeneralLedgerAccountReadRequestBody
}

func (c *Client) NewGeneralLedgerAccountReadQueryParams() *GeneralLedgerAccountReadQueryParams {
	selectFields, _ := utils.Fields(&GeneralLedgerAccount{})
	return &GeneralLedgerAccountReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type GeneralLedgerAccountReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p GeneralLedgerAccountReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GeneralLedgerAccountReadRequest) QueryParams() *GeneralLedgerAccountReadQueryParams {
	return r.queryParams
}

func (c *Client) NewGeneralLedgerAccountReadPathParams() *GeneralLedgerAccountReadPathParams {
	return &GeneralLedgerAccountReadPathParams{}
}

type GeneralLedgerAccountReadPathParams struct {
}

func (p *GeneralLedgerAccountReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GeneralLedgerAccountReadRequest) PathParams() *GeneralLedgerAccountReadPathParams {
	return r.pathParams
}

func (r *GeneralLedgerAccountReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *GeneralLedgerAccountReadRequest) Method() string {
	return r.method
}

func (s *Client) NewGeneralLedgerAccountReadRequestBody() GeneralLedgerAccountReadRequestBody {
	return GeneralLedgerAccountReadRequestBody{}
}

type GeneralLedgerAccountReadRequestBody GeneralLedgerAccounts

func (r *GeneralLedgerAccountReadRequest) RequestBody() *GeneralLedgerAccountReadRequestBody {
	return &r.requestBody
}

func (r *GeneralLedgerAccountReadRequest) SetRequestBody(body GeneralLedgerAccountReadRequestBody) {
	r.requestBody = body
}

func (r *GeneralLedgerAccountReadRequest) NewResponseBody() *GeneralLedgerAccountReadResponseBody {
	return &GeneralLedgerAccountReadResponseBody{}
}

type GeneralLedgerAccountReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			GeneralLedgerAccount
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *GeneralLedgerAccountReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("GeneralLedgerAccount", r.PathParams())
}

func (r *GeneralLedgerAccountReadRequest) Do() (GeneralLedgerAccountReadResponseBody, error) {
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

type GeneralLedgerAccounts []GeneralLedgerAccount

type GeneralLedgerAccount struct {
	AccountConversionType             int
	AccountReportCategory             string
	AdjustmentAccount                 string
	AlternativeLedger                 string
	Assets                            bool
	BalanceProfitAndLoss              string
	Blocked                           bool
	CentralizationAccount             string
	ChequeAccount                     bool
	Class1                            string
	Class1Description                 string
	Class2                            string
	Class2Description                 string
	Class3                            string
	Class3Description                 string
	Class4                            string
	Class4Description                 string
	Class5                            string
	Class5Description                 string
	Class6                            string
	Class6Description                 string
	Class7                            string
	Class7Description                 string
	Class8                            string
	Class8Description                 string
	Class9                            string
	Class9Description                 string
	Class10                           string
	Class10Description                string
	Class11                           string
	Class12                           string
	ClosingEntryAccount               string
	Compress                          bool
	CorporateGeneralLedgerCode        string
	CorporateGeneralLedgerDescription string
	CostCentreAccount                 bool
	CostUnitAccount                   bool
	CreatedDate                       edm.DateTime
	Creator                           int
	CreatorFullName                   string
	DebitCredit                       string
	DefaultCostCentre                 string
	DefaultCostUnit                   string
	DefaultDebtorCreditor             bool
	CompanyCode                       string
	CompanyName                       string
	DocumentID                        *uuid.UUID
	DocumentSubject                   string
	ExcludeVATListing                 bool
	ExternalBalanceCredit             string
	TextFreeField1                    string
	TextFreeField2                    string
	TextFreeField3                    string
	TextFreeField4                    string
	TextFreeField5                    string
	TextFreeField6                    string
	TextFreeField7                    string
	TextFreeField8                    string
	TextFreeField9                    string
	TextFreeField10                   string
	AmountFreeField1                  float64
	AmountFreeField2                  float64
	AmountFreeField3                  float64
	AmountFreeField4                  float64
	AmountFreeField5                  float64
	YesNoFreeField1                   bool
	YesNoFreeField2                   bool
	YesNoFreeField3                   bool
	YesNoFreeField4                   bool
	YesNoFreeField5                   bool
	GeneralLedgerAccountNumber        string
	GLAssociate                       string
	GeneralLedgerCode                 string
	GeneralLedgerDescription          string
	GeneralLedgerDescription1         string
	GeneralLedgerDescription2         string
	GeneralLedgerDescription3         string
	GeneralLedgerDescription4         string
	ID                                int
	InvoiceRegisterType               string
	ItemAccount                       bool
	Match                             bool
	ModifiedDate                      edm.DateTime
	Modifier                          int
	ModifierFullName                  string
	PageAfterClosing                  int
	PercentagePrivate                 float64
	PersonalAccount                   bool
	ProjectAccount                    bool
	PurchaseVATReturnType             string
	Quantities                        bool
	RegisterMinerals                  string
	ResourceAccount                   bool
	Revalue                           bool
	RevenueAccount                    string
	RewardType                        string
	SchemeType                        string
	ShowNotes                         bool
	SubTotalAccount                   string
	TypeAdjustmentInflation           string
	UseCostcenterAllocation           bool
	UseIntercompany                   bool
	VATCode                           string
	VATNonDeductibleAccount           string
	VatNonDeductiblePercent           float64
}
