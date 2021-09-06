package proxychain

import (
	"context"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
)

type ProxyConfig struct {
	address string
	auth    *proxy.Auth
}

func GetClient(chain []ProxyConfig) *http.Client {
	client := http.DefaultClient
	client.Transport = getTransport(chain)
	return client
}

func getTransport(chain []ProxyConfig) *http.Transport {
	dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
		return getProxyDialer(chain).Dial(network, address)
	}
	tr := &http.Transport{
		DialContext: dialContext,
	}
	return tr
}

// dialer1 -> dialer2 -> dialer... -> url
func getProxyDialer(chain []ProxyConfig) (dialer proxy.Dialer) {
	firstProxy := chain[0]
	dialer, _ = proxy.SOCKS5("tcp", firstProxy.address, firstProxy.auth, proxy.Direct)
	for i := 1; i < len(chain); i++ {
		chainProxy := chain[i]
		dialer, _ = proxy.SOCKS5("tcp", chainProxy.address, chainProxy.auth, dialer)
	}
	return
}
