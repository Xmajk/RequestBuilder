package main

import (
	"fmt"

	request_builder "github.com/Xmajk/RequestBuilder/internal/RequestBuilder"
)

func main() {
	builder := request_builder.NewRequestBuilder()

	builder.SetSchema(request_builder.HTTPS)
	builder.SetHostnameAndPort("google.com")

	res, err := builder.Do()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(res.StatusCode)
}
