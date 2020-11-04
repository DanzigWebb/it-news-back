package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"newsparser/source/habr"
	"newsparser/source/models"
	"strings"
)

type GetPost struct {
	Page    int      `json:"page"`
	Subject []string `json:"subject"`
}

type WebResponse struct {
	All   int           `json:"allPosts"`
	Posts []models.Post `json:"posts"`
}

func postsRoute() {
	http.HandleFunc("/posts", postsRouteHandler)
}

func postsRouteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

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

		fmt.Fprint(w, string(getHabrPostsByRss(p.Page, p.Subject)))
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func getHabrPostsByRss(pages int, subjects []string) []byte {
	var postsCollection []models.Post

	for i := 1; i <= pages; i++ {
		posts := habr.ParseRSSByPageIndex(i)
		postsCollection = append(postsCollection, posts...)
	}

	// Todo вынести сортировку статей в парсинг (ParseAndFilterRss)
	var output []models.Post
	for _, post := range postsCollection {
		for _, subject := range subjects {
			for _, category := range post.Categories {
				if strings.Contains(category, subject) {
					output = append(output, post)
					break
				}
			}
		}
	}

	var resp = WebResponse{
		All:   len(postsCollection),
		Posts: output,
	}

	toJSON, err := json.Marshal(resp)

	if err != nil {
		log.Fatal(err)
	}

	return toJSON
}

//func getPosts(pages int, subject []string) []byte {
//	var postsCollection []models.Post
//
//	for i := 1; i <= pages; i++ {
//		posts := habr.GetPostsByIndex(i, subject)
//		postsCollection = append(postsCollection, posts...)
//	}
//
//	habr.SendPostToBot(postsCollection)
//	toJSON, _ := json.Marshal(postsCollection)
//	return toJSON
//}
