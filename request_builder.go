package request_builder

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RequestBuilder wraps an [net/http.Request] and its associated [net/http.Client],
// providing a convenient API for building and executing HTTP requests.
//
// Use [NewRequestBuilder] to obtain a properly initialised instance.
type RequestBuilder struct {
	// Request is the underlying HTTP request being built.
	Request http.Request
	// Client is the HTTP client used to execute the request.
	Client *http.Client
}

// NewRequestBuilder returns a new [RequestBuilder] with sensible defaults:
// the HTTP method is set to GET, the scheme to HTTPS, and the client to
// [net/http.DefaultClient].
func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		Client: http.DefaultClient,
		Request: http.Request{
			Method: GET,
			URL: &url.URL{
				Scheme: HTTPS,
			},
			Header: make(http.Header),
		},
	}
}

// ===HEADERS===

// Headers returns the [net/http.Header] map of the underlying request,
// allowing callers to add, set, or delete headers directly.
//
//	rb.Headers().Set("Accept", "application/json")
//	rb.Headers().Add("X-Request-ID", "abc123")
func (this *RequestBuilder) Headers() http.Header {
	return this.Request.Header
}

// ===CLIENT===

// SetDefaultClient resets the HTTP client to [net/http.DefaultClient].
func (this *RequestBuilder) SetDefaultClient() error {
	return this.SetClient(http.DefaultClient)
}

// SetClient replaces the HTTP client used to execute requests.
// Pass a custom [net/http.Client] to control timeouts, redirects, or transport.
func (this *RequestBuilder) SetClient(client *http.Client) error {
	this.Client = client
	return nil
}

// ===BODY===

// SetBody sets the request body to the content of the provided string.
// Passing nil results in an empty body.
func (this *RequestBuilder) SetBody(body *string) {
	if body == nil {
		this.Request.Body = io.NopCloser(strings.NewReader(""))
		return
	}
	this.Request.Body = io.NopCloser(strings.NewReader(*body))
}

// ===DESTINATION===

// SetSchema sets the URL scheme of the request.
// Currently always sets the scheme to HTTPS regardless of the provided value.
func (this *RequestBuilder) SetSchema(schema string) error {
	this.Request.URL.Scheme = HTTPS
	return nil
}

// SetHostnameAndPort sets the host (and optional port) of the request URL.
// The value is passed directly to [net/url.URL.Host], e.g. "example.com" or
// "example.com:8080".
func (this *RequestBuilder) SetHostnameAndPort(hostname string) {
	this.Request.URL.Host = hostname
}

// SetURLPath sets the path component of the request URL, e.g. "/v1/users".
func (this *RequestBuilder) SetURLPath(path string) {
	this.Request.URL.Path = path
}

// ===PROCESS===

// Do executes the HTTP request and returns a [BuilderResponse] whose body has
// been fully read into memory.
//
// An error is returned if the request fails or if the response body cannot be
// read.
func (this *RequestBuilder) Do() (*BuilderResponse, error) {
	res, err := this.Client.Do(&this.Request)
	if err != nil {
		return nil, err
	}

	builderReponse := newBuilderResponse(res)
	if builderReponse == nil {
		return nil, errors.New("ERROR: create response failed")
	}

	return builderReponse, nil
}
