package rttp

import (
	"net/http"
	"time"
)

type Client struct {
	*http.Client
}

// 系统的http
func NewClient() *Client {
	return NewClientWithTimeout(5 * time.Second)
}

func NewClientWithTimeout(timeout time.Duration) *Client {
	return &Client{&http.Client{
		Timeout: timeout,
	}}
}
