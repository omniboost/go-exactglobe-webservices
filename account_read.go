package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
	uuid "github.com/satori/go.uuid"
)

func (c *Client) NewAccountReadRequest() AccountReadRequest {
	return AccountReadRequest{
		client:      c,
		queryParams: c.NewAccountReadQueryParams(),
		pathParams:  c.NewAccountReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewAccountReadRequestBody(),
	}
}

type AccountReadRequest struct {
	client      *Client
	queryParams *AccountReadQueryParams
	pathParams  *AccountReadPathParams
	method      string
	headers     http.Header
	requestBody AccountReadRequestBody
}

func (c *Client) NewAccountReadQueryParams() *AccountReadQueryParams {
	selectFields, _ := utils.Fields(&Project{})
	return &AccountReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type AccountReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p AccountReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountReadRequest) QueryParams() *AccountReadQueryParams {
	return r.queryParams
}

func (c *Client) NewAccountReadPathParams() *AccountReadPathParams {
	return &AccountReadPathParams{}
}

type AccountReadPathParams struct {
}

func (p *AccountReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountReadRequest) PathParams() *AccountReadPathParams {
	return r.pathParams
}

func (r *AccountReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountReadRequest) Method() string {
	return r.method
}

func (s *Client) NewAccountReadRequestBody() AccountReadRequestBody {
	return AccountReadRequestBody{}
}

type AccountReadRequestBody struct{}

func (r *AccountReadRequest) RequestBody() *AccountReadRequestBody {
	return &r.requestBody
}

func (r *AccountReadRequest) SetRequestBody(body AccountReadRequestBody) {
	r.requestBody = body
}

func (r *AccountReadRequest) NewResponseBody() *AccountReadResponseBody {
	return &AccountReadResponseBody{}
}

type AccountReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			Account
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *AccountReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("Account", r.PathParams())
}

func (r *AccountReadRequest) Do() (AccountReadResponseBody, error) {
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

type Accounts []Account

type Account struct {
	AccountEmployeeId           int
	AccountEmployeeName         string
	AccountManagerID            int
	AccountManagerName          string
	AccountManagerType          string
	AddExtraReceiptToOrder      bool
	AddressID                   string
	CompanyCode                 string
	CompanyName                 string
	AllowBackorder              bool
	AllowPartialDelivery        bool
	ApplyShippingCharges        bool
	ASPDatabase                 string
	ASPDatabaseServer           string
	ASPWebsiteServer            string
	ASPWebsiteURL               string
	AttachUBL                   bool
	Attention                   bool
	AutoAllocate                bool
	AvalaraEntityUseCode        string
	BankAccount                 string
	BankAccountType             string
	BICCode                     string
	Blocked                     bool
	BOELimitAmount              int
	BOEMaxAmount                int
	Category_01                 string
	Category_02                 string
	Category_03                 string
	Category_04                 string
	Category_05                 string
	Category_06                 string
	Category_07                 string
	Category_08                 string
	Category_09                 string
	Category_10                 string
	Category_11                 string
	Category_12                 string
	Category_13                 string
	Category_14                 string
	Category_15                 string
	DebtorCreditorAccount       string
	DebtorAccountName           string
	ChamberofCommerceNumber     string
	ChangeVatData               bool
	ChangeVatServices           bool
	CheckDate                   *DateTime
	CityOfLoadUnload            string
	ClassificationDescription   string
	ClassificationID            string
	ContactAccountManager       int
	ContactCompanyCode          string
	ContactCreator              int
	ContactCreatedDate          *DateTime
	ContactModifiedDate         *DateTime
	ContactEmail                string
	ContactFaxNumber            string
	ContactFirstName            string
	ContactInitials             string
	JobDescription              string
	JobTitle                    string
	JobTitleDescription         string
	ContactLanguage             string
	ContactLanguageDescription  string
	ContactLastName             string
	ContactMiddleName           string
	ContactMobileNumber         string
	ContactModifier             int
	ContactPhoneNumber          string
	Extension                   string
	PictureFilename             string
	ContactTitle                string
	Abbreviation                string
	TitleDescription            string
	CiscoDirectory              string
	CollectionInvoice           string
	CompanySize                 string
	SizeDescription             string
	CompanyType                 int
	ContactFullName             string
	ContactID                   uuid.UUID
	ContactNumberFreeField1     float64
	ContactNumberFreeField2     float64
	ContactNumberFreeField3     float64
	ContactNumberFreeField4     float64
	ContactNumberFreeField5     float64
	CountryOfAssembly           string
	CountryOfOrigin             string
	CreditorCode                string
	CreditorNumber              string
	Creator                     int
	CreatorName                 string
	CreditabilityScenario       string
	CreditcardNumberReference   string
	CreditCardDescription       string
	CreditCardHolder            string
	CreditLine                  int
	CreditorNumberReference     string
	CreditorType                string
	CurrencyCode                string
	CurrencyDescription         string
	AccountCode                 string
	CodeatSupplier              string
	Source                      string
	SourceDescription           string
	CustomerSince               *DateTime
	CustomerStatus              string
	CustomerStatusDescription   string
	CustomerType                string
	CustomerTypeDescription     string
	CreatedDate                 DateTime
	DateFreeField1              *DateTime
	DateFreeField2              *DateTime
	DateFreeField3              *DateTime
	DateFreeField4              *DateTime
	DateFreeField5              *DateTime
	DateLastReminder            *DateTime
	ModifiedDate                DateTime
	DebtorCode                  string
	DebtorNumber                string
	DebtorNumberReference       string
	DefaultSelCode              string
	DefaultSelDescription       string
	DeliveryAddressline1        string
	DeliveryAddressline2        string
	DeliveryAddressline3        string
	DeliveryCity                string
	DeliveryCountry             string
	DeliveryCounty              string
	DeleteBankAccount           int
	DeliveryMethod              string
	DeliverymethodDescription   string
	DeliveryTerms               string
	DeliveryPostCode            string
	DeliveryState               string
	DestinationCountry          string
	Directory                   string
	Discount                    int
	SupplierID                  uuid.UUID
	CompanyCodeDebtorAccountID  *uuid.UUID
	CustomerID                  *uuid.UUID
	Document                    string
	DocumentCreator             string
	DocumentCreatedDate         *DateTime
	DocumentSubject             string
	DunBradstreetNumber         string
	EmailAddress                string
	ESFContactID                *uuid.UUID
	ExemptNumber                string
	CreditCardExpirationDate    string
	ExtraDuty                   bool
	FaxNumber                   string
	FedCategory                 string
	FedIDNumber                 string
	FedIDType                   string
	FiscalCode                  string
	ContactGender               string
	GLN                         string
	GroupDueDates               string
	GuidFreeField1              *uuid.UUID
	GuidFreeField2              *uuid.UUID
	GuidFreeField3              *uuid.UUID
	GuidFreeField4              *uuid.UUID
	GuidFreeField5              *uuid.UUID
	IBAN                        string
	ID                          string
	IncoTermAcknowledgeOrder    bool
	IncoTermCode                string
	IncoTermConfirmPrices       bool
	IncoTermProperty            string
	SectorCode                  string
	SectorDescription           string
	IntermediaryassociateID     *uuid.UUID
	IntrastatSystem             string
	IntRegion                   string
	InvoiceAddress              string
	InvoiceAddress2             string
	InvoiceAddress3             string
	InvoiceCity                 string
	InvoiceCountry              string
	InvoiceCounty               string
	InvoiceCopies               int
	InvoiceDebtor               string
	InvoiceMethod               string
	InvoiceAddressPostCode      string
	InvoiceState                string
	IsBinaryServiceEnabled      bool
	IsCommissionable            bool
	IsTaxExempted               bool
	ExtraItemDescription        string
	ExtraDescriptionName        string
	JournalCode                 string
	LogoFilename                string
	MailingAddress1             string
	MailingAddress2             string
	MailingAddress3             string
	Mailbox                     string
	MailingCity                 string
	MailingCountry              string
	MailingCountryDescription   string
	MailingCounty               string
	MailingPostcode             string
	MailingState                string
	ManagersCostCenter          string
	ManagersCCDescription       string
	MessageIDCompanyLogo        string
	MessageIDRemarks            string
	Modifier                    int
	ModifierName                string
	AccountName                 string
	Note                        interface{} `json:"Note"`
	NumberFreeField1            float64
	NumberFreeField2            float64
	NumberFreeField3            float64
	NumberFreeField4            float64
	NumberFreeField5            float64
	IntegerFreeField1           int
	IntegerFreeField2           int
	IntegerFreeField3           int
	IntegerFreeField4           int
	IntegerFreeField5           int
	OffsetLedgerAccount         string
	OffsetAccountName           string
	Parent                      string
	ParentCode                  string
	ParentName                  string
	PayeeName                   string
	PaymentCondition            string
	PaymentConditionDescription string
	PaymentDay1                 int
	PaymentDay2                 int
	PaymentDay3                 int
	PaymentDay4                 int
	PaymentDay5                 int
	PhoneNumber                 string
	PhoneExtension              string
	PhoneQueue                  string
	PostalAddressLine1          string
	PostalAddressLine2          string
	PostalAddressLine3          string
	PostalCity                  string
	PostalCountry               string
	PostalCounty                string
	PostalPostCode              string
	PostalState                 string
	PriceList                   string
	PriceListDescription        string
	PurchaseOrderConfirmation   bool
	OrderConfirmation           string
	Rating                      string
	RatingDescription           string
	Reminder                    bool
	ResellerCode                string
	ResellerID                  *uuid.UUID
	ResellerName                string
	SecurityLevel               int
	SignDate                    *DateTime
	MailingStateDescription     string
	StatementDate               *DateTime
	StatementFrequency          string
	StatementNumber             int
	StatementPrint              bool
	StatFactor                  string
	AccountStatus               string
	SubSector                   string
	SubSectorDescription        string
	SyncID                      *uuid.UUID
	Terms                       string
	TerritoryCode               string
	TextFreeField1              string
	TextFreeField10             string
	TextFreeField11             string
	TextFreeField12             string
	TextFreeField13             string
	TextFreeField14             string
	TextFreeField15             string
	TextFreeField16             string
	TextFreeField17             string
	TextFreeField18             string
	TextFreeField19             string
	TextFreeField2              string
	TextFreeField20             string
	TextFreeField21             string
	TextFreeField22             string
	TextFreeField23             string
	TextFreeField24             string
	TextFreeField25             string
	TextFreeField26             string
	TextFreeField27             string
	TextFreeField28             string
	TextFreeField29             string
	TextFreeField3              string
	TextFreeField30             string
	TextFreeField4              string
	TextFreeField5              string
	TextFreeField6              string
	TextFreeField7              string
	TextFreeField8              string
	TextFreeField9              string
	TransactionA                string
	TransactionB                string
	Transport                   string
	Transshipment               string
	TriangularCountry           string
	TypeSince                   *DateTime
	UseAvalaraTaxation          bool
	VatCode                     string
	VATLiability                string
	VATNumber                   string
	VatServices                 string
	WebAddress                  string
	YesNoFreeField1             bool
	YesNoFreeField2             bool
	YesNoFreeField3             bool
	YesNoFreeField4             bool
	YesNoFreeField5             bool
}
