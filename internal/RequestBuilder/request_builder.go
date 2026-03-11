package request_builder

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestBuilder struct {
	Request http.Request
	Client  *http.Client
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		Client: http.DefaultClient,
		Request: http.Request{
			Method: GET,
			URL: &url.URL{
				Scheme: HTTPS,
			},
		},
	}
}

// ===HEADERS===
func (this *RequestBuilder) Headers() http.Header {
	return this.Request.Header
}

// ===CLIENT===
func (this *RequestBuilder) SetDefaultClient() error {
	return this.SetClient(http.DefaultClient)
}

func (this *RequestBuilder) SetClient(client *http.Client) error {
	this.Client = client
	return nil
}

// ===BODY===
func (this *RequestBuilder) SetBody(body *string) {
	if body == nil {
		this.Request.Body = io.NopCloser(strings.NewReader(""))
		return
	}
	this.Request.Body = io.NopCloser(strings.NewReader(*body))
}

// ===DESTINATION===
func (this *RequestBuilder) SetSchema(schema string) error {

	this.Request.URL.Scheme = HTTPS
	return nil
}

func (this *RequestBuilder) SetHostnameAndPort(hostname string) {
	this.Request.URL.Host = hostname
}

func (this *RequestBuilder) SetURLPath(path string) {
	this.Request.URL.Path = path
}

// ===PROCESS===

func (this *RequestBuilder) Do() (*http.Response, error) {
	return this.Client.Do(&this.Request)
}
