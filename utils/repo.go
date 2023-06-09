package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func CloneRepo(url string, dir string) (*git.Repository, error) {
	privateSSHKeyPath := fmt.Sprintf("./ssh_key/%s", getSSHKeyFileName())

	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateSSHKeyPath, "")
	if err != nil {
		return nil, err
	}
	
	// Check if the directory already exists
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		// If the directory exists, delete it
		err = os.RemoveAll(dir)
		if err != nil {
			return nil, err
		}
	}
	
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
		Auth: publicKeys,
		Progress: os.Stdout,
	})

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func getSSHKeyFileName() string {
	files, err := ioutil.ReadDir("./ssh_key")
	if err != nil {
		log.Fatal(err)
	}

	if len(files) > 2 {
		log.Fatal("Please provide only one SSH key file in the ssh_key folder.")
	}

	return files[1].Name()
}