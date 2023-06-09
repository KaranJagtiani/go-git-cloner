# GitHub Repository Cloner and Commit Crawler

This is a Go application that clones a list of repositories provided by the user and then crawls the commits of those repositories without using the GitHub APIs.

## Features

- Clone multiple GitHub repositories using SSH.
- Crawl through the commit history of each repository.
- Crawl till specific date in the past.
- Perform these actions in a secure manner using your personal SSH keys.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/dl/) installed on your local machine.
- SSH keys setup on your GitHub account - [Instructions](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).

### Installation

1. Clone the repository:

```bash
git clone git@github.com:KaranJagtiani/go-git-cloner.git
```

2. Modify the `config.yaml` file:

- author_email - Provide the email for the author whose commits you want to crawl.
- crawl_x_days_in_past - How many days in the past do you want to crawl.
- repositories - List of repositories that you want to crawl.

3. Build the project

```bash
go build -o out/go-git-cloner
```

4. Run the project

```bash
./out/go-git-cloner
```

## Contributing

We welcome contributions from the community! If you wish to contribute, please follow these steps:

1. Fork the project
2. Create your feature branch (git checkout -b feature/AmazingFeature)
3. Commit your changes (git commit -m 'Add some AmazingFeature')
4. Push to the branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
