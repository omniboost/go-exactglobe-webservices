package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
	"github.com/omniboost/go-exactglobe-webservices/odata"
	"github.com/omniboost/go-exactglobe-webservices/utils"
)

func (c *Client) NewFinancialHeaderCreateRequest() FinancialHeaderCreateRequest {
	return FinancialHeaderCreateRequest{
		client:      c,
		queryParams: c.NewFinancialHeaderCreateQueryParams(),
		pathParams:  c.NewFinancialHeaderCreatePathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: FinancialHeaderCreateRequestBody{},
	}
}

type FinancialHeaderCreateRequest struct {
	client      *Client
	queryParams *FinancialHeaderCreateQueryParams
	pathParams  *FinancialHeaderCreatePathParams
	method      string
	headers     http.Header
	requestBody FinancialHeaderCreateRequestBody
}

func (c *Client) NewFinancialHeaderCreateQueryParams() *FinancialHeaderCreateQueryParams {
	selectFields, _ := utils.Fields(&Project{})
	return &FinancialHeaderCreateQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type FinancialHeaderCreateQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p FinancialHeaderCreateQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FinancialHeaderCreateRequest) QueryParams() *FinancialHeaderCreateQueryParams {
	return r.queryParams
}

func (c *Client) NewFinancialHeaderCreatePathParams() *FinancialHeaderCreatePathParams {
	return &FinancialHeaderCreatePathParams{}
}

type FinancialHeaderCreatePathParams struct {
}

func (p *FinancialHeaderCreatePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *FinancialHeaderCreateRequest) PathParams() *FinancialHeaderCreatePathParams {
	return r.pathParams
}

func (r *FinancialHeaderCreateRequest) SetMethod(method string) {
	r.method = method
}

func (r *FinancialHeaderCreateRequest) Method() string {
	return r.method
}

func (r *FinancialHeaderCreateRequest) NewRequestBody() FinancialHeaderCreateRequestBody {
	return FinancialHeaderCreateRequestBody{}
}

type FinancialHeaderCreateRequestBody NewFinancialHeader

func (r *FinancialHeaderCreateRequest) RequestBody() *FinancialHeaderCreateRequestBody {
	return &r.requestBody
}

func (r *FinancialHeaderCreateRequest) SetRequestBody(body FinancialHeaderCreateRequestBody) {
	r.requestBody = body
}

func (r *FinancialHeaderCreateRequest) NewResponseBody() *FinancialHeaderCreateResponseBody {
	return &FinancialHeaderCreateResponseBody{}
}

type FinancialHeaderCreateResponseBody struct {
	D struct {
		Next     string       `json:"__next"`
		MetaData edm.MetaData `json:"__metadata"`
	} `json:"d"`
}

func (r *FinancialHeaderCreateRequest) URL() url.URL {
	return r.client.GetEndpointURL("FinancialHeader", r.PathParams())
}

func (r *FinancialHeaderCreateRequest) Do() (FinancialHeaderCreateResponseBody, error) {
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

type NewFinancialHeader struct {
	TransactionKey string `json:"TransactionKey"`
	Journal        string `json:"Journal"`
	DebtorNumber   string `json:"DebtorNumber"`
	OurReference   string `json:"OurReference"`
	YourReference  string `json:"YourReference"`
	Description    string `json:"Description"`
}
