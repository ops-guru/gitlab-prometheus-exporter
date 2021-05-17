package config

import (
	"fmt"
	goGitlab "github.com/xanzy/go-gitlab"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"

	"os"
)

// GetEnv - Allows us to supply a fallback option if nothing specified
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// Config struct holds all of the runtime confgiguration for the application
type Config struct {
	APIURL        string
	Repositories  []string
	Organisations []string
	Users         string
	APITokenEnv   string
	APITokenFile  string
	APIToken      string
}

// Init populates the Config struct based on environmental runtime configuration
func Init() Config {

	listenPort := GetEnv("LISTEN_PORT", "9171")
	os.Setenv("LISTEN_PORT", listenPort)
	url := "https://gitlab.com"
	repos := os.Getenv("REPOS")
	groups := os.Getenv("GROUPS")
	users := os.Getenv("USERS")
	tokenEnv := os.Getenv("GITLAB_TOKEN")
	tokenFile := os.Getenv("GITLAB_TOKEN_FILE")
	token, err := getAuth(tokenEnv, tokenFile)

	if err != nil {
		log.Errorf("Error initialising Configuration, Error: %v", err)
	}

	var reposList []string
	if repos != "" {
		rs := strings.Split(repos, ", ")
		for _, x := range rs {
			reposList = append(reposList, x)
		}
	}

	var groupList []string
	if groups != "" {
		gs := strings.Split(groups, ", ")
		for _, x := range gs {
			groupList = append(reposList, x)
			groupRepos := getReposByGroup(x, token)
			reposList = append(reposList, groupRepos...)
		}
	}


	appConfig := Config{
		url,
		reposList,
		groupList,
		users,
		tokenEnv,
		tokenFile,
		token,
	}

	return appConfig
}

// getAuth returns oauth2 token as string for usage in http.request
func getAuth(token string, tokenFile string) (string, error) {

	if token != "" {
		return token, nil
	} else if tokenFile != "" {
		b, err := ioutil.ReadFile(tokenFile)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(b)), err

	}

	return "", nil
}

// getReposByGroup returns a list of repositories
// that belong to a group
func getReposByGroup(group string, token string) []string {
	fmt.Println(token)
	git, err := goGitlab.NewClient(token)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil
	}
	groupInfo, _, err := git.Groups.GetGroup(group)
	if err != nil {
		log.Errorf("Error obtaining repositories from group, Error: %v", err)
		return nil
	}
	var repos []string
	for _, p := range groupInfo.Projects {
		fullName := group + "/" + p.Path
		repos = append(repos, fullName)
	}

	return repos
}