package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewDocumentCreateRequest() DocumentCreateRequest {
	return DocumentCreateRequest{
		client:      c,
		queryParams: c.NewDocumentCreateQueryParams(),
		pathParams:  c.NewDocumentCreatePathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: DocumentCreateRequestBody{},
	}
}

type DocumentCreateRequest struct {
	client      *Client
	queryParams *DocumentCreateQueryParams
	pathParams  *DocumentCreatePathParams
	method      string
	headers     http.Header
	requestBody DocumentCreateRequestBody
}

func (c *Client) NewDocumentCreateQueryParams() *DocumentCreateQueryParams {
	return &DocumentCreateQueryParams{}
}

type DocumentCreateQueryParams struct{}

func (p DocumentCreateQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DocumentCreateRequest) QueryParams() *DocumentCreateQueryParams {
	return r.queryParams
}

func (c *Client) NewDocumentCreatePathParams() *DocumentCreatePathParams {
	return &DocumentCreatePathParams{}
}

type DocumentCreatePathParams struct{}

func (p *DocumentCreatePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DocumentCreateRequest) PathParams() *DocumentCreatePathParams {
	return r.pathParams
}

func (r *DocumentCreateRequest) SetMethod(method string) {
	r.method = method
}

func (r *DocumentCreateRequest) Method() string {
	return r.method
}

func (r *DocumentCreateRequest) NewRequestBody() DocumentCreateRequestBody {
	return DocumentCreateRequestBody{}
}

type DocumentCreateRequestBody NewDocument

func (r *DocumentCreateRequest) RequestBody() *DocumentCreateRequestBody {
	return &r.requestBody
}

func (r *DocumentCreateRequest) SetRequestBody(body DocumentCreateRequestBody) {
	r.requestBody = body
}

func (r *DocumentCreateRequest) NewResponseBody() *DocumentCreateResponseBody {
	return &DocumentCreateResponseBody{}
}

type DocumentCreateResponseBody struct {
	D struct {
		Next     string       `json:"__next"`
		MetaData edm.MetaData `json:"__metadata"`

		NewDocument
	} `json:"d"`
}

func (r *DocumentCreateRequest) URL() url.URL {
	return r.client.GetEndpointURL("Document", r.PathParams())
}

func (r *DocumentCreateRequest) Do() (DocumentCreateResponseBody, error) {
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

type NewDocument Document


