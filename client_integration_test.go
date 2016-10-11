package api_utils

import (
	"net/http/httptest"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
	"github.com/onsi/gomega"
	"log"
	"io/ioutil"
	"fmt"
	"os"
)

const mockPostResponse = "{\"Title\":\"bugfix\",\"Body\":\"test body\",\"Milestone\":0,\"Labels\":[],\"Assignees\":{\"Login\":\"mauipipe@gmail.com\",\"html_url\":\"http://google.com\"}}"

func Handlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc(ExpectedCall, addIssueMockHandler).Methods(POST)
	r.HandleFunc(ExpectedCall, addIssueMockHandler).Methods(PUT)

	return r
}

func addIssueMockHandler(w http.ResponseWriter, r *http.Request) {
	if ((r.Method == POST) || (r.Method == PUT)) {
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprint(w, mockPostResponse)
		return
	}
	log.Panicf("wrong method %s", r.Method)
}

var ts *httptest.Server

func TestMain(m *testing.M) {
	ts = httptest.NewServer(Handlers())
	defer ts.Close()
	os.Exit(m.Run())
}

type IntegrationCallNoIdempotent struct {
	method string
}

var integrationCallNoIdempotents = []IntegrationCallNoIdempotent{
	{method:POST},
	{method:PUT},
}

func TestClient_Call(t *testing.T) {
	gomega.RegisterTestingT(t)

	client := NewClient(NewClientRequest())
	uri := ts.URL + ExpectedCall
	for _, expectation := range integrationCallNoIdempotents {
		rp := NewRequestParameters(expectation.method, PostBodytParams(), uri)
		rp.AuthToken = "test_token"

		res, err := client.Call(rp)

		if (err != nil) {
			log.Panicf("%v", err);
		}

		if (err != nil) {
			log.Panic(err.Error())
		}

		body, _ := ioutil.ReadAll(res.Body);

		gomega.Expect(http.StatusOK).Should(gomega.Equal(res.StatusCode))
		gomega.Expect(mockPostResponse).Should(gomega.Equal(string(body)))
	}

}