package proxychain

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetClient(t *testing.T) {
	proxyChain := []ProxyConfig{
		{
			address: "localhost:1080",
			auth:    nil,
		},
		{
			address: "212.119.47.150:1080",
			auth: &proxy.Auth{
				User:     "myLogin",
				Password: "myPassword",
			},
		},
		{
			address: "212.119.47.229:1085",
			auth:    nil,
		},
	}
	req, _ := http.NewRequest(http.MethodGet, "https://api.ipify.org/", nil)
	client := GetClient(proxyChain)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "212.119.47.229", string(body))
}
