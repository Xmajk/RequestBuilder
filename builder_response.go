package requestBuilder

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

// BuilderResponse is returned by [RequestBuilder.Do] and holds the HTTP status
// code, the original [net/http.Response], and the response body pre-loaded as a
// plain string.
type BuilderResponse struct {
	// StatusCode is the HTTP status code of the response (e.g. 200, 404).
	StatusCode int
	// Response is the underlying [net/http.Response]. Note that Body has already
	// been read and closed; do not attempt to read it again.
	Response http.Response
	// Body contains the full response body as a string. It is nil only when the
	// response itself was nil.
	Body *string
}

// newBuilderResponse creates a BuilderResponse from an *http.Response, reading
// the body into memory. Returns nil if response is nil or reading fails.
func newBuilderResponse(response *http.Response) *BuilderResponse {
	if response == nil {
		return nil
	}

	statusCode := response.StatusCode
	new := &BuilderResponse{
		StatusCode: statusCode,
		Response:   *response,
		Body:       nil,
	}
	err := new.loadBody()
	if err != nil {
		return nil
	}

	return new
}

// loadBody reads the full response body into this.Body using a fixed-size
// buffer and closes the body when done.
func (this *BuilderResponse) loadBody() error {
	body := ""

	defer this.Response.Body.Close()

	for {
		tmp := make([]byte, RESPONSE_LOADER_BUFFER_SIZE)

		i, err := this.Response.Body.Read(tmp)
		if err == io.EOF {
			if i > 0 {
				body += string(tmp[:i])
			}

			break
		} else if err != nil {
			return err
		}

		body += string(tmp[:i])
	}

	this.Body = &body

	return nil
}

// DecodeJSON parses the response body as JSON and stores the result in the
// value pointed to by v. It is a convenience wrapper around [encoding/json.Unmarshal].
//
//	var result MyStruct
//	if err := res.DecodeJSON(&result); err != nil {
//	    log.Fatal(err)
//	}
func (this *BuilderResponse) DecodeJSON(v any) error {
	return json.Unmarshal([]byte(*this.Body), v)
}

// DecodeXML parses the response body as XML and stores the result in the
// value pointed to by v. It is a convenience wrapper around [encoding/xml.Unmarshal].
//
//	var result MyStruct
//	if err := res.DecodeXML(&result); err != nil {
//	    log.Fatal(err)
//	}
func (this *BuilderResponse) DecodeXML(v any) error {
	return xml.Unmarshal([]byte(*this.Body), v)
}
