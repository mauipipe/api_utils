package api

import (
	"testing"
	"github.com/onsi/gomega"
	"io/ioutil"
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
	{method:GET, params:GetTestParams(), host:ExpectedHost, expectedResult:"q=b", expectedCall:ExpectedCall},
	{method:DELETE, params:GetTestParams(), host:ExpectedHost, expectedResult:"q=b", expectedCall:ExpectedCall},
}

func TestClientRequest_NewRequestIdemPotent(t *testing.T) {
	gomega.RegisterTestingT(t)

	client := NewClientRequest()

	for _, ct := range callTests {
		res, err := client.NewRequest(ct.method, ct.params, TestCall)

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
	{method:POST, params: PostBodytParams(), host:ExpectedHost, expectedResult: PostBodytParams(), expectedCall:ExpectedCall},
	{method:PUT, params:PostBodytParams(), host:ExpectedHost, expectedResult: PostBodytParams(), expectedCall:ExpectedCall},
}

func TestClientRequest_NewRequestPOST(t *testing.T) {
	gomega.RegisterTestingT(t)

	client := NewClientRequest()

	for _, ct := range callTestsNotidemPotent {
		res, err := client.NewRequest(ct.method, ct.params, TestCall)

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