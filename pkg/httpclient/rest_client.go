package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
	urlUtils "net/url"
	"time"
)

type Config struct {
	Host            string
	Params          map[string]string
	TimeoutInMillis int // Timeout in millisecond
	Headers         map[string]string
}

type RESTClient struct {
	host    string
	timeout time.Duration
	headers map[string]string
}

func NewRESTClient(config Config) RESTClient {
	timeout := time.Millisecond * 3000
	if config.TimeoutInMillis != 0 {
		timeout = time.Millisecond * time.Duration(config.TimeoutInMillis)
	}

	return RESTClient{
		host:    config.Host,
		timeout: timeout,
		headers: config.Headers,
	}
}

func (rcl RESTClient) Get(url string, params map[string]string, dest interface{}) (err error) {
	if err = rcl.sendRequest(http.MethodGet, url, params, nil, dest); err != nil {
		return
	}

	return
}

func (rcl RESTClient) sendRequest(method, endpoint string, params map[string]string, body io.Reader, dest interface{}) (err error) {
	var client = &http.Client{
		Timeout: rcl.timeout,
	}

	url := rcl.host + endpoint

	if len(params) > 0 {
		qs := urlUtils.Values{}

		for parKey, parValue := range params {
			qs.Add(parKey, parValue)
		}

		url += qs.Encode()
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(dest); err != nil {
		return
	}

	return nil
}
