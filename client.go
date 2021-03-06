package api_utils

import (
	"net/http"
	"log"
	"bytes"
	"fmt"
	"io/ioutil"
)

const (
	AddedStatusCode = 201
)

type RequestParameters struct {
	Method    string
	Params    string
	Uri       string
	AuthToken string
}

func NewRequestParameters(method string, params string, uri string) *RequestParameters {
	return &RequestParameters{
		Method:method,
		Params:params,
		Uri:uri,
		AuthToken:"",
	}
}

type ClientRequest struct {
}

func NewClientRequest() *ClientRequest {
	return &ClientRequest{}
}

type RequestFactory interface {
	NewRequest(rp *RequestParameters) (*http.Request, error)
}

func (cr ClientRequest)NewRequest(rp *RequestParameters) (*http.Request, error) {
	var req *http.Request
	var err error

	method := rp.Method
	uri := rp.Uri
	token := rp.AuthToken

	switch method{
	case http.MethodGet:
		fallthrough
	case http.MethodDelete:
		uri := uri + "?" + rp.Params
		req, err = http.NewRequest(method, uri, nil)
	case http.MethodPost:
		fallthrough
	case http.MethodPatch:
		fallthrough
	case http.MethodPut:
		log.Printf("Body is %s", rp)
		req, err = http.NewRequest(method, uri, bytes.NewBuffer([]byte(rp.Params)))
		req.Header.Set("Content-Type", "application/json")
		if (token != "") {
			req.SetBasicAuth(token, "x-oauth-basic")
		}
	default:
		panic(fmt.Sprintf("invalid method consumed %s", method))
	}

	Errorchecker(err)

	return req, err
}

type Client struct {
	rf RequestFactory
	dc http.Client
}

type Callable interface {
	Call(rp *RequestParameters) (*http.Response, error)
}

func (c Client)Call(rp *RequestParameters) (*http.Response, error) {
	req, err := c.rf.NewRequest(rp)
	Errorchecker(err)

	resp, err := c.dc.Do(req)
	Errorchecker(err)

	processStatusCode(resp, rp)

	return resp, err
}

func processStatusCode(resp *http.Response, rp *RequestParameters) {
	if http.StatusOK <= resp.StatusCode && resp.StatusCode >= AddedStatusCode {
		body, err := ioutil.ReadAll(resp.Body);
		log.Printf("Call %s failed with status code %d", rp.Method, resp.StatusCode)
		log.Printf("The reasons are: %v", string(body))

		if err != nil {
			resp.Body.Close()
			panic(err)
		}
	}
}

func NewClient() *Client {
	return &Client{rf:NewClientRequest(), dc: http.Client{}}
}
