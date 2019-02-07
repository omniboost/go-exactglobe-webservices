package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
	uuid "github.com/satori/go.uuid"
)

func (c *Client) NewProjectReadRequest() ProjectReadRequest {
	return ProjectReadRequest{
		client:      c,
		queryParams: c.NewProjectReadQueryParams(),
		pathParams:  c.NewProjectReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewProjectReadRequestBody(),
	}
}

type ProjectReadRequest struct {
	client      *Client
	queryParams *ProjectReadQueryParams
	pathParams  *ProjectReadPathParams
	method      string
	headers     http.Header
	requestBody ProjectReadRequestBody
}

func (c *Client) NewProjectReadQueryParams() *ProjectReadQueryParams {
	selectFields, _ := utils.Fields(&Project{})
	return &ProjectReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type ProjectReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p ProjectReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ProjectReadRequest) QueryParams() *ProjectReadQueryParams {
	return r.queryParams
}

func (c *Client) NewProjectReadPathParams() *ProjectReadPathParams {
	return &ProjectReadPathParams{}
}

type ProjectReadPathParams struct {
}

func (p *ProjectReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ProjectReadRequest) PathParams() *ProjectReadPathParams {
	return r.pathParams
}

func (r *ProjectReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *ProjectReadRequest) Method() string {
	return r.method
}

func (s *Client) NewProjectReadRequestBody() ProjectReadRequestBody {
	return ProjectReadRequestBody{}
}

type ProjectReadRequestBody struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}

func (r *ProjectReadRequest) RequestBody() *ProjectReadRequestBody {
	return &r.requestBody
}

func (r *ProjectReadRequest) SetRequestBody(body ProjectReadRequestBody) {
	r.requestBody = body
}

func (r *ProjectReadRequest) NewResponseBody() *ProjectReadResponseBody {
	return &ProjectReadResponseBody{}
}

type ProjectReadResponseBody Projects

func (r *ProjectReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("Project", r.PathParams())
}

func (r *ProjectReadRequest) Do() (ProjectReadResponseBody, error) {
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

type Projects []Project

type Project struct {
	BackFlushing           bool
	CompressionSO          string
	ProjectCostCenter      string
	ProjectCostCenterName  string
	ProjectCostUnit        string
	ProjectCostUnitName    string
	Creator                int
	CreatorName            string
	CreatedDate            *DateTime
	CustomerCode           string
	CustomerCountry        string
	CustomerID             *uuid.UUID
	CustomerName           string
	CustomerStatus         string
	CustomerType           string
	DateFreeField1         *DateTime
	DateFreeField2         *DateTime
	DateFreeField3         *DateTime
	DateFreeField4         *DateTime
	DateFreeField5         *DateTime
	Description            string
	CompanyCode            string
	CompanyName            string
	Member                 bool
	MemberCompanyCode      string
	MemberCompanyName      string
	MemberName             string
	MemberJobTitle         string
	ImageFileName          string
	EndDate                *DateTime
	StartDate              *DateTime
	IsBinaryServiceEnabled bool
	ItemCode               string
	ItemName               string
	JobGroup               int
	JobGroupName           string
	LinkHour               bool
	Material               string
	Memo                   string
	MessageID              *uuid.UUID
	Modifie                int
	ModifierName           string
	ModifiedDate           DateTime
	NumberFreeField1       int
	NumberFreeField2       int
	NumberFreeField3       int
	NumberFreeField4       int
	NumberFreeField5       int
	ParentProject          string
	ParentProjectName      string
	PercentCompletion      int
	Assortment             int
	AssortmentName         string
	ProjectNumber          string
	ProjectManager         int
	SecurityLevel          int
	SerialNumber           int
	SerialNumberName       string
	Status                 string
	TextFreeField1         string
	TextFreeField2         string
	TextFreeField3         string
	TextFreeField4         string
	TextFreeField5         string
	TrainingCosts          int
	TrainingQuantity       int
	Type                   string
	VisibleMember          bool
	Warehouse              string
	WarehouseName          string
	WIPMethod              string
	WorkInProgressCost     string
	YesNoFreeField1        bool
	YesNoFreeField2        bool
	YesNoFreeField3        bool
	YesNoFreeField4        bool
	YesNoFreeField5        bool
}
