package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newsparser/source/habr"
	"newsparser/source/models"
)

type GetPost struct {
	Page    int      `json:"page"`
	Subject []string `json:"subject"`
}

func postsRoute() {
	http.HandleFunc("/posts", postsRouteHandler)
}

func postsRouteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {

	case "POST":

		var p GetPost

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(p)

		fmt.Fprintf(w, string(getPosts(p.Page, p.Subject)))
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func getPosts(pages int, subject []string) []byte {
	var postsCollection []models.Post

	for i := 1; i <= pages; i++ {
		posts := habr.GetPostsByIndex(i, subject)
		postsCollection = append(postsCollection, posts...)
	}

	habr.SendPostToBot(postsCollection)
	toJSON, _ := json.Marshal(postsCollection)
	return toJSON
}
