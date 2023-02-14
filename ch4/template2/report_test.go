package template2

import (
	"log"
	"os"
	"testing"

	"lemonguge.cn/learn_go/ch4/github"
)

func TestSearchIssues(t *testing.T) {
	result, err := github.SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func TestSearchIssues2(t *testing.T) {
	result, err := github.SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report2.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
