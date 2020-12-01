package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewCostUnitReadRequest() CostUnitReadRequest {
	return CostUnitReadRequest{
		client:      c,
		queryParams: c.NewCostUnitReadQueryParams(),
		pathParams:  c.NewCostUnitReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCostUnitReadRequestBody(),
	}
}

type CostUnitReadRequest struct {
	client      *Client
	queryParams *CostUnitReadQueryParams
	pathParams  *CostUnitReadPathParams
	method      string
	headers     http.Header
	requestBody CostUnitReadRequestBody
}

func (c *Client) NewCostUnitReadQueryParams() *CostUnitReadQueryParams {
	selectFields, _ := utils.Fields(&CostUnit{})
	return &CostUnitReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type CostUnitReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CostUnitReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CostUnitReadRequest) QueryParams() *CostUnitReadQueryParams {
	return r.queryParams
}

func (c *Client) NewCostUnitReadPathParams() *CostUnitReadPathParams {
	return &CostUnitReadPathParams{}
}

type CostUnitReadPathParams struct {
}

func (p *CostUnitReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CostUnitReadRequest) PathParams() *CostUnitReadPathParams {
	return r.pathParams
}

func (r *CostUnitReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *CostUnitReadRequest) Method() string {
	return r.method
}

func (s *Client) NewCostUnitReadRequestBody() CostUnitReadRequestBody {
	return CostUnitReadRequestBody{}
}

type CostUnitReadRequestBody CostUnits

func (r *CostUnitReadRequest) RequestBody() *CostUnitReadRequestBody {
	return &r.requestBody
}

func (r *CostUnitReadRequest) SetRequestBody(body CostUnitReadRequestBody) {
	r.requestBody = body
}

func (r *CostUnitReadRequest) NewResponseBody() *CostUnitReadResponseBody {
	return &CostUnitReadResponseBody{}
}

type CostUnitReadResponseBody CostUnits

func (r *CostUnitReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("CostUnit", r.PathParams())
}

func (r *CostUnitReadRequest) Do() (CostUnitReadResponseBody, error) {
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

type CostUnits []CostUnit

type CostUnit struct {
}
