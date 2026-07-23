package main

import (
	"fmt"
	"net/http"
)

type GzipResponseWriter struct {
	http.ResponseWriter
	index int
}

type GzipResponseWriterWithCloseNotify struct {
	*GzipResponseWriter
}

func (w GzipResponseWriterWithCloseNotify) CloseNotify() <-chan bool {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	fmt.Println("hello")
}
