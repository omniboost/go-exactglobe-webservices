package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewAccountUpdateRequest() AccountUpdateRequest {
	return AccountUpdateRequest{
		client:      c,
		queryParams: c.NewAccountUpdateQueryParams(),
		pathParams:  c.NewAccountUpdatePathParams(),
		method:      http.MethodPut,
		headers:     http.Header{},
		requestBody: AccountUpdateRequestBody{},
	}
}

type AccountUpdateRequest struct {
	client      *Client
	queryParams *AccountUpdateQueryParams
	pathParams  *AccountUpdatePathParams
	method      string
	headers     http.Header
	requestBody AccountUpdateRequestBody
}

func (c *Client) NewAccountUpdateQueryParams() *AccountUpdateQueryParams {
	return &AccountUpdateQueryParams{}
}

type AccountUpdateQueryParams struct {
}

func (p AccountUpdateQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountUpdateRequest) QueryParams() *AccountUpdateQueryParams {
	return r.queryParams
}

func (c *Client) NewAccountUpdatePathParams() *AccountUpdatePathParams {
	return &AccountUpdatePathParams{}
}

type AccountUpdatePathParams struct {
	GUID string
}

func (p *AccountUpdatePathParams) Params() map[string]string {
	return map[string]string{
		"GUID": p.GUID,
	}
}

func (r *AccountUpdateRequest) PathParams() *AccountUpdatePathParams {
	return r.pathParams
}

func (r *AccountUpdateRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountUpdateRequest) Method() string {
	return r.method
}

func (r *AccountUpdateRequest) NewRequestBody() AccountUpdateRequestBody {
	return AccountUpdateRequestBody{}
}

type AccountUpdateRequestBody NewAccount

func (r *AccountUpdateRequest) RequestBody() *AccountUpdateRequestBody {
	return &r.requestBody
}

func (r *AccountUpdateRequest) SetRequestBody(body AccountUpdateRequestBody) {
	r.requestBody = body
}

func (r *AccountUpdateRequest) NewResponseBody() *AccountUpdateResponseBody {
	return &AccountUpdateResponseBody{}
}

type AccountUpdateResponseBody struct {
	D struct {
		Next     string       `json:"__next"`
		MetaData edm.MetaData `json:"__metadata"`

		Account
	} `json:"d"`
}

func (r *AccountUpdateRequest) URL() url.URL {
	return r.client.GetEndpointURL("Account(guid'{{.GUID}}')", r.PathParams())
}

func (r *AccountUpdateRequest) Do() (AccountUpdateResponseBody, error) {
	// Update http request
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
