package request_builder_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	request_builder "github.com/Xmajk/RequestBuilder"
)

// ExampleNewRequestBuilder demonstrates creating a builder and performing a
// simple GET request.
func ExampleNewRequestBuilder() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	}))
	defer srv.Close()

	rb := request_builder.NewRequestBuilder()
	rb.SetHostnameAndPort(srv.Listener.Addr().String())
	rb.SetSchema(request_builder.HTTP)
	rb.Request.URL.Scheme = "http"

	res, err := rb.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(*res.Body)
	// Output:
	// hello
}

// ExampleRequestBuilder_Headers shows how to add custom headers before sending
// the request.
func ExampleRequestBuilder_Headers() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Header.Get("X-Custom"))
	}))
	defer srv.Close()

	rb := request_builder.NewRequestBuilder()
	rb.Request.URL.Scheme = "http"
	rb.SetHostnameAndPort(srv.Listener.Addr().String())
	rb.Headers().Set("X-Custom", "my-value")

	res, err := rb.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(*res.Body)
	// Output:
	// my-value
}

// ExampleBuildResponse_DecodeJSON demonstrates unmarshalling a JSON response
// body into a Go struct.
func ExampleBuilderResponse_DecodeJSON() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"id":1,"name":"Alice"}`)
	}))
	defer srv.Close()

	rb := request_builder.NewRequestBuilder()
	rb.Request.URL.Scheme = "http"
	rb.SetHostnameAndPort(srv.Listener.Addr().String())

	res, err := rb.Do()
	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	var u User
	if err := res.DecodeJSON(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID=%d Name=%s\n", u.ID, u.Name)
	// Output:
	// ID=1 Name=Alice
}

// ExampleBuildResponse_DecodeXML demonstrates unmarshalling an XML response
// body into a Go struct.
func ExampleBuilderResponse_DecodeXML() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, `<user><id>2</id><name>Bob</name></user>`)
	}))
	defer srv.Close()

	rb := request_builder.NewRequestBuilder()
	rb.Request.URL.Scheme = "http"
	rb.SetHostnameAndPort(srv.Listener.Addr().String())

	res, err := rb.Do()
	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	}

	var u User
	if err := res.DecodeXML(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID=%d Name=%s\n", u.ID, u.Name)
	// Output:
	// ID=2 Name=Bob
}

// ExampleRequestBuilder_SetBody shows how to attach a request body for a POST
// request.
func ExampleRequestBuilder_SetBody() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 512)
		n, _ := r.Body.Read(buf)
		fmt.Fprintf(w, "received: %s", buf[:n])
	}))
	defer srv.Close()

	rb := request_builder.NewRequestBuilder()
	rb.Request.URL.Scheme = "http"
	rb.Request.Method = request_builder.POST
	rb.SetHostnameAndPort(srv.Listener.Addr().String())

	payload := `{"action":"ping"}`
	rb.SetBody(&payload)

	res, err := rb.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(*res.Body)
	// Output:
	// received: {"action":"ping"}
}
