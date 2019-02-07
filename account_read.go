package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
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
	return &AccountReadQueryParams{}
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

type AccountReadRequestBody struct {
}

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
