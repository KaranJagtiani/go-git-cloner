package main

import (
	"fmt"
	"os"
	"path"

	"github.com/KaranJagtiani/go-git-cloner/types"
	"github.com/KaranJagtiani/go-git-cloner/utils"
)

func main() {
	// Parse config file
	config, err := utils.ParseYaml("config.yaml")
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		os.Exit(1)
	}

	// Clone repositories
	var repositories []types.Repository
	for _, repo := range config.Repositories {
		fmt.Printf("Cloning repository: %s\n", repo.URL)
		
		dir := path.Join("/tmp", path.Base(repo.URL))
		r, err := utils.CloneRepo(repo.URL, dir)
		if err != nil {
			fmt.Printf("Error cloning repository: %s\n", err)
			continue
		}

		newRepo := types.Repository{
			URL: repo.URL,
		}

		commits, err := utils.GetAuthorCommits(config.AuthorEmail, repo.URL, config.CrawlXDaysInPast, r)
		if err != nil {
			fmt.Printf("Error in fetching commits: %s\n", err)
		}
		
		newRepo.Commits = commits
		repositories = append(repositories, newRepo)
	}
	
	// Print commits
	for _, repo := range repositories {
		fmt.Printf("Printing commits for repository: %s\n", repo.URL)
		for _, commit := range repo.Commits {
			fmt.Println("-----")
			fmt.Printf("URL: %s\n", commit.Url)
			fmt.Printf("Date: %s\n", commit.FormattedDate)
			fmt.Printf("Message: %s\n", commit.Message)	
		}
	}
}
