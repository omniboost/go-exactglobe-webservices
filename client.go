package webservices

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"text/template"

	ntlmssp "github.com/Azure/go-ntlmssp"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-exactglobe-webservices/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "dijk177740vm001:8020",
		Path:   "Services/Exact.Entity.REST.EG/",
	}
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client, baseURL url.URL, databaseName string, databaseServerName string, username string, password string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetDatabaseName(databaseName)
	client.SetDatabaseServerName(databaseServerName)
	client.SetUsername(username)
	client.SetPassword(password)
	client.SetBaseURL(baseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	databaseName       string
	databaseServerName string
	username           string
	password           string

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	// set NTLM authentication
	client.Transport = ntlmssp.Negotiator{
		RoundTripper: http.DefaultTransport,
	}

	c.http = client
}

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) DatabaseName() string {
	return c.databaseName
}

func (c *Client) SetDatabaseName(databaseName string) {
	c.databaseName = databaseName
}

func (c *Client) DatabaseServerName() string {
	return c.databaseServerName
}

func (c *Client) SetDatabaseServerName(databaseServerName string) {
	c.databaseServerName = databaseServerName
}

func (c *Client) Username() string {
	return c.username
}

func (c *Client) SetUsername(username string) {
	c.username = username
}

func (c *Client) Password() string {
	return c.password
}

func (c *Client) SetPassword(password string) {
	c.password = password
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) GetEndpointURL(path string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	clientURL.Path = clientURL.Path + path

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set credentials
	req.SetBasicAuth(c.Username(), c.Password())

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	req.Header.Add("DatabaseName", c.databaseName)
	req.Header.Add("ServerName", c.databaseServerName)
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())

	return req, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if responseBody == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	type Wrapper struct {
		D struct {
			Results json.RawMessage `json:"results"`
			Next    string          `json:"__next"`
		} `json:"d"`
	}

	// try to decode body into interface parameter
	w := &Wrapper{}
	dec := json.NewDecoder(httpResp.Body)
	if c.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err = dec.Decode(w)
	if err != nil && err != io.EOF {
		// create a simple error response
		errorResponse := &ErrorResponse{Response: httpResp}
		errorResponse.Errors = append(errorResponse.Errors, err)
		return httpResp, errorResponse
	}

	err = json.Unmarshal(w.D.Results, responseBody)
	if err != nil && err != io.EOF {
		// @TODO: fix this
		log.Fatal(err)
	}

	return httpResp, nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		errorResponse.Errors = append(errorResponse.Errors, errors.New(r.Status))
		return errorResponse
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		errorResponse.Errors = append(errorResponse.Errors, err)
		return errorResponse
	}

	return errorResponse
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Errors []error
}

type Error struct {
	Code       string     `json:"code"`
	Message    Message    `json:"message"`
	InnerError InnerError `json:"innererror"`
}

func (e Error) Empty() bool {
	return e.Code == "" && e.Message.Value == "" && e.InnerError.Message == ""
}

func (e Error) Error() string {
	if e.Code != "" {
		if e.InnerError.Message != "" {
			return fmt.Sprintf("%s: %s", e.Code, e.InnerError.Message)
		}
		return fmt.Sprintf("%s: %s", e.Code, e.Message.Value)
	}

	if e.InnerError.Message != "" {
		return fmt.Sprintf("%s", e.InnerError.Message)
	}

	return fmt.Sprintf("%s", e.Message.Value)
}

type Message struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

type InnerError struct {
	Message    string `json:"message"`
	Type       string `json:"type"`
	Stacktrace string `json:"stacktrace"`
}

func (r *ErrorResponse) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Error Error `json:"error"`
	}{}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	if !tmp.Error.Empty() {
		r.Errors = append(r.Errors, tmp.Error)
	}

	return nil
}

func (r ErrorResponse) Error() string {
	if len(r.Errors) > 0 {
		str := []string{}
		for _, err := range r.Errors {
			str = append(str, err.Error())
		}
		return strings.Join(str, ", ")
	}

	return r.Errors[0].Error()
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type PathParams interface {
	Params() map[string]string
}
