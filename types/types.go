package types

import "time"

type Config struct {
	AuthorEmail string `yaml:"author_email"`
	CrawlXDaysInPast int `yaml:"crawl_x_days_in_past"`
	Repositories []Repository `yaml:"repositories"`
}

type Repository struct {
	URL string `yaml:"url"`
	Commits []Commit
}

type Commit struct {
	FormattedDate string
	Message string
	Url string
	Date time.Time
}