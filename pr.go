package main

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// TODO extract real domain
// ಠ_ಠ
func main() {
	_, err := exec.Command("ping", "-c 1", "gitlab.liip.ch").Output()
	if err != nil {
		log.Fatal("Network is down... or you forgot VPN")
	}

	remote, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		log.Fatal("Go inside git repo!")
	}

	repoUrl := GetRepoUrl(string(remote))

	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatal("You should be on a branch...")
	}

	prUrl := GetPullRequestUrl(repoUrl, string(branch))

    exec.Command("open", prUrl).Output()

	log.Println("(╯°□°）╯︵ ┻━┻ ", prUrl)
}

// ¯\_(ツ)_/¯
func GetRepoUrl(url string) string {
	r := regexp.MustCompile(`origin\s*http([a-z]://.+)`)

	switch r.MatchString(url) {
	case true:
		url = r.FindString(url)
		words := strings.Fields(url)
		url = words[1]
	default:
		r := regexp.MustCompile(`^origin(.*)git@(.*)git`)
		url = strings.Replace(r.FindString(url), ":", "/", -1)
		url = strings.Replace(url, `git@`, "https://", 1)
	}
	url = strings.Replace(url, ".git", "", 1)
	url = strings.Replace(url, "origin", "", 1)
	url = strings.TrimSpace(url)

	return url
}

func GetPullRequestUrl(repoUrl string, branchName string) string {
	r := regexp.MustCompile("[^a-zA-Z0-9/-]+")
	branch := r.ReplaceAllString(branchName, "")
	return repoUrl + "/compare/" + branch + "?expand=1"
}
