// +build e2e

package test

import (
	"github.com/go-resty/v2"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Running E2E test for health check endoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fail()
	}

	fmt.Println(resp.StatusCode())
}
