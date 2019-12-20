package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewCostCenterGroupReadRequest() CostCenterGroupReadRequest {
	return CostCenterGroupReadRequest{
		client:      c,
		queryParams: c.NewCostCenterGroupReadQueryParams(),
		pathParams:  c.NewCostCenterGroupReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCostCenterGroupReadRequestBody(),
	}
}

type CostCenterGroupReadRequest struct {
	client      *Client
	queryParams *CostCenterGroupReadQueryParams
	pathParams  *CostCenterGroupReadPathParams
	method      string
	headers     http.Header
	requestBody CostCenterGroupReadRequestBody
}

func (c *Client) NewCostCenterGroupReadQueryParams() *CostCenterGroupReadQueryParams {
	selectFields, _ := utils.Fields(&Project{})
	return &CostCenterGroupReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CostCenterGroupReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CostCenterGroupReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CostCenterGroupReadRequest) QueryParams() *CostCenterGroupReadQueryParams {
	return r.queryParams
}

func (c *Client) NewCostCenterGroupReadPathParams() *CostCenterGroupReadPathParams {
	return &CostCenterGroupReadPathParams{}
}

type CostCenterGroupReadPathParams struct {
}

func (p *CostCenterGroupReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CostCenterGroupReadRequest) PathParams() *CostCenterGroupReadPathParams {
	return r.pathParams
}

func (r *CostCenterGroupReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *CostCenterGroupReadRequest) Method() string {
	return r.method
}

func (s *Client) NewCostCenterGroupReadRequestBody() CostCenterGroupReadRequestBody {
	return CostCenterGroupReadRequestBody{}
}

type CostCenterGroupReadRequestBody CostCenterGroups

func (r *CostCenterGroupReadRequest) RequestBody() *CostCenterGroupReadRequestBody {
	return &r.requestBody
}

func (r *CostCenterGroupReadRequest) SetRequestBody(body CostCenterGroupReadRequestBody) {
	r.requestBody = body
}

func (r *CostCenterGroupReadRequest) NewResponseBody() *CostCenterGroupReadResponseBody {
	return &CostCenterGroupReadResponseBody{}
}

type CostCenterGroupReadResponseBody CostCenterGroups

func (r *CostCenterGroupReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("CostCenterGroup", r.PathParams())
}

func (r *CostCenterGroupReadRequest) Do() (CostCenterGroupReadResponseBody, error) {
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

type CostCenterGroups []CostCenterGroup

type CostCenterGroup struct {
	ClassID     int
	Description string
}
