# Prometheus Gitlab Exporter

Exposes metrics for your repositories from the Gitlab API, to a Prometheus compatible endpoint.

## Configuration

This exporter is setup to take input from environment variables. All variables are optional:

* `GROUPS` If supplied, the exporter will enumerate all repositories for that group. Expected in the format "group1, group2".
* `REPOS` If supplied, The repos you wish to monitor, expected in the format "group/repo1, group/repo2". Can be across different Gitlab users/orgs.
* `GITLAB_TOKEN` - gitlab authentication token that allows the API to be queried. If none supplied, the exporter will only have access to public repos. The token must have `read_api` access.
* `GITLAB_TOKEN_FILE` If supplied _instead of_ `GITLAB_TOKEN`, enables the user to supply a path to a file containing a gitlab authentication token.
* `API_URL` Gitlab API URL
* `LOG_LEVEL` The level of logging the exporter will run with, defaults to `debug`

## Install and deploy

Build a docker image:
```
docker build -t <image-name> .
docker tag <built_image_hash> gitlab-exporter:latest
```