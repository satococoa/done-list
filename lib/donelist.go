package donelist

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func CreateClient(accessToken string) *github.Client {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tokenClient := oauth2.NewClient(ctx, tokenSource)
	return github.NewClient(tokenClient)
}

func FetchIssues(client *github.Client) ([]*github.Issue, error) {
	options := &github.IssueListOptions{
		Filter: "subscribed",
		State:  "all",
		Sort:   "updated",
		Since:  time.Now().AddDate(0, 0, -1),
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}
	issues, _, err := client.Issues.List(context.Background(), true, options)
	return issues, err
}

func issuesByRepo(issues []*github.Issue) map[string][]*github.Issue {
	repoNameWithType := func(issue *github.Issue) string {
		htmlURL := *issue.HTMLURL
		exp, _ := regexp.Compile("https://github.com/(.+)/[0-9]+$")
		match := exp.FindStringSubmatch(htmlURL)
		return match[1]
	}
	issuesByRepo := make(map[string][]*github.Issue)

	for _, issue := range issues {
		name := repoNameWithType(issue)
		issuesByRepo[name] = append(issuesByRepo[name], issue)
	}
	return issuesByRepo
}

func PrintDoneList(writer io.Writer, issues []*github.Issue) {
	collectedIssues := issuesByRepo(issues)
	for name, issues := range collectedIssues {
		fmt.Fprintf(writer, "## %s\n", name)
		for _, issue := range issues {
			var mark string
			if *issue.State == "closed" {
				mark = "x"
			} else {
				mark = " "
			}
			fmt.Fprintf(writer, "- [%s] [%s by %s](%s)\n", mark, *issue.Title, *issue.User.Login, *issue.HTMLURL)
		}
		fmt.Fprintln(writer, "")
	}
}
