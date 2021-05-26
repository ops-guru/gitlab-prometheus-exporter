package exporter

import (
	"github.com/andreip-og/gitlab-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	goGitlab "github.com/xanzy/go-gitlab"
)

// Exporter is used to store Metrics data and embeds the config struct.
// This is done so that the relevant functions have easy access to the
// user defined runtime configuration when the Collect method is called.
type Exporter struct {
	APIMetrics map[string]*prometheus.Desc
	config.Config
}

// Data is used to store an array of Datums.
// This is useful for the JSON array detection
type Data []Datum

// Datum is used to store data from all the relevant endpoints in the API
type Datum struct {
	RepoName      string
	Branches      []*goGitlab.Branch
	MergeRequests []*goGitlab.MergeRequest
	Releases      []*goGitlab.Release
	Commits       []*goGitlab.Commit
}

// type CommitsPerBranch struct {
// 	Branch        string
// 	BranchCommits []*goGitlab.Commit
// }
