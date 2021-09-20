# go-proxychain
Http client with connection through proxy chain

## Usage example

```go
/*
Proxy chain: 
127.0.0.1:1080 => 212.119.47.150:1080 => 212.119.47.229:1085 => api.ipify.org
*/
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
```
