package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
	uuid "github.com/satori/go.uuid"
)

func (c *Client) NewResourceReadRequest() ResourceReadRequest {
	return ResourceReadRequest{
		client:      c,
		queryParams: c.NewResourceReadQueryParams(),
		pathParams:  c.NewResourceReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewResourceReadRequestBody(),
	}
}

type ResourceReadRequest struct {
	client      *Client
	queryParams *ResourceReadQueryParams
	pathParams  *ResourceReadPathParams
	method      string
	headers     http.Header
	requestBody ResourceReadRequestBody
}

func (c *Client) NewResourceReadQueryParams() *ResourceReadQueryParams {
	selectFields, _ := utils.Fields(&Resource{})
	return &ResourceReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type ResourceReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p ResourceReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ResourceReadRequest) QueryParams() *ResourceReadQueryParams {
	return r.queryParams
}

func (c *Client) NewResourceReadPathParams() *ResourceReadPathParams {
	return &ResourceReadPathParams{}
}

type ResourceReadPathParams struct {
}

func (p *ResourceReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ResourceReadRequest) PathParams() *ResourceReadPathParams {
	return r.pathParams
}

func (r *ResourceReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *ResourceReadRequest) Method() string {
	return r.method
}

func (s *Client) NewResourceReadRequestBody() ResourceReadRequestBody {
	return ResourceReadRequestBody{}
}

type ResourceReadRequestBody struct{}

func (r *ResourceReadRequest) RequestBody() *ResourceReadRequestBody {
	return &r.requestBody
}

func (r *ResourceReadRequest) SetRequestBody(body ResourceReadRequestBody) {
	r.requestBody = body
}

func (r *ResourceReadRequest) NewResponseBody() *ResourceReadResponseBody {
	return &ResourceReadResponseBody{}
}

type ResourceReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			Resource
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *ResourceReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("Resource", r.PathParams())
}

func (r *ResourceReadRequest) Do() (ResourceReadResponseBody, error) {
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

type Resource struct {
	Blocked                     bool       `json:"Blocked"`
	AddressLine1                string     `json:"AddressLine1"`
	AddressLine2                string     `json:"AddressLine2"`
	AddressHouseNumber          string     `json:"AddressHouseNumber"`
	AdjustedHireDate            *DateTime  `json:"AdjustedHireDate"`
	Suffix                      string     `json:"Suffix"`
	ApplicationPhase            string     `json:"ApplicationPhase"`
	AssistantFullName           string     `json:"AssistantFullName"`
	AssistantID                 int        `json:"AssistantID"`
	AccountAddressID            *uuid.UUID `json:"AccountAddressID"`
	AccountContactID            *uuid.UUID `json:""AccountContactID"`
	AccountID                   int        `json:"AccountID"`
	BackOfficeBlocked           bool       `json:"BackOfficeBlocked"`
	BankAccount1                string     `json:"BankAccount1"`
	BirthDate                   *DateTime  `json:"BirthDate"`
	BirthPrefix                 string     `json:"BirthPrefix"`
	CentralizationAccount       string     `json:"CentralizationAccount"`
	City                        string     `json:"City"`
	Classification              string     `json:"Classification"`
	ClassificationID            *uuid.UUID `json:"ClassificationID"`
	CreditorID                  *uuid.UUID `json:"CreditorID"`
	CompanyCode                 string     `json:"CompanyCode"`
	CompanyName                 string     `json:"CompanyName"`
	ContractEnd                 *DateTime  `json:"ContractEnd"`
	CostCenter                  string     `json:"CostCenter"`
	CostCenterDescription       string     `json:"CostCenterDescription"`
	CostUnit                    string     `json:"CostUnit"`
	CostUnitDescription         string     `json:"CostUnitDescription"`
	CountryCode                 string     `json:"CountryCode"`
	CountryDescription          string     `json:"CountryDescription"`
	CreatedDate                 DateTime   `json:"CreatedDate"`
	Creator                     int        `json:"Creator"`
	CreatorName                 string     `json:"CreatorName"`
	CreditLine                  int        `json:"CreditLine"`
	AccountPayable              string     `json:"AccountPayable"`
	AccountPayableName          string     `json:"AccountPayableName"`
	AccountPayableNumber        string     `json:"AccountPayableNumber"`
	Currency                    string     `json:"Currency"`
	CurrencyDescription         string
	DateFreeField1              *DateTime
	DateFreeField2              *DateTime
	DateFreeField3              *DateTime
	DateFreeField4              *DateTime
	DateFreeField5              *DateTime
	DebtorCode                  string
	DebtorNumber                string
	Discount                    int
	EmailPersonal               string
	EmailWork                   string
	EmploymentEndDate           *DateTime
	EmploymentStartDate         DateTime
	PhoneWorkExtension          string
	Fax                         string
	FirstName                   string
	FiscalCode                  string
	FTE                         float64
	FullName                    string
	Gender                      string
	HRContractID                *uuid.UUID
	ID                          int
	Initials                    string
	InternalRate                int
	IsBinaryServiceEnabled      string
	ISPAccount                  string
	ItemCode                    string
	JobCategory                 string
	JobGroup                    string
	SecurityLevel               int
	JobTitle                    string
	JobTitleDescription         string
	JobTitleLevel               int
	Langua                      string
	LanguageDescription         string
	LastName                    string
	PhysicalLocation            string
	PhysicalLocationDescription string
	MaidenNa                    string
	ContractLocati              string
	ContractLocationDescription string
	ManagerFullName             string
	ManagerID                   int
	ManagerUserID               *uuid.UUID
	MaritalState                string
	MaritalStateDate            string
	MessageIDPicture            string
	MessageIDSignature          string
	MiddleName                  string
	ModifiedDate                DateTime
	Modifier                    int
	ModifierName                string
	MsnID                       string
	Nationality                 string
	NationalityDescription      string
	NumberFreeField1            float64
	NumberFreeField2            float64
	NumberFreeField3            float64
	NumberFreeField4            float64
	NumberFreeField5            float64
	OffSetAccount               string
	OriginCountryCode           string
	OriginCountryDescription    string
	PayeeName                   string
	PaymentCondition            string
	PaymentConditionDescription string
	PaymentLimit                int
	PaymentMethod               string
	PhonePersonal               string
	PhoneWorkMobile             string
	PhonePersonalMobile         string
	PhoneWorkMobileExtension    string
	PhoneWork                   string
	PictureFilename             string
	PlaceOfBirth                string
	PostalAddressLine1          string
	PostalAddressLine2          string
	PostalCity                  string
	PostalCountry               string
	PostalCountryDescription    string
	PostalKeepSameAsVisit       bool
	PostalPhone                 string
	PostalPostCode              string
	PostalStateCode             string
	PostalStateDescription      string
	Postcode                    string
	Prefix                      string
	PrevRating                  string
	PriceList                   string
	PurchaseCurrency            string
	PurchaseLimit               int
	Qualifications              string
	Race                        string
	Rating                      string
	ReasonResign                string
	RMALimit                    int
	Room                        string
	SalesInvoiceLimit           int
	SalesOrderLimit             int
	JobLevel                    int
	SignatureFile               string
	SkypeID                     string
	SocialSecurityNumber        string
	State                       string
	StateDescription            string
	Status                      string
	sysguid                     uuid.UUID
	Focus                       string
	TextFreeField1              string
	TextFreeField2              string
	TextFreeField3              string
	TextFreeField4              string
	TextFreeField5              string
	TimeZoneDescription         string
	TimeZone                    string
	Title                       string
	TitleDescription            string
	ResourceType                string
	UserID                      string
	VacancyStartDate            *DateTime
	VacancyType                 string
	VatNumber                   string
	YesNoFreeField1             bool
	YesNoFreeField2             bool
	YesNoFreeField3             bool
	YesNoFreeField4             bool
	YesNoFreeField5             bool
}
