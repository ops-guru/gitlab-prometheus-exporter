package exporter

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// AddMetrics - Add's all of the metrics to a map of strings, returns the map.
func AddMetrics() map[string]*prometheus.Desc {

	APIMetrics := make(map[string]*prometheus.Desc)

	APIMetrics["MergeRequestID"] = prometheus.NewDesc(
		prometheus.BuildFQName("gitlab", "repo", "merge_req_id"),
		"Merge request ID",
		[]string{"repo_name", "mr_id", "author", "title", "merge_status"}, nil,
	)

	APIMetrics["Commits"] = prometheus.NewDesc(
		prometheus.BuildFQName("gitlab", "repo", "commit_info"),
		"Commit info",
		[]string{"repo_name", "commit_id", "author"}, nil,
	)

	APIMetrics["Releases"] = prometheus.NewDesc(
		prometheus.BuildFQName("gitlab", "repo", "releases"),
		"Releases",
		[]string{"repo_name", "release_name", "release_tag"}, nil,
	)

	return APIMetrics
}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(data []*Datum, ch chan<- prometheus.Metric) error {

	// APIMetrics - range through the data slice
	for _, x := range data {
		for _, mr := range x.MergeRequests {
			ch <- prometheus.MustNewConstMetric(e.APIMetrics["MergeRequestID"], prometheus.CounterValue, 1, x.RepoName, strconv.Itoa(mr.ID), mr.Author.Name, mr.Title, mr.State)
		}

		//for _, b := range x.Commits {
		// for _, c := range b.BranchCommits {
		// 	ch <- prometheus.MustNewConstMetric(e.APIMetrics["Commits"], prometheus.GaugeValue, 0.0, x.RepoName, b.Branch, c.ID, c.AuthorName, c.CreatedAt.String())
		// }
		for _, c := range x.Commits {
			ch <- prometheus.MustNewConstMetric(e.APIMetrics["Commits"], prometheus.CounterValue, 1, x.RepoName, c.ID, c.AuthorName)

		}

		for _, r := range x.Releases {
			ch <- prometheus.MustNewConstMetric(e.APIMetrics["Releases"], prometheus.CounterValue, 1, x.RepoName, r.Name, r.TagName)
		}

		//ch <- prometheus.MustNewConstMetric(e.APIMetrics["MergeRequestTotal"], prometheus.GaugeValue, float64(mrCount))
	}

	return nil
}
