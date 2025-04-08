package service

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v49/github"
	"golang.org/x/oauth2"
)

var token string

func init(){
	token = //github token
}

func GeneratePR(){
	
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	owner := "your-github-username"  
	repo := "your-repository"
	baseBranch := "main"            
	compareBranch := "feature-branch"
	
	pr := &github.NewPullRequest{
		Title:               github.String("Your PR Title"),
		Body:                github.String("A description of the changes you're making."),
		Head:                github.String(compareBranch), 
		Base:                github.String(baseBranch), 
	}

	createdPR, _, err := client.PullRequests.Create(context.Background(), owner, repo, pr)
	if err != nil {
		log.Fatalf("Error creating PR: %v", err)
	}

	fmt.Printf("Created PR #%d: %s\n", *createdPR.Number, *createdPR.HTMLURL)

}