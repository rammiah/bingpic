package rttp

import (
	"bytes"
	"net/http"
	"strings"
)

func (c *Client) PostWithHeader(url string, header map[string]string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	// set header
	for k, v := range header {
		req.Header.Set(k, v)
	}
	return c.Do(req)
}

func (c *Client) PostFormData(url string, header map[string]string, form interface{}) {
	//req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	//if err != nil {
	//	return nil, err
	//}
	// set header
	//for k, v := range header {
	//	req.Header.Set(k, v)
	//}
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(""))


}
