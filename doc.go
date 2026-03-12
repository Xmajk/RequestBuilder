// Package request_builder provides a fluent HTTP request builder for
// constructing and executing HTTP requests with automatic response body loading.
//
// # Overview
//
// The central type is [RequestBuilder], which wraps [net/http.Request] and
// [net/http.Client] and exposes a method-chaining-friendly API for configuring
// the method, URL, headers, and body before executing the request with [RequestBuilder.Do].
//
// The response is returned as a [BuilderResponse], which reads the entire body
// into memory and exposes it as a plain string pointer. For structured payloads
// the response can be decoded directly with [BuilderResponse.DecodeJSON] or
// [BuilderResponse.DecodeXML].
//
// # Basic usage
//
//	rb := request_builder.NewRequestBuilder()
//	rb.SetHostnameAndPort("api.example.com")
//	rb.SetURLPath("/v1/users")
//	rb.Headers().Set("Authorization", "Bearer my-token")
//
//	res, err := rb.Do()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(*res.Body)
//
// # Decoding a JSON response
//
//	type User struct {
//	    ID   int    `json:"id"`
//	    Name string `json:"name"`
//	}
//
//	var u User
//	if err := res.DecodeJSON(&u); err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(u.Name)
package request_builder
