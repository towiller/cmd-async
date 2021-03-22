package connpool

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type HttpMsg struct {
	Headers map[string]string
	Body    string
	Uri     string
	Args    string
	Method  string
}

type HttpPool struct {
	HttpClient *http.Client
}

func NewHttpPool(opts ...HttpOption) *HttpPool {
	options := &HttpOptions{
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     5,
		DialTimeout:         5,
		KeepAlive:           10,
	}

	for _, o := range opts {
		o(options)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(options.DialTimeout) * time.Second,
				KeepAlive: time.Duration(options.KeepAlive) * time.Second,
			}).DialContext,
			MaxIdleConns:        options.MaxIdleConns,
			MaxIdleConnsPerHost: options.MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(options.IdleConnTimeout) * time.Second,
		},
	}

	return &HttpPool{
		HttpClient: httpClient,
	}
}

func (h *HttpPool) Request(method string, header map[string][]string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header = header
	if err != nil {
		log.Println("http request error:", err)
		return nil, err
	}

	return h.HttpClient.Do(req)
}
