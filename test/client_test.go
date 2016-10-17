package api_utils_test

import (
	"testing"
	"github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"github.com/mauipipe/api_utils"
)

const (
	Host = "http://test.com"
	ExpectedHost = "test.com"
	ExpectedCall = "/test"
	TestCall = Host + ExpectedCall
)

type callTestClient struct {
	method         string
	params         string
	host           string
	expectedResult string
	expectedCall   string
}

var callTests = []callTestClient{
	{method:http.MethodGet, params:GetTestParams(), host:ExpectedHost, expectedResult:"q=b", expectedCall:ExpectedCall},
	{method:http.MethodDelete, params:GetTestParams(), host:ExpectedHost, expectedResult:"q=b", expectedCall:ExpectedCall},
}

func TestClientRequest_NewRequestIdemPotent(t *testing.T) {
	gomega.RegisterTestingT(t)

	request := api_utils.NewClientRequest()

	for _, ct := range callTests {
		rp := api_utils.NewRequestParameters(ct.method, ct.params, TestCall)
		res, err := request.NewRequest(rp)

		if err != nil {
			panic(err)
		}

		gomega.Expect(ct.method).Should(gomega.Equal(res.Method))
		gomega.Expect(ct.host).Should(gomega.Equal(res.Host))
		gomega.Expect(ct.expectedResult).Should(gomega.Equal(res.URL.RawQuery))
		gomega.Expect(ct.expectedCall).Should(gomega.Equal(res.URL.Path))
	}
}

var callTestsNotidemPotent = []callTestClient{
	{method:http.MethodPost, params: PostBodytParams(), host:ExpectedHost, expectedResult: PostBodytParams(), expectedCall:ExpectedCall},
	{method:http.MethodPut, params:PostBodytParams(), host:ExpectedHost, expectedResult: PostBodytParams(), expectedCall:ExpectedCall},
}

func TestClientRequest_NewRequestPOST(t *testing.T) {
	gomega.RegisterTestingT(t)

	client := api_utils.NewClientRequest()

	for _, ct := range callTestsNotidemPotent {
		rp := api_utils.NewRequestParameters(ct.method, ct.params, TestCall)
		res, err := client.NewRequest(rp)

		if err != nil {
			panic(err)
		}

		res.ParseForm()
		body, _ := ioutil.ReadAll(res.Body)

		gomega.Expect(ct.method).Should(gomega.Equal(res.Method))
		gomega.Expect(ct.host).Should(gomega.Equal(res.Host))
		gomega.Expect(PostBodytParams()).Should(gomega.Equal(string(body)))
		gomega.Expect(ct.expectedCall).Should(gomega.Equal(res.URL.Path))
	}
}