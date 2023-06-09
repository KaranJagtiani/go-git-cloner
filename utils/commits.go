package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/KaranJagtiani/go-git-cloner/types"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func GetAuthorCommits(authorEmail string, repoUrl string, daysInPast int, r *git.Repository) ([]types.Commit, error) {
	ref, err := r.Head()
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, err
	}

	// Compute zeroed 'x' days from now in the past
	t := time.Now()
	dateXDaysInPast := t.AddDate(0, 0, -daysInPast)
	dateXDaysInPast = time.Date(dateXDaysInPast.Year(), dateXDaysInPast.Month(), dateXDaysInPast.Day(), 0, 0, 0, 0, dateXDaysInPast.Location())

	var repoCommits []types.Commit
	err = cIter.ForEach(func(c *object.Commit) error {
		if (c.Author.Email == authorEmail) {
			formattedDate := c.Author.When.Format("02-01-2006 15:04")
			splitCommitUrl := strings.Split(repoUrl, ":")[1]
			commitUrl := fmt.Sprintf("https://github.com/%s/commit/%s", splitCommitUrl[0 : len(splitCommitUrl)-4], c.Hash)
			if (c.Author.When.After(dateXDaysInPast)) {
				repoCommits = append(repoCommits, types.Commit{
					Message: c.Message,
					Url: commitUrl,
					FormattedDate: formattedDate,
					Date: c.Author.When,
				})
			}
		}
		return nil
	})

	return repoCommits, err
}
