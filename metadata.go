package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewMetadataRequest() MetadataRequest {
	return MetadataRequest{
		client:      c,
		queryParams: c.NewMetadataQueryParams(),
		pathParams:  c.NewMetadataPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewMetadataRequestBody(),
	}
}

type MetadataRequest struct {
	client      *Client
	queryParams *MetadataQueryParams
	pathParams  *MetadataPathParams
	method      string
	headers     http.Header
	requestBody MetadataRequestBody
}

func (c *Client) NewMetadataQueryParams() *MetadataQueryParams {
	selectFields := []string{}
	return &MetadataQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type MetadataQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p MetadataQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *MetadataRequest) QueryParams() *MetadataQueryParams {
	return r.queryParams
}

func (c *Client) NewMetadataPathParams() *MetadataPathParams {
	return &MetadataPathParams{}
}

type MetadataPathParams struct {
}

func (p *MetadataPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MetadataRequest) PathParams() *MetadataPathParams {
	return r.pathParams
}

func (r *MetadataRequest) SetMethod(method string) {
	r.method = method
}

func (r *MetadataRequest) Method() string {
	return r.method
}

func (s *Client) NewMetadataRequestBody() MetadataRequestBody {
	return MetadataRequestBody{}
}

type MetadataRequestBody struct{}

func (r *MetadataRequest) RequestBody() *MetadataRequestBody {
	return &r.requestBody
}

func (r *MetadataRequest) SetRequestBody(body MetadataRequestBody) {
	r.requestBody = body
}

func (r *MetadataRequest) NewResponseBody() *MetadataResponseBody {
	return &MetadataResponseBody{}
}

type MetadataResponseBody struct{}

func (r *MetadataRequest) URL() url.URL {
	return r.client.GetEndpointURL("$metadata", r.PathParams())
}

func (r *MetadataRequest) Do() (MetadataResponseBody, error) {
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

type Metadata struct {
}
