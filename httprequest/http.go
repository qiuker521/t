package httprequest
import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetData(url1 string, timeout int) ([]byte, error) {
	return DataWithHeader("GET", url1, timeout, map[string]string{}, map[string]string{})
}

func DataWithHeader(
	method string,
	url1 string,
	timeout int,
	header map[string]string,
	params map[string]string,
) ([]byte, error) {

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	paramsA := ""
	if method == "POST" {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v)
		}
		paramsA = values.Encode()
	}

	req, _ := http.NewRequest(method, url1, strings.NewReader(paramsA))

	for k, v := range header {
		req.Header.Set(k, v)
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if method == "GET" {
		values := req.URL.Query()
		for k, v := range params {
			values.Add(k, v)
		}
		req.URL.RawQuery = values.Encode()
	}

	getResp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer getResp.Body.Close()
	body, err := ioutil.ReadAll(getResp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

