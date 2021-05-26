package exporter

import (
	log "github.com/sirupsen/logrus"
	goGitlab "github.com/xanzy/go-gitlab"
)

// XXX gatherData - Collects the data from the API and stores into struct
func (e *Exporter) gatherData() ([]*Datum, error) {

	data := []*Datum{}

	// this needs to use something returned from getGroups
	for _, repo := range e.Config.Repositories {
		git, err := goGitlab.NewClient(e.Config.APIToken, goGitlab.WithBaseURL(e.Config.APIURL))
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		d := new(Datum)

		d.RepoName = repo
		log.Infof("API data started for repository: %s", repo)
		d.Branches = getBranches(repo, git)
		log.Infof("Branch gathered for repository: %s", repo)
		d.Releases = getReleases(repo, git)
		log.Infof("Releases gathered for repository: %s", repo)
		d.MergeRequests = getMergeRequests(repo, git)
		log.Infof("MergeRequests gathered for repository: %s", repo)
		d.Commits = getCommits(repo, git)
		log.Infof("Commits gathered for repository: %s", repo)
		// for _, b := range d.Branches {
		// 	c := new(CommitsPerBranch)
		// 	c.Branch = b.Name
		// 	c.BranchCommits = getCommits(repo, b.Name, git)
		// 	log.Infof("Branch gathered for repository: %s branch: %s ", repo, b.Name)
		// 	d.Commits = append(d.Commits, c)
		// }

		data = append(data, d)

		log.Infof("API data fetched for repository: %s", repo)
	}

	//return data, rates, err
	return data, nil

}

type MergeRequests []*goGitlab.MergeRequest

func getMergeRequests(repo string, client *goGitlab.Client) MergeRequests {
	var mrs MergeRequests
	opts := &goGitlab.ListProjectMergeRequestsOptions{
		ListOptions: goGitlab.ListOptions{
			Page: 0,
		},
	}
	for {
		pageMRs, resp, err := client.MergeRequests.ListProjectMergeRequests(repo, opts)
		if err != nil {
			log.Errorf("Unable to obtain merge requests from API, Error: %s", err)
		}

		mrs = append(mrs, pageMRs...)

		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opts.Page = resp.NextPage
	}

	return mrs
}

type Commits []*goGitlab.Commit

func getCommits(repo string, branch string, client *goGitlab.Client) Commits {
	var commits Commits
	opts := &goGitlab.ListCommitsOptions{
		RefName: &branch,
		ListOptions: goGitlab.ListOptions{
			Page: 0,
		},
	}
	for {
		pageCs, resp, err := client.Commits.ListCommits(repo, opts)
		if err != nil {
			log.Errorf("Unable to obtain commits from API, Error: %s", err)
		}

		commits = append(commits, pageCs...)

		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opts.Page = resp.NextPage
	}

	return commits
}

type Releases []*goGitlab.Release

func getReleases(repo string, client *goGitlab.Client) Releases {
	var releases Releases
	releases, _, err := client.Releases.ListReleases(repo, &goGitlab.ListReleasesOptions{PerPage: 100})
	if err != nil {
		log.Errorf("Unable to obtain releases from API, Error: %s", err)
	}
	return releases
}

type Branches []*goGitlab.Branch

func getBranches(repo string, client *goGitlab.Client) Branches {
	var branches Branches
	pageOpts := goGitlab.ListOptions{PerPage: 100}
	branches, _, err := client.Branches.ListBranches(repo, &goGitlab.ListBranchesOptions{ListOptions: pageOpts})
	if err != nil {
		log.Errorf("Unable to obtain branches from API, Error: %s", err)
	}
	return branches
}
