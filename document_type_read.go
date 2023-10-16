package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewDocumentTypeReadRequest() DocumentTypeReadRequest {
	return DocumentTypeReadRequest{
		client:      c,
		queryParams: c.NewDocumentTypeReadQueryParams(),
		pathParams:  c.NewDocumentTypeReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewDocumentTypeReadRequestBody(),
	}
}

type DocumentTypeReadRequest struct {
	client      *Client
	queryParams *DocumentTypeReadQueryParams
	pathParams  *DocumentTypeReadPathParams
	method      string
	headers     http.Header
	requestBody DocumentTypeReadRequestBody
}

func (c *Client) NewDocumentTypeReadQueryParams() *DocumentTypeReadQueryParams {
	selectFields, _ := utils.Fields(&Project{})
	return &DocumentTypeReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type DocumentTypeReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p DocumentTypeReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DocumentTypeReadRequest) QueryParams() *DocumentTypeReadQueryParams {
	return r.queryParams
}

func (c *Client) NewDocumentTypeReadPathParams() *DocumentTypeReadPathParams {
	return &DocumentTypeReadPathParams{}
}

type DocumentTypeReadPathParams struct {
}

func (p *DocumentTypeReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DocumentTypeReadRequest) PathParams() *DocumentTypeReadPathParams {
	return r.pathParams
}

func (r *DocumentTypeReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *DocumentTypeReadRequest) Method() string {
	return r.method
}

func (s *Client) NewDocumentTypeReadRequestBody() DocumentTypeReadRequestBody {
	return DocumentTypeReadRequestBody{}
}

type DocumentTypeReadRequestBody struct{}

func (r *DocumentTypeReadRequest) RequestBody() *DocumentTypeReadRequestBody {
	return &r.requestBody
}

func (r *DocumentTypeReadRequest) SetRequestBody(body DocumentTypeReadRequestBody) {
	r.requestBody = body
}

func (r *DocumentTypeReadRequest) NewResponseBody() *DocumentTypeReadResponseBody {
	return &DocumentTypeReadResponseBody{}
}

type DocumentTypeReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			DocumentType
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *DocumentTypeReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("DocumentType", r.PathParams())
}

func (r *DocumentTypeReadRequest) Do() (DocumentTypeReadResponseBody, error) {
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

type DocumentTypes []DocumentType

type DocumentType struct {
	edm.MetaData                 `json:"__metadata,omitempty"`
	AccountLabel                 string      `json:"AccountLabel,omitempty"`
	AccountSelection             string      `json:"AccountSelection,omitempty"`
	AccountType                  string      `json:"AccountType,omitempty"`
	UseAccount                   int         `json:"UseAccount,omitempty"`
	Active                       bool        `json:"Active,omitempty"`
	AllowApproval                bool        `json:"AllowApproval,omitempty"`
	AssortmentDefault            int         `json:"AssortmentDefault,omitempty"`
	UseAssortment                int         `json:"UseAssortment,omitempty"`
	AllowAttachment              bool        `json:"AllowAttachment,omitempty"`
	CreatorRoleLevel1            int         `json:"CreatorRoleLevel1,omitempty"`
	CreatorRoleLevel2            int         `json:"CreatorRoleLevel2,omitempty"`
	CreatorRoleLevel3            int         `json:"CreatorRoleLevel3,omitempty"`
	CreatorRole1                 int         `json:"CreatorRole1,omitempty"`
	CreatorRoleLevel1Description string      `json:"CreatorRoleLevel1Description,omitempty"`
	CreatorRole2                 int         `json:"CreatorRole2,omitempty"`
	CreatorRoleLevel2Description string      `json:"CreatorRoleLevel2Description,omitempty"`
	CreatorRole3                 int         `json:"CreatorRole3,omitempty"`
	CreatorRoleLevel3Description string      `json:"CreatorRoleLevel3Description,omitempty"`
	CreatorSecurityLevel         int         `json:"CreatorSecurityLevel,omitempty"`
	IsSystemGenerated            bool        `json:"IsSystemGenerated,omitempty"`
	Creator                      int         `json:"Creator,omitempty"`
	CreatedDate                  string      `json:"CreatedDate,omitempty"`
	DaysBeforeDeletion           int         `json:"DaysBeforeDeletion,omitempty"`
	DefaultCategory              string      `json:"DefaultCategory,omitempty"`
	DefaultGroupID               int         `json:"DefaultGroupID,omitempty"`
	DefaultSecurityLevel         int         `json:"DefaultSecurityLevel,omitempty"`
	DefaultSubCategory           string      `json:"DefaultSubCategory,omitempty"`
	DeleteRights                 string      `json:"DeleteRights,omitempty"`
	TypeDescription              string      `json:"TypeDescription,omitempty"`
	EditRights                   string      `json:"EditRights,omitempty"`
	AllowEditSecurityLevel       bool        `json:"AllowEditSecurityLevel,omitempty"`
	FinancialOurReference        int         `json:"FinancialOurReference,omitempty"`
	FinancialSelection           string      `json:"FinancialSelection,omitempty"`
	UseTransaction               int         `json:"UseTransaction,omitempty"`
	FinancialYourReference       int         `json:"FinancialYourReference,omitempty"`
	GroupID                      string      `json:"GroupID,omitempty"`
	GroupName                    string      `json:"GroupName,omitempty"`
	ID                           int         `json:"ID,omitempty"`
	ItemDefault                  string      `json:"ItemDefault,omitempty"`
	ItemLabel                    string      `json:"ItemLabel,omitempty"`
	ItemDescription              string      `json:"ItemDescription,omitempty"`
	ItemSelection                string      `json:"ItemSelection,omitempty"`
	ItemStatus                   string      `json:"ItemStatus,omitempty"`
	UseItem                      int         `json:"UseItem,omitempty"`
	UseSerialNumber              int         `json:"UseSerialNumber,omitempty"`
	Manager                      int         `json:"Manager,omitempty"`
	ManagerFullName              string      `json:"ManagerFullName,omitempty"`
	Modifier                     int         `json:"Modifier,omitempty"`
	ModifiedDate                 string      `json:"ModifiedDate,omitempty"`
	UseOrderNumber               int         `json:"UseOrderNumber,omitempty"`
	UsePaymentReference          int         `json:"UsePaymentReference,omitempty"`
	Policy                       string      `json:"Policy,omitempty"`
	PolicyDescription            string `json:"PolicyDescription,omitempty"`
	ProjectDefault               string      `json:"ProjectDefault,omitempty"`
	ProjectDescription           string      `json:"ProjectDescription,omitempty"`
	ProjectSelection             string      `json:"ProjectSelection,omitempty"`
	UseProject                   int         `json:"UseProject,omitempty"`
	ResourceLabel                string      `json:"ResourceLabel,omitempty"`
	ResourceSelection            string      `json:"ResourceSelection,omitempty"`
	ResourceStatus               string      `json:"ResourceStatus,omitempty"`
	ResourceType                 string      `json:"ResourceType,omitempty"`
	UseResource                  int         `json:"UseResource,omitempty"`
	UseSecurityApproval          bool        `json:"UseSecurityApproval,omitempty"`
	UseShipmentMethod            int         `json:"UseShipmentMethod,omitempty"`
	SystemType                   bool        `json:"SystemType,omitempty"`
	Layout                       string      `json:"Layout,omitempty"`
	LayoutSubject                string      `json:"LayoutSubject,omitempty"`
	UseOwner                     bool        `json:"UseOwner,omitempty"`
	AllowPublish                 bool        `json:"AllowPublish,omitempty"`
	AllowRating                  bool        `json:"AllowRating,omitempty"`
	UsedInStatisticReports       bool        `json:"UsedInStatisticReports,omitempty"`
	VersionType                  string      `json:"VersionType,omitempty"`
	UseWarehouse                 int         `json:"UseWarehouse,omitempty"`
}
