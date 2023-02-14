package github_test

import (
	"fmt"
	"log"
	"testing"

	"lemonguge.cn/learn_go/ch4/github"
)

func TestSearchIssues(t *testing.T) {
	result, err := github.SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
