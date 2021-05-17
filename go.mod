module github.com/andreip-og/gitlab-exporter

go 1.16

require (
	github.com/fatih/structs v1.1.0
	github.com/infinityworks/github-exporter v0.0.0-20201016091012-831b72461034
	github.com/prometheus/client_golang v1.10.0
	github.com/sirupsen/logrus v1.8.1
	github.com/steinfletcher/apitest v1.5.6
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80
	github.com/xanzy/go-gitlab v0.49.0
)

replace github.com/andreip-og/gitlab-exporter => ../gitlab-exporter
