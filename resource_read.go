package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-globe-webservices/odata"
	"github.com/omniboost/go-globe-webservices/utils"
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
	}
}

type ResourceReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
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

type ResourceReadRequestBody struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
}

func (r *ResourceReadRequest) RequestBody() *ResourceReadRequestBody {
	return &r.requestBody
}

func (r *ResourceReadRequest) SetRequestBody(body ResourceReadRequestBody) {
	r.requestBody = body
}

func (r *ResourceReadRequest) NewResponseBody() *ResourceReadResponseBody {
	return &ResourceReadResponseBody{}
}

type ResourceReadResponseBody Resources

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

type Resources []Resource

type Resource struct {
	CompanyName string
	FirstName   string
	LastName    string
	FullName    string
}
