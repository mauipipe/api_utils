package api_test

import (
	"net/http/httptest"
	"github.com/gorilla/mux"
	"net/http"
	"api"
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

	r.HandleFunc(api.ExpectedCall, addIssueMockHandler).Methods(api.POST)
	r.HandleFunc(api.ExpectedCall, addIssueMockHandler).Methods(api.PUT)

	return r
}

func addIssueMockHandler(w http.ResponseWriter, r *http.Request) {
	if ((r.Method == api.POST) || (r.Method == api.PUT)) {
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
	{method:api.POST},
	//{method:api.PUT},
}

func TestClient_Call(t *testing.T) {
	gomega.RegisterTestingT(t)

	client := api.NewClient(api.NewClientRequest())
	uri := ts.URL + api.ExpectedCall
	for _, expectation := range integrationCallNoIdempotents {
		res, err := client.Call(expectation.method, api.PostBodytParams(), uri)
		log.Printf("%v",res);

		if (err != nil) {
			log.Panic(err.Error())
		}

		body, _ := ioutil.ReadAll(res.Body);

		gomega.Expect(http.StatusOK).Should(gomega.Equal(res.StatusCode))
		gomega.Expect(mockPostResponse).Should(gomega.Equal(string(body)))
	}

}