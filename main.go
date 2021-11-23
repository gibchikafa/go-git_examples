package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Basic example of how to commit changes to the current branch to an existing
// repository.
func main() {
	CheckArgs("<directory>")
	directory := os.Args[1]

	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	// We can verify the current status of the worktree using the method Status.
	Info("git status --porcelain")
	status, err := w.Status()
	CheckIfError(err)
	fmt.Println(status)

	err = w.AddWithOptions(&git.AddOptions{All: true})
	CheckIfError(err)

	Info("git status --porcelain. After add")
	s, err := w.Status()
	CheckIfError(err)
	fmt.Println(s)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit Since version 5.0.1, we can omit the Author signature, being read
	// from the git config files.
	Info("git commit -m \"example go-git commit\"")
	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	CheckIfError(err)

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)
	fmt.Println(obj)

	// We can verify the current status of the worktree using the method Status.
	Info("git status --porcelain. After commit")
	status2, err := w.Status()
	CheckIfError(err)
	fmt.Println(status2)
}
