package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewCostCenterSubGroupReadRequest() CostCenterSubGroupReadRequest {
	return CostCenterSubGroupReadRequest{
		client:      c,
		queryParams: c.NewCostCenterSubGroupReadQueryParams(),
		pathParams:  c.NewCostCenterSubGroupReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCostCenterSubGroupReadRequestBody(),
	}
}

type CostCenterSubGroupReadRequest struct {
	client      *Client
	queryParams *CostCenterSubGroupReadQueryParams
	pathParams  *CostCenterSubGroupReadPathParams
	method      string
	headers     http.Header
	requestBody CostCenterSubGroupReadRequestBody
}

func (c *Client) NewCostCenterSubGroupReadQueryParams() *CostCenterSubGroupReadQueryParams {
	selectFields, _ := utils.Fields(&Project{})
	return &CostCenterSubGroupReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CostCenterSubGroupReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CostCenterSubGroupReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CostCenterSubGroupReadRequest) QueryParams() *CostCenterSubGroupReadQueryParams {
	return r.queryParams
}

func (c *Client) NewCostCenterSubGroupReadPathParams() *CostCenterSubGroupReadPathParams {
	return &CostCenterSubGroupReadPathParams{}
}

type CostCenterSubGroupReadPathParams struct {
}

func (p *CostCenterSubGroupReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CostCenterSubGroupReadRequest) PathParams() *CostCenterSubGroupReadPathParams {
	return r.pathParams
}

func (r *CostCenterSubGroupReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *CostCenterSubGroupReadRequest) Method() string {
	return r.method
}

func (s *Client) NewCostCenterSubGroupReadRequestBody() CostCenterSubGroupReadRequestBody {
	return CostCenterSubGroupReadRequestBody{}
}

type CostCenterSubGroupReadRequestBody CostCenterSubGroups

func (r *CostCenterSubGroupReadRequest) RequestBody() *CostCenterSubGroupReadRequestBody {
	return &r.requestBody
}

func (r *CostCenterSubGroupReadRequest) SetRequestBody(body CostCenterSubGroupReadRequestBody) {
	r.requestBody = body
}

func (r *CostCenterSubGroupReadRequest) NewResponseBody() *CostCenterSubGroupReadResponseBody {
	return &CostCenterSubGroupReadResponseBody{}
}

type CostCenterSubGroupReadResponseBody CostCenterSubGroups

func (r *CostCenterSubGroupReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("CostCenterSubGroup", r.PathParams())
}

func (r *CostCenterSubGroupReadRequest) Do() (CostCenterSubGroupReadResponseBody, error) {
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

type CostCenterSubGroups []CostCenterSubGroup

type CostCenterSubGroup struct {
}
