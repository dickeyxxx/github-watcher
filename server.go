package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/google/go-github/github"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/github/webhook", GithubHook)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

func GithubHook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var push github.PushEvent
	json.Unmarshal(body, &push)
	if err != nil {
		panic(err)
	}
	if push.Ref != nil {
		branch := strings.Split(*push.Ref, "/")[2]
		fmt.Printf("[%s] %s updated\n", *push.Repo.FullName, branch)
	}
}
