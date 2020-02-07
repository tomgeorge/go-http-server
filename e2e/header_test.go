package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestReturnsHeaders(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/headers")
	if err != nil {
		t.Errorf("Expected no error but was %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	respBody := string(body)
	if respBody != "asdf" {
		t.Errorf("Expected asdf but was %v", resp)
	}
}
