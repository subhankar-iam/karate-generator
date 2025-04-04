package service

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var branch string
var token string
var username string

func init() {
	godotenv.Load()
	branch = os.Getenv("BRANCH")
	token = os.Getenv("GITHUB_API_KEY")
	username = os.Getenv("USER_NAME")

}

func CommitAndPush(basePath string) {
	fmt.Println(username)
	r, err := git.PlainOpen(basePath)
	if err != nil {
		fmt.Println(err)
	}
	w, err := r.Worktree()
	if err != nil {
		fmt.Println(err)
	}
	_, err = w.Add(".")
	if err != nil {
		fmt.Println(err)
	}
	_, err = w.Commit("Test Genie generated", &git.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: "TestGenie@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
		Create: false,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = r.Push(&git.PushOptions{
		RemoteName: "origin2",
		Auth: &http.BasicAuth{
			Username: "git",
			Password: token,
		},
		RefSpecs: []config.RefSpec{
			config.RefSpec("refs/heads/" + branch + ":refs/heads/" + branch),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Commit Success")
}
