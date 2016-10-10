package api

const (
	Expected_Post_Value = "b"
)

func GetTestParams() string {
	return "q=b"
}

func PostBodytParams() string {
	return "{\"Title\":\"bugfix\",\"Body\":\"test body\",\"Milestone\":0,\"Labels\":[],\"Assignees\":{\"Login\":\"mauipipe@gmail.com\",\"html_url\":\"http://google.com\"}}"
}