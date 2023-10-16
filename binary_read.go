package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewBinaryReadRequest() BinaryReadRequest {
	return BinaryReadRequest{
		client:      c,
		queryParams: c.NewBinaryReadQueryParams(),
		pathParams:  c.NewBinaryReadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewBinaryReadRequestBody(),
	}
}

type BinaryReadRequest struct {
	client      *Client
	queryParams *BinaryReadQueryParams
	pathParams  *BinaryReadPathParams
	method      string
	headers     http.Header
	requestBody BinaryReadRequestBody
}

func (c *Client) NewBinaryReadQueryParams() *BinaryReadQueryParams {
	selectFields, _ := utils.Fields(&Project{})
	return &BinaryReadQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type BinaryReadQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p BinaryReadQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *BinaryReadRequest) QueryParams() *BinaryReadQueryParams {
	return r.queryParams
}

func (c *Client) NewBinaryReadPathParams() *BinaryReadPathParams {
	return &BinaryReadPathParams{}
}

type BinaryReadPathParams struct {
}

func (p *BinaryReadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BinaryReadRequest) PathParams() *BinaryReadPathParams {
	return r.pathParams
}

func (r *BinaryReadRequest) SetMethod(method string) {
	r.method = method
}

func (r *BinaryReadRequest) Method() string {
	return r.method
}

func (s *Client) NewBinaryReadRequestBody() BinaryReadRequestBody {
	return BinaryReadRequestBody{}
}

type BinaryReadRequestBody struct{}

func (r *BinaryReadRequest) RequestBody() *BinaryReadRequestBody {
	return &r.requestBody
}

func (r *BinaryReadRequest) SetRequestBody(body BinaryReadRequestBody) {
	r.requestBody = body
}

func (r *BinaryReadRequest) NewResponseBody() *BinaryReadResponseBody {
	return &BinaryReadResponseBody{}
}

type BinaryReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func (r *BinaryReadRequest) URL() url.URL {
	return r.client.GetEndpointURL("Binary", r.PathParams())
}

func (r *BinaryReadRequest) Do() (BinaryReadResponseBody, error) {
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
