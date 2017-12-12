package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// ( •_•)
func TestGetUrl(t *testing.T) {
	testCase := `origin   git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (fetch)`
	expected := `https://git.intern.orange-food.net/agentur-liip/ch.orange-food.produkte`
	assert.Equal(t, expected, GetRepoUrl(testCase), "works for ssh remote")
}

// ( •_•)>⌐■-■
func TestGetUrl3(t *testing.T) {
	testCase := `origin	git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (fetch)
origin	git@git.intern.orange-food.net:agentur-liip/ch.orange-food.produkte.git (push)`
	expected := `https://git.intern.orange-food.net/agentur-liip/ch.orange-food.produkte`
	assert.Equal(t, expected, GetRepoUrl(testCase), "works for ssh remote")
}

// (⌐■_■)
func TestGetUrl4(t *testing.T)  {
	testCase := `heroku	https://git.heroku.com/murmuring-island-99377.git (fetch)
heroku	https://git.heroku.com/murmuring-island-99377.git (push)
origin	https://github.com/syzer/poker-player-go (fetch)
origin	https://github.com/syzer/poker-player-go (push)`
	expected  := `https://github.com/syzer/poker-player-go`
	assert.Equal(t, expected, GetRepoUrl(testCase), "works for with heroku remotes too")
}

// ┬─┬ ノ(º_ºノ)
func TestGetUrl2(t *testing.T) {
	testCase := `origin  https://github.com/syzer/repo.git (fetch)`
	expected := `https://github.com/syzer/repo`
	assert.Equal(t, GetRepoUrl(testCase), expected, "works for https remote")
}

// (╯°□°）╯┬─┬
func TestGetPullRequestUrl(t *testing.T) {
	// given
	testUrl, branchName := "https://github.com/syzer/pr", "feature/branch-cheacking"

	// expected
	expected := "https://github.com/syzer/pr/compare/feature/branch-cheacking?expand=1"
	//expected := "https://github.com/syzer/pr/compare/master...feature/branch-cheacking?expand=1"

	// test
	assert.Equal(t, GetPullRequestUrl(testUrl, branchName), expected, "Works for github PR")
}
