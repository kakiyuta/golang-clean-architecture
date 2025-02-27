// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// ProductVariants defines model for ProductVariants.
type ProductVariants struct {
	Id          *int64     `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Validations *[]Variant `json:"validations,omitempty"`
}

// Prouct defines model for Prouct.
type Prouct struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Variant defines model for Variant.
type Variant struct {
	Id    *int64  `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Price *int    `json:"price,omitempty"`
}

// PostV1LoginJSONBody defines parameters for PostV1Login.
type PostV1LoginJSONBody struct {
	Email    openapi_types.Email `json:"email"`
	Password string              `json:"password"`
}

// GetV1ProductsParams defines parameters for GetV1Products.
type GetV1ProductsParams struct {
	// Limit 取得する商品数
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset 取得する商品のオフセット
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// PostV1ProductsJSONBody defines parameters for PostV1Products.
type PostV1ProductsJSONBody struct {
	Name string `json:"name"`
}

// PostV1LoginJSONRequestBody defines body for PostV1Login for application/json ContentType.
type PostV1LoginJSONRequestBody PostV1LoginJSONBody

// PostV1ProductsJSONRequestBody defines body for PostV1Products for application/json ContentType.
type PostV1ProductsJSONRequestBody PostV1ProductsJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// PostV1LoginWithBody request with any body
	PostV1LoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostV1Login(ctx context.Context, body PostV1LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetV1Products request
	GetV1Products(ctx context.Context, params *GetV1ProductsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostV1ProductsWithBody request with any body
	PostV1ProductsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostV1Products(ctx context.Context, body PostV1ProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostV1LoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostV1LoginRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostV1Login(ctx context.Context, body PostV1LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostV1LoginRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetV1Products(ctx context.Context, params *GetV1ProductsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetV1ProductsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostV1ProductsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostV1ProductsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostV1Products(ctx context.Context, body PostV1ProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostV1ProductsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostV1LoginRequest calls the generic PostV1Login builder with application/json body
func NewPostV1LoginRequest(server string, body PostV1LoginJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostV1LoginRequestWithBody(server, "application/json", bodyReader)
}

// NewPostV1LoginRequestWithBody generates requests for PostV1Login with any type of body
func NewPostV1LoginRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/login")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetV1ProductsRequest generates requests for GetV1Products
func NewGetV1ProductsRequest(server string, params *GetV1ProductsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/products")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Limit != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Offset != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "offset", runtime.ParamLocationQuery, *params.Offset); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostV1ProductsRequest calls the generic PostV1Products builder with application/json body
func NewPostV1ProductsRequest(server string, body PostV1ProductsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostV1ProductsRequestWithBody(server, "application/json", bodyReader)
}

// NewPostV1ProductsRequestWithBody generates requests for PostV1Products with any type of body
func NewPostV1ProductsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/products")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// PostV1LoginWithBodyWithResponse request with any body
	PostV1LoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostV1LoginResponse, error)

	PostV1LoginWithResponse(ctx context.Context, body PostV1LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*PostV1LoginResponse, error)

	// GetV1ProductsWithResponse request
	GetV1ProductsWithResponse(ctx context.Context, params *GetV1ProductsParams, reqEditors ...RequestEditorFn) (*GetV1ProductsResponse, error)

	// PostV1ProductsWithBodyWithResponse request with any body
	PostV1ProductsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostV1ProductsResponse, error)

	PostV1ProductsWithResponse(ctx context.Context, body PostV1ProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostV1ProductsResponse, error)
}

type PostV1LoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Token *string `json:"token,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r PostV1LoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostV1LoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetV1ProductsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Products *[]ProductVariants `json:"products,omitempty"`
		Total    *int64             `json:"total,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetV1ProductsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetV1ProductsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostV1ProductsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *Prouct
}

// Status returns HTTPResponse.Status
func (r PostV1ProductsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostV1ProductsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostV1LoginWithBodyWithResponse request with arbitrary body returning *PostV1LoginResponse
func (c *ClientWithResponses) PostV1LoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostV1LoginResponse, error) {
	rsp, err := c.PostV1LoginWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostV1LoginResponse(rsp)
}

func (c *ClientWithResponses) PostV1LoginWithResponse(ctx context.Context, body PostV1LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*PostV1LoginResponse, error) {
	rsp, err := c.PostV1Login(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostV1LoginResponse(rsp)
}

// GetV1ProductsWithResponse request returning *GetV1ProductsResponse
func (c *ClientWithResponses) GetV1ProductsWithResponse(ctx context.Context, params *GetV1ProductsParams, reqEditors ...RequestEditorFn) (*GetV1ProductsResponse, error) {
	rsp, err := c.GetV1Products(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetV1ProductsResponse(rsp)
}

// PostV1ProductsWithBodyWithResponse request with arbitrary body returning *PostV1ProductsResponse
func (c *ClientWithResponses) PostV1ProductsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostV1ProductsResponse, error) {
	rsp, err := c.PostV1ProductsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostV1ProductsResponse(rsp)
}

func (c *ClientWithResponses) PostV1ProductsWithResponse(ctx context.Context, body PostV1ProductsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostV1ProductsResponse, error) {
	rsp, err := c.PostV1Products(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostV1ProductsResponse(rsp)
}

// ParsePostV1LoginResponse parses an HTTP response from a PostV1LoginWithResponse call
func ParsePostV1LoginResponse(rsp *http.Response) (*PostV1LoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostV1LoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Token *string `json:"token,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetV1ProductsResponse parses an HTTP response from a GetV1ProductsWithResponse call
func ParseGetV1ProductsResponse(rsp *http.Response) (*GetV1ProductsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetV1ProductsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Products *[]ProductVariants `json:"products,omitempty"`
			Total    *int64             `json:"total,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostV1ProductsResponse parses an HTTP response from a PostV1ProductsWithResponse call
func ParsePostV1ProductsResponse(rsp *http.Response) (*PostV1ProductsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostV1ProductsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest Prouct
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /v1/login)
	PostV1Login(ctx echo.Context) error

	// (GET /v1/products)
	GetV1Products(ctx echo.Context, params GetV1ProductsParams) error

	// (POST /v1/products)
	PostV1Products(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostV1Login converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostV1Login(ctx)
	return err
}

// GetV1Products converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1Products(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetV1ProductsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetV1Products(ctx, params)
	return err
}

// PostV1Products converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1Products(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostV1Products(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/v1/login", wrapper.PostV1Login)
	router.GET(baseURL+"/v1/products", wrapper.GetV1Products)
	router.POST(baseURL+"/v1/products", wrapper.PostV1Products)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RVQW/jRBT+K9WDo5XEYbtofaIbLSuHlgYK2exWPUzsl3hSe2Z2Ztytu8qhiUBIICGB",
	"Cto7B4QEWsGxEn/GCoh/gWbsxglx1QLdizXzZt7nb9775puXEPBEcIZMK/BeggoiTIgd9iQP00D3iaSk",
	"XBWSC5Saop3R0HxHXCZEgweU6fv3wAGdCSymOEYJUwcYSdBsLVeUlpSNzcIJiWlINOWsANSY2MHbEkfg",
	"wVvNilyzZNYs+Zj0Eo9ISTKYVgE+nGBgd/QkTwN999Tr/nVF7A3USUgarK4sczZ5TB1QGKSS6uzAVKyg",
	"MEQiUe6kOlo22SQV4YpLpLWAqcGgbMTN1hBVIKkwLQIPHnW2DiIutnZ6vkmiOsYqCg6coFTFTrfRarQM",
	"dS6QEUHBg3dsyAFBdGRJNU/cZszHlNmScaU3f5jPf85nr/PZD/n8N7Bg0srFD8GDHle67+5aBAckPk9R",
	"6Yc8zAxOwJnGohtEiJgGNq85UZxVOt/sFSaExnZwShJhjxfxMb5XThsBT8Cpellsd2o6RpR6wWW4DrWM",
	"OjWCMgegEkPwDpe4y4SjzUavpWiZog0owZkqztJutf5HJTQ/RrZOH7NuNHwc0H3a9T89890Pqa989vF2",
	"0PHv+8di0O90HzQw67pBu58N2vGxP+Gnu51uPHzyvhiaTUl0MnzS33k2iKLh4KF6drA9GbZbdLfTFU8H",
	"H9H9ySN375On2d7Z8Yu9SfdB493uvbPaYtUUY105+x+Y6NSxKhOFldlzjbFGZ4vPflxcfL749jyffbP4",
	"+rvF799vqO0x6r7buwIyrZEkQY1SgXe4gWcx8vNX+ezLAviPi9dgrhV48DxFmcHVdYeYJlSDs9KIEEck",
	"jTV4bsuBhJzSJE3MxMwoK2dOjRncTCM//yWf/ZTPL/LZZT6f5/MvrmHFRyOF19Ba5dGq4XF0p0Jcbd6t",
	"Xoh/PlwbL4UDmmsS38qVb621ynetHlYd9/BoemQ8odbilrr789XlX1/9eo3LrQjvbozu6q2pbveBHWyV",
	"fwKrvF1kY3OC9vb2TZZlAf+bTbn/6hw3tD6t71FHItEYWlOYTv8OAAD//0fMucL3CAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
