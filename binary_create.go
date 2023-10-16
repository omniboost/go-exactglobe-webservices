package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewBinaryCreateRequest() BinaryCreateRequest {
	return BinaryCreateRequest{
		client:      c,
		queryParams: c.NewBinaryCreateQueryParams(),
		pathParams:  c.NewBinaryCreatePathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: BinaryCreateRequestBody{},
	}
}

type BinaryCreateRequest struct {
	client      *Client
	queryParams *BinaryCreateQueryParams
	pathParams  *BinaryCreatePathParams
	method      string
	headers     http.Header
	requestBody BinaryCreateRequestBody
}

func (c *Client) NewBinaryCreateQueryParams() *BinaryCreateQueryParams {
	return &BinaryCreateQueryParams{}
}

type BinaryCreateQueryParams struct{}

func (p BinaryCreateQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *BinaryCreateRequest) QueryParams() *BinaryCreateQueryParams {
	return r.queryParams
}

func (c *Client) NewBinaryCreatePathParams() *BinaryCreatePathParams {
	return &BinaryCreatePathParams{}
}

type BinaryCreatePathParams struct{}

func (p *BinaryCreatePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BinaryCreateRequest) PathParams() *BinaryCreatePathParams {
	return r.pathParams
}

func (r *BinaryCreateRequest) SetMethod(method string) {
	r.method = method
}

func (r *BinaryCreateRequest) Method() string {
	return r.method
}

func (r *BinaryCreateRequest) NewRequestBody() BinaryCreateRequestBody {
	return BinaryCreateRequestBody{}
}

type BinaryCreateRequestBody NewBinary

func (r *BinaryCreateRequest) RequestBody() *BinaryCreateRequestBody {
	return &r.requestBody
}

func (r *BinaryCreateRequest) SetRequestBody(body BinaryCreateRequestBody) {
	r.requestBody = body
}

func (r *BinaryCreateRequest) NewResponseBody() *BinaryCreateResponseBody {
	return &BinaryCreateResponseBody{}
}

type BinaryCreateResponseBody struct {
	D struct {
		Next     string       `json:"__next"`
		MetaData edm.MetaData `json:"__metadata"`

		NewBinary
	} `json:"d"`
}

func (r *BinaryCreateRequest) URL() url.URL {
	return r.client.GetEndpointURL("Binary", r.PathParams())
}

func (r *BinaryCreateRequest) Do() (BinaryCreateResponseBody, error) {
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

type NewBinary Binary

type Binary struct {
	Data       string `json:"Data,omitempty"`
	DataString string `json:"DataString,omitempty"`
	Encoded    bool   `json:"Encoded,omitempty"`
	MessageID  string `json:"MessageID,omitempty"`
	Sequence   int    `json:"Sequence,omitempty"`
}
