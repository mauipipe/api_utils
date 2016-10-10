package api

import (
	"net/http"
	"log"
	"bytes"
	"io/ioutil"
)

const (
	GET = "GET"
	POST = "POST"
	DELETE = "DELETE"
	PUT = "PUT"
	PATCH = "PATCH"
	OAuthBasicToken = "22c853d518616e71914c0e7e66be61c482eb9d82"
)

type ClientRequest struct {
}

func NewClientRequest() *ClientRequest {
	return &ClientRequest{}
}

type RequestFactory interface {
	NewRequest(method string, requestParams string, uri string) (*http.Request, error)
}

func (cr ClientRequest)NewRequest(method string, rp string, uri string) (*http.Request, error) {
	var req *http.Request
	var err error

	switch method{
	case GET:
		fallthrough
	case DELETE:
		uri := uri + "?" + rp
		req, err = http.NewRequest(method, uri, nil)
	case POST:
		fallthrough
	case PATCH:
		fallthrough
	case PUT:
		log.Printf("Body is %s", rp)
		req, err = http.NewRequest(method, uri, bytes.NewBuffer([]byte(rp)))
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(OAuthBasicToken, "x-oauth-basic")
	default:
		panic("error")
	}

	return req, err
}

type Client struct {
	rf RequestFactory
	dc http.Client
}

type Callable interface {
	Call(method string, requestParams string, uri string) (*http.Response, error)
}

func (c Client)Call(method string, rp string, uri string) (*http.Response, error) {

	log.Print(uri)
	req, err := c.rf.NewRequest(method, rp, uri)
	if (err != nil) {
		log.Printf("%v", err)
	}

	resp, err := c.dc.Do(req)
	if (err != nil) {
		log.Printf("%v", err)
	}

	log.Printf("%v %v %v", int(resp.StatusCode) != http.StatusOK || int(resp.StatusCode) != 201, resp.StatusCode, http.StatusOK)

	if http.StatusOK <= resp.StatusCode && resp.StatusCode >= 201 {
		body, _ := ioutil.ReadAll(resp.Body);
		resp.Body.Close()

		log.Printf("Call %s failed with status code %d", method, resp.StatusCode)
		log.Printf("The reasons are: %v", string(body))
	}

	return resp, err
}

func NewClient(cr RequestFactory) *Client {
	return &Client{rf:cr, dc: http.Client{}}
}
