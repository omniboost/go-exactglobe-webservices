package webservices

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-exactglobe-webservices/edm"
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

type CostUnitReadResponseBody struct {
	D struct {
		Results []struct {
			Next     string       `json:"__next"`
			MetaData edm.MetaData `json:"__metadata"`

			CostUnit
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

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
	Code             string `json:"Code"`
	CompanyCode      string `json:"CompanyCode"`
	CompanyName      string `json:"CompanyName"`
	CreatedDate      string `json:"CreatedDate"`
	Creator          int    `json:"Creator"`
	CreatorFullName  string `json:"CreatorFullName"`
	Description      string `json:"Description"`
	Description1     string `json:"Description1,omitempty"`
	Description2     string `json:"Description2,omitempty"`
	Description3     string `json:"Description3,omitempty"`
	Description4     string `json:"Description4,omitempty"`
	ID               int    `json:"ID"`
	ModifiedDate     string `json:"ModifiedDate"`
	Modifier         int    `json:"Modifier"`
	ModifierFullName string `json:"ModifierFullName"`
	NumberFreeField1 int    `json:"NumberFreeField1,omitempty"`
	NumberFreeField2 int    `json:"NumberFreeField2,omitempty"`
	NumberFreeField3 int    `json:"NumberFreeField3,omitempty"`
	NumberFreeField4 int    `json:"NumberFreeField4,omitempty"`
	NumberFreeField5 int    `json:"NumberFreeField5,omitempty"`
	TextFreeField1   string `json:"TextFreeField1,omitempty"`
	TextFreeField2   string `json:"TextFreeField2,omitempty"`
	TextFreeField3   string `json:"TextFreeField3,omitempty"`
	TextFreeField4   string `json:"TextFreeField4,omitempty"`
	TextFreeField5   string `json:"TextFreeField5,omitempty"`
}
