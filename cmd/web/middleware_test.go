package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myh myHandler
	h := NoSurf(&myh)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("Type is not http handler but is %t ", v))
	}
}
func TestSessionLoad(t *testing.T) {
	var myh myHandler
	h := SessionLoad(&myh)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("Type is not http handler but is %t ", v))
	}
}
