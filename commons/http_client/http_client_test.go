package http_client

import (
	"encoding/json"
	"testing"
	"time"
)

func TestHttpClientGetRequest(t *testing.T) {
	client := NewHttpClient(true, ConfigOption(Config{
		ClientTimeout: 10 * time.Second, // default: 5s
		RetryCount:    2,                // default: 3
	}))

	req := client.NewRequest()
	resp, err := req.Get("https://www.boredapi.com/api/activity")
	if err != nil {
		t.Error("error making http get request")
	}

	result := make(map[string]interface{})
	body := resp.Body()
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Error("error parsing result from http response")
	}
}
