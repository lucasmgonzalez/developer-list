package resources

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not could found this resource!")
}

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Tudo certo :D")
}

func PegarUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%v", "mocando users")
}
