package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewDocumentReadRequest() DocumentReadRequest {
	return DocumentReadRequest{
		client:      c,
		queryParams: c.NewDocumentReadQueryParams(),
		pathParams:  c.NewDocumentReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewDocumentReadRequestBody(),
	}
}

type DocumentReadRequest struct {
	client      *Client
	queryParams *DocumentReadQueryParams
	pathParams  *DocumentReadPathParams
	method      string
	headers     http.Header
	requestBody DocumentReadRequestBody
}

func (c *Client) NewDocumentReadQueryParams() *DocumentReadQueryParams {
	selectFields, _ := utils.Fields(&Project{})
	return &DocumentReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type DocumentReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p DocumentReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DocumentReadRequest) QueryParams() *DocumentReadQueryParams {
	return r.queryParams
}

func (c *Client) NewDocumentReadPathParams() *DocumentReadPathParams {
	return &DocumentReadPathParams{}
}

type DocumentReadPathParams struct {
}

func (p *DocumentReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DocumentReadRequest) PathParams() *DocumentReadPathParams {
	return r.pathParams
}

func (r *DocumentReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *DocumentReadRequest) Method() string {
	return r.method
}

func (s *Client) NewDocumentReadRequestBody() DocumentReadRequestBody {
	return DocumentReadRequestBody{}
}

type DocumentReadRequestBody struct{}

func (r *DocumentReadRequest) RequestBody() *DocumentReadRequestBody {
	return &r.requestBody
}

func (r *DocumentReadRequest) SetRequestBody(body DocumentReadRequestBody) {
	r.requestBody = body
}

func (r *DocumentReadRequest) NewResponseBody() *DocumentReadResponseBody {
	return &DocumentReadResponseBody{}
}

type DocumentReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			Document
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *DocumentReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("Document", r.PathParams())
}

func (r *DocumentReadRequest) Do() (DocumentReadResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
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

type Documents []Document

type Document struct {
	edm.MetaData `json:"__metadata,omitempty"`

	AccountContactEmail     string   `json:"AccountContactEmail,omitempty"`
	AccountContactID        string   `json:"AccountContactID,omitempty"`
	AccountContactName      string   `json:"AccountContactName,omitempty"`
	AllowPublish            bool     `json:"AllowPublish,omitempty"`
	AllowReplies            bool     `json:"AllowReplies,omitempty"`
	AccountCode             string   `json:"AccountCode,omitempty"`
	AccountID               string   `json:"AccountID,omitempty"`
	AccountEmail            string   `json:"AccountEmail,omitempty"`
	AccountName             string   `json:"AccountName,omitempty"`
	Category                string   `json:"Category,omitempty"`
	CategoryID              string   `json:"CategoryID,omitempty"`
	CheckedOut              edm.Date `json:"CheckedOut,omitempty"`
	CheckedOutBy            string   `json:"CheckedOutBy,omitempty"`
	CheckedOutByName        string   `json:"CheckedOutByName,omitempty"`
	Assortment              float64  `json:"Assortment,omitempty"`
	Description             string   `json:"Description,omitempty"`
	ProductReleaseOn        string   `json:"ProductReleaseOn,omitempty"`
	CompanyCode             string   `json:"CompanyCode,omitempty"`
	CompanyName             string   `json:"CompanyName,omitempty"`
	Creator                 string   `json:"Creator,omitempty"`
	CreatorName             string   `json:"CreatorName,omitempty"`
	CreatedDate             string   `json:"CreatedDate,omitempty"`
	CreditorNumber          string   `json:"CreditorNumber,omitempty"`
	DebtorNumber            string   `json:"DebtorNumber,omitempty"`
	EntryKey                string   `json:"EntryKey,omitempty"`
	ExpiryDate              edm.Date `json:"ExpiryDate,omitempty"`
	FileName                string   `json:"FileName,omitempty"`
	FrontPageIntroText      string   `json:"FrontPageIntroText,omitempty"`
	GroupID                 int      `json:"GroupID,omitempty"`
	GroupName               string   `json:"GroupName,omitempty"`
	GroupTemplate           string   `json:"GroupTemplate,omitempty"`
	Hid                     int      `json:"HID,omitempty"`
	ID                      string   `json:"ID,omitempty"`
	IsBinaryServiceEnabled  bool     `json:"IsBinaryServiceEnabled,omitempty"`
	IsMailMerge             bool     `json:"IsMailMerge,omitempty"`
	IsTemplate              bool     `json:"IsTemplate,omitempty"`
	ItemCode                string   `json:"ItemCode,omitempty"`
	ItemName                string   `json:"ItemName,omitempty"`
	ItemSerialNumberCode    string   `json:"ItemSerialNumberCode,omitempty"`
	ItemSerialNumber        string   `json:"ItemSerialNumber,omitempty"`
	ItemSerialNumberName    string   `json:"ItemSerialNumberName,omitempty"`
	LanguageID              string   `json:"LanguageID,omitempty"`
	LanguageName            string   `json:"LanguageName,omitempty"`
	MessageIDBody           string   `json:"MessageIDBody,omitempty"`
	MessageIDDocument       string   `json:"MessageIDDocument,omitempty"`
	Modifier                string   `json:"Modifier,omitempty"`
	ModifierName            string   `json:"ModifierName,omitempty"`
	ModifiedDate            edm.Date `json:"ModifiedDate,omitempty"`
	NeedsApproval           bool     `json:"NeedsApproval,omitempty"`
	NewsType                int      `json:"NewsType,omitempty"`
	Note                    string   `json:"Note,omitempty"`
	OrderNumber             string   `json:"OrderNumber,omitempty"`
	OurRef                  string   `json:"OurRef,omitempty"`
	Owner                   int      `json:"Owner,omitempty"`
	OwnerName               string   `json:"OwnerName,omitempty"`
	OwnerType               int      `json:"OwnerType,omitempty"`
	OwnerTypeValue          string   `json:"OwnerTypeValue,omitempty"`
	ParentID                string   `json:"ParentID,omitempty"`
	PaymentReference        string   `json:"PaymentReference,omitempty"`
	PersonID                int      `json:"PersonID,omitempty"`
	PersonEmail             string   `json:"PersonEmail,omitempty"`
	PersonName              string   `json:"PersonName,omitempty"`
	ProjectDescription      string   `json:"ProjectDescription,omitempty"`
	ProjectNumber           string   `json:"ProjectNumber,omitempty"`
	ProjectCompanyCode      string   `json:"ProjectCompanyCode,omitempty"`
	Security                bool     `json:"Security,omitempty"`
	SecurityLevel           int      `json:"SecurityLevel,omitempty"`
	ShipmentMethod          string   `json:"ShipmentMethod,omitempty"`
	Source                  string   `json:"Source,omitempty"`
	Status                  int      `json:"Status,omitempty"`
	StatusText              string   `json:"StatusText,omitempty"`
	SubCategory             string   `json:"SubCategory,omitempty"`
	Subject                 string   `json:"Subject,omitempty"`
	Type                    int      `json:"Type,omitempty"`
	TypeAccountLabel        string   `json:"TypeAccountLabel,omitempty"`
	TypeAccountSelection    string   `json:"TypeAccountSelection,omitempty"`
	TypeAccountType         string   `json:"TypeAccountType,omitempty"`
	TypeAllowAttachment     int      `json:"TypeAllowAttachment,omitempty"`
	TypeDaysToExpiry        float64  `json:"TypeDaysToExpiry,omitempty"`
	TypeDefSecurity         int      `json:"TypeDefSecurity,omitempty"`
	TypeDeleteRights        string   `json:"TypeDeleteRights,omitempty"`
	TypeEditRights          string   `json:"TypeEditRights,omitempty"`
	TypeEditSecurity        int      `json:"TypeEditSecurity,omitempty"`
	TypeFinanceSelection    string   `json:"TypeFinanceSelection,omitempty"`
	TypeItemLabel           string   `json:"TypeItemLabel,omitempty"`
	TypeItemSelection       string   `json:"TypeItemSelection,omitempty"`
	TypeItemStatus          string   `json:"TypeItemStatus,omitempty"`
	TypeName                string   `json:"TypeName,omitempty"`
	TypeProjectSelection    string   `json:"TypeProjectSelection,omitempty"`
	TypeResourceLabel       string   `json:"TypeResourceLabel,omitempty"`
	TypeResourceSelection   string   `json:"TypeResourceSelection,omitempty"`
	TypeResourceStatus      string   `json:"TypeResourceStatus,omitempty"`
	TypeResourceType        string   `json:"TypeResourceType,omitempty"`
	TypeTemplate            string   `json:"TypeTemplate,omitempty"`
	TypeUseAccount          int      `json:"TypeUseAccount,omitempty"`
	TypeUseApprove          int      `json:"TypeUseApprove,omitempty"`
	TypeUseAssortment       int      `json:"TypeUseAssortment,omitempty"`
	TypeUseFinance          int      `json:"TypeUseFinance,omitempty"`
	TypeUseItem             int      `json:"TypeUseItem,omitempty"`
	TypeUseItemSerialNumber int      `json:"TypeUseItemSerialNumber,omitempty"`
	TypeUseOrderNumber      int      `json:"TypeUseOrderNumber,omitempty"`
	TypeUseOurReference     int      `json:"TypeUseOurReference,omitempty"`
	TypeUseOwner            int      `json:"TypeUseOwner,omitempty"`
	TypeUsePaymentReference int      `json:"TypeUsePaymentReference,omitempty"`
	TypeUseProject          int      `json:"TypeUseProject,omitempty"`
	TypeUsePublish          int      `json:"TypeUsePublish,omitempty"`
	TypeUseRating           int      `json:"TypeUseRating,omitempty"`
	TypeUseResource         int      `json:"TypeUseResource,omitempty"`
	TypeUseShipmentMethod   int      `json:"TypeUseShipmentMethod,omitempty"`
	TypeUseWarehouse        int      `json:"TypeUseWarehouse,omitempty"`
	TypeUseYourReference    int      `json:"TypeUseYourReference,omitempty"`
	TypeVersionType         string   `json:"TypeVersionType,omitempty"`
	UseAssociate            int      `json:"UseAssociate,omitempty"`
	UseClass                int      `json:"UseClass,omitempty"`
	UseCompany              int      `json:"UseCompany,omitempty"`
	UseEntryKey             int      `json:"UseEntryKey,omitempty"`
	UseItem                 int      `json:"UseItem,omitempty"`
	UseItemSerialNumber     int      `json:"UseItemSerialNumber,omitempty"`
	UseLanguage             int      `json:"UseLanguage,omitempty"`
	UsePerson               int      `json:"UsePerson,omitempty"`
	UseProjectNumber        int      `json:"UseProjectNumber,omitempty"`
	Version                 string   `json:"Version,omitempty"`
	VersionActivate         bool     `json:"VersionActivate,omitempty"`
	VersionCheckedOut       edm.Date `json:"VersionCheckedOut,omitempty"`
	VersionCheckedOutBy     string   `json:"VersionCheckedOutBy,omitempty"`
	VersionCheckedOutByName string   `json:"VersionCheckedOutByName,omitempty"`
	VersionCreatedBy        string   `json:"VersionCreatedBy,omitempty"`
	VersionCreatedDate      string   `json:"VersionCreatedDate,omitempty"`
	VersionFilename         string   `json:"VersionFilename,omitempty"`
	VersionFPIntroText      string   `json:"VersionFPIntroText,omitempty"`
	VersionHID              float64  `json:"VersionHID,omitempty"`
	VersionID               string   `json:"VersionID,omitempty"`
	VersionModifiedBy       string   `json:"VersionModifiedBy,omitempty"`
	VersionModifiedDate     edm.Date `json:"VersionModifiedDate,omitempty"`
	VersionNote             string   `json:"VersionNote,omitempty"`
	VersionStartDate        edm.Date `json:"VersionStartDate,omitempty"`
	VersionStatus           float64  `json:"VersionStatus,omitempty"`
	VersionStatusText       string   `json:"VersionStatusText,omitempty"`
	VersionSubject          string   `json:"VersionSubject,omitempty"`
	VersionToDealWith       string   `json:"VersionToDealWith,omitempty"`
	ViewDate                string   `json:"ViewDate,omitempty"`
	Warehouse               string   `json:"Warehouse,omitempty"`
	YourRef                 string   `json:"YourRef,omitempty"`
}
