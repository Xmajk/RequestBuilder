package request_builder

// HTTP method constants for use with [RequestBuilder].
const (
	POST   = "POST"
	GET    = "GET"
	DELETE = "DELETE"
	PUT    = "PUT"
	HEAD   = "HEAD"
)

// URL scheme constants.
const (
	HTTP  = "http"
	HTTPS = "https"
)

// errors
const (
	// NOT_IMPLEMENTED_ERROR is returned as a panic message for features that
	// have not been implemented yet.
	NOT_IMPLEMENTED_ERROR = "ERROR: Not implemented yet!"
)

// RESPONSE_LOADER_BUFFER_SIZE is the size in bytes of the read buffer used
// when loading the HTTP response body into memory.
const (
	RESPONSE_LOADER_BUFFER_SIZE = 1_000_000
)
