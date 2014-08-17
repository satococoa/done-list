package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/google/go-github/github"
	"github.com/satococoa/github-auth/client"
)

func main() {
	client := createClient()
	issues, err := fetchIssues(client)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printDoneList(issues)
}

func createClient() *github.Client {
	return client.CreateClient("done-list-golang", []string{"repo", "public_repo", "read:org"})
}

func fetchIssues(client *github.Client) ([]github.Issue, error) {
	options := &github.IssueListOptions{
		Filter: "subscribed",
		Sort:   "updated",
		Since:  time.Now().AddDate(0, 0, -1),
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}
	issues, _, err := client.Issues.List(true, options)
	return issues, err
}

func issuesByRepo(issues []github.Issue) map[string][]github.Issue {
	repoNameWithType := func(issue github.Issue) string {
		url := *issue.URL
		exp, _ := regexp.Compile("https://api.github.com/repos/(.+)/[0-9]+$")
		match := exp.FindStringSubmatch(url)
		return match[1]
	}
	issuesByRepo := make(map[string][]github.Issue)

	for _, issue := range issues {
		name := repoNameWithType(issue)
		issuesByRepo[name] = append(issuesByRepo[name], issue)
	}
	return issuesByRepo
}

func printDoneList(issues []github.Issue) {
	collectedIssues := issuesByRepo(issues)
	for name, issues := range collectedIssues {
		fmt.Printf("## %s\n", name)
		for _, issue := range issues {
			var mark string
			if *issue.State == "closed" {
				mark = "x"
			} else {
				mark = " "
			}
			fmt.Printf("- [%s] [%s by %s](%s)\n", mark, *issue.Title, *issue.User.Login, *issue.HTMLURL)
		}
		fmt.Println("")
	}
}
