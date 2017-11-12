package github

import (
	"io/ioutil"
	"log"
	"net/http"
)

const serviceEndpoint = "http://api.github.com/"
const accessToken = ""

func GetContributorsStatsFromRepository(owner, repo string) (response []byte) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", serviceEndpoint+"repos/"+owner+"/"+repo+"/stats/contributors?access_token="+accessToken, nil)

	if err != nil {
		// Handle Error
		log.Fatal("Error na criação request: %+v", err)
	}

	resp, err := client.Do(req)

	if err != nil {
		// Handle Error
		log.Fatal("Error na request: %+v", err)
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		// Handle Error
		log.Fatal("Erro na leitura: %+v", err)
	}

	return
}
