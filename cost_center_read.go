package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewCostCenterReadRequest() CostCenterReadRequest {
	return CostCenterReadRequest{
		client:      c,
		queryParams: c.NewCostCenterReadQueryParams(),
		pathParams:  c.NewCostCenterReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCostCenterReadRequestBody(),
	}
}

type CostCenterReadRequest struct {
	client      *Client
	queryParams *CostCenterReadQueryParams
	pathParams  *CostCenterReadPathParams
	method      string
	headers     http.Header
	requestBody CostCenterReadRequestBody
}

func (c *Client) NewCostCenterReadQueryParams() *CostCenterReadQueryParams {
	selectFields, _ := utils.Fields(&CostCenter{})
	return &CostCenterReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CostCenterReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CostCenterReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CostCenterReadRequest) QueryParams() *CostCenterReadQueryParams {
	return r.queryParams
}

func (c *Client) NewCostCenterReadPathParams() *CostCenterReadPathParams {
	return &CostCenterReadPathParams{}
}

type CostCenterReadPathParams struct {
}

func (p *CostCenterReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CostCenterReadRequest) PathParams() *CostCenterReadPathParams {
	return r.pathParams
}

func (r *CostCenterReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *CostCenterReadRequest) Method() string {
	return r.method
}

func (s *Client) NewCostCenterReadRequestBody() CostCenterReadRequestBody {
	return CostCenterReadRequestBody{}
}

type CostCenterReadRequestBody CostCenters

func (r *CostCenterReadRequest) RequestBody() *CostCenterReadRequestBody {
	return &r.requestBody
}

func (r *CostCenterReadRequest) SetRequestBody(body CostCenterReadRequestBody) {
	r.requestBody = body
}

func (r *CostCenterReadRequest) NewResponseBody() *CostCenterReadResponseBody {
	return &CostCenterReadResponseBody{}
}

type CostCenterReadResponseBody CostCenters

func (r *CostCenterReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("CostCenter", r.PathParams())
}

func (r *CostCenterReadRequest) Do() (CostCenterReadResponseBody, error) {
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

type CostCenters []CostCenter

type CostCenter struct {
	AllocationLevel   int
	Class1            string
	Class2            string
	Class             string
	Class4            string
	ClassDescription1 string
	ClassDescription2 string
	ClassDescription3 string
	ClassDescription4 string
	Code              string
	CompanyCode       string
	CompanyName       string
	CreatedDate       DateTime
	Creator           int
	CreatorName       string
	Description       string
	Description1      string
	Description2      string
	Description3      string
	Description4      string
	DirectManager     int
	DirectManagerName string
	Enabled           bool
	GLAccount         string
	GLOffsetAccount   string
	ID                int
	ModifiedDate      DateTime
	Modifier          int
	ModifierName      string
	StandardRate      int
	TextFreeField1    string
	TextFreeField2    string
	TextFreeField3    string
	TextFreeField4    string
	TextFreeField5    string
	NumberFreeField1  int
	NumberFreeField2  int
	NumberFreeField3  int
	NumberFreeField4  int
	NumberFreeField5  int
}
