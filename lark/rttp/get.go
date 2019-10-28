package rttp

import "net/http"

func (c *Client) GetWithHeader(url string, header map[string]string) (*http.Response, error) {
	// GET没有body
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// 设置一下header
	for k, v := range header {
		req.Header.Add(k, v)
	}
	return c.Do(req)
}
