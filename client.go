package api_utils

import (
	"net/http"
	"log"
	"bytes"
	"io/ioutil"
	"fmt"
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

	if err != nil {
		panic(err)
	}

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
	req, _ := c.rf.NewRequest(rp)
	resp, err := c.dc.Do(req)

	log.Printf("%v %v %v", int(resp.StatusCode) != http.StatusOK || int(resp.StatusCode) != 201, resp.StatusCode, http.StatusOK)

	if http.StatusOK <= resp.StatusCode && resp.StatusCode >= AddedStatusCode {
		body, _ := ioutil.ReadAll(resp.Body);
		resp.Body.Close()

		log.Printf("Call %s failed with status code %d", rp.Method, resp.StatusCode)
		log.Printf("The reasons are: %v", string(body))
	}

	if err != nil {
		panic(err)
	}

	return resp, err
}

func NewClient() *Client {
	return &Client{rf:NewClientRequest(), dc: http.Client{}}
}
