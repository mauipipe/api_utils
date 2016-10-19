package main

import (
	"github.com/mauipipe/api_utils"
	"fmt"
	"bitbucket/mauipipe/githubclient"
)

func main() {
	l := api_utils.NewLogger()
	is := githubclient.NewIssueService()
	isq := githubclient.IssueSearchQuery{
		Q:[]string{"json"},
	}

	defer l.CollectPanic(api_utils.FileSystemWriter)

	res, err := is.SearchIssue(&isq);
	if err != nil {

	}

	fmt.Sprintf("%v", res)

	fmt.Println("End")
}
