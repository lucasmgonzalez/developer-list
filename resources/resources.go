package resources

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucasmgonzalez/developer-list/repositories/github"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not could found this resource!")
}

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Tudo certo :D")
}

func PegarUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var owner = r.URL.Query().Get("owner")
	var repo = r.URL.Query().Get("repo")
	var order = r.URL.Query().Get("order")

	contributions := github.GetContributionsFromRepositoryGroupedByAuthor(owner, repo, order)

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(contributions)

}
