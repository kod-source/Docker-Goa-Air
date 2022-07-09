// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "test_build": test Resource Client
//
// Command:
// $ goagen
// --design=github.com/kod-source/docker-goa-mysql/design
// --out=/Users/horikoudai/program-practice/docker-goa/app
// --version=v1.5.13

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// AddTestPath computes a request path to the add action of test.
func AddTestPath(left int, right int) string {
	param0 := strconv.Itoa(left)
	param1 := strconv.Itoa(right)
	return fmt.Sprintf("/add/%s/%s", param0, param1)
}

// add returns the sum of the left and right parameters in the response body
func (c *Client) AddTest(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAddTestRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddTestRequest create the request corresponding to the add action endpoint of the test resource.
func (c *Client) NewAddTestRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
