package github

import (
	"encoding/json"
	"log"
	"sort"
	"strings"

	"github.com/lucasmgonzalez/developer-list/services/github"
)

type GithubUser struct {
	Login string `json:"login"`
	Id    int    `json:"id"`
}

type GithubContribution struct {
	Total  int        `json:"total"`
	Author GithubUser `json:"author"`
}

func GetContributionsFromRepositoryGroupedByAuthor(repo, owner, order string) (contributions []GithubContribution) {

	// Retrieve contributions from service
	serviceResponse := github.GetContributorsStatsFromRepository(repo, owner)

	dec := json.NewDecoder(strings.NewReader(string(serviceResponse)))

	for dec.More() {
		err := dec.Decode(&contributions)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Rank contributions
	orderFunction := func(i, j int) bool {
		return contributions[i].Total > contributions[j].Total
	}

	if order == "asc" {
		orderFunction = func(i, j int) bool {
			return contributions[i].Total < contributions[j].Total
		}
	}

	sort.Slice(contributions, orderFunction)

	return
}
